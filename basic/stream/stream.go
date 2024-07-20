package main

import (
	"bufio"
	"io"
	"log"
	"net/http"
)

// 发送请求给LLM 并流式获取响应，同时

func main() {
	http.HandleFunc("/stream", func(w http.ResponseWriter, r *http.Request) {
		// 检查请求方法
		if r.Method != http.MethodGet {
			http.Error(w, "只支持 GET 请求", http.StatusMethodNotAllowed)
			return
		}

		// 创建 HTTP 客户端
		client := &http.Client{}

		// 发送请求给 LLM
		resp, err := client.Get("http://llm.example.com/stream")
		if err != nil {
			log.Println("发送请求给 LLM 失败:", err)
			http.Error(w, "内部服务器错误", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// 检查响应状态码
		if resp.StatusCode != http.StatusOK {
			log.Println("LLM 返回非 200 响应:", resp.StatusCode)
			http.Error(w, "LLM 服务错误", http.StatusInternalServerError)
			return
		}

		// 设置响应头，告诉前端我们将发送流式数据
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		// 流式传输响应
		flusher, ok := w.(http.Flusher)
		if !ok {
			log.Println("无法使用 Flusher")
			http.Error(w, "内部服务器错误", http.StatusInternalServerError)
			return
		}

		reader := bufio.NewReader(resp.Body)
		for {
			// 逐字读取
			bytee, err := reader.ReadByte()
			if err != nil {
				if err == io.EOF {
					break // 达到输入流的末尾
				}
				log.Println("读取 LLM 响应失败:", err)
				return
			}

			// 逐字写入
			_, err = w.Write([]byte{bytee})
			if err != nil {
				log.Println("写入响应失败:", err)
				return
			}

			flusher.Flush() // 每次发送后立即刷新
		}
	})

	// 启动服务器
	log.Println("服务器启动在 http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("启动服务器失败:", err)
	}
}

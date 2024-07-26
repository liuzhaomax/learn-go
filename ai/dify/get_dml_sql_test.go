package dify

/*
  这个有问题，似乎dify返回有token数量的限制
*/

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

type DifyRequest struct {
	Inputs       map[string]string `json:"inputs"`
	ResponseMode string            `json:"response_mode"`
	User         string            `json:"user"`
}

const (
	ChunkSize = 1024 // 1KB的常量定义
)

// 定义结构体用于提取嵌套的 text 字段
type ResponseData struct {
	WorkflowID string `json:"workflow_run_id"`
	TaskID     string `json:"task_id"`
	Data       struct {
		Outputs struct {
			Text string `json:"text"`
		} `json:"outputs"`
	} `json:"data"`
}

func sendDifyRequest(c *gin.Context) {
	// 构建请求体
	reqBody := DifyRequest{
		Inputs: map[string]string{
			"url": "https://arxiv.org/category_taxonomy",
		},
		ResponseMode: "streaming",
		User:         "abc-123",
	}
	reqBodyJson, err := json.Marshal(reqBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 创建HTTP请求
	req, err := http.NewRequest("POST", "https://prismer.dev/v1/workflows/run", bytes.NewBuffer(reqBodyJson))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	req.Header.Set("Authorization", "Bearer ")
	req.Header.Set("Content-Type", "application/json")

	// 发送HTTP请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer resp.Body.Close()

	// 设置内容类型和状态码
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)

	// 逐行读取响应并仅发送text字段内容
	reader := bufio.NewReader(resp.Body)
	buffer := make([]byte, 0, ChunkSize) // 定义1KB的缓冲区

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				// 发送最后的剩余数据
				if len(buffer) > 0 {
					jsonData, _ := json.Marshal(gin.H{
						"dify": string(buffer),
					})
					c.Writer.Write(jsonData)
					c.Writer.Write([]byte("\n"))
				}
				break
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}
		}

		// 清理行数据
		lineStr := strings.TrimSpace(string(line))
		if strings.HasPrefix(lineStr, "data:") {
			lineStr = strings.TrimPrefix(lineStr, "data: ")
		}

		// 解析每行JSON并提取outputs.text内容拼接到缓冲区中
		var responseData ResponseData
		err = json.Unmarshal([]byte(lineStr), &responseData)
		if err != nil {
			// 如果遇到解析错误，继续处理下一行
			continue
		}

		text := responseData.Data.Outputs.Text
		if text != "" {
			buffer = append(buffer, text...)
		}

		// 如果缓冲区超过1KB，则发送
		for len(buffer) >= ChunkSize {
			jsonData, _ := json.Marshal(gin.H{
				"dify": string(buffer[:ChunkSize]),
			})
			c.Writer.Write(jsonData)
			c.Writer.Write([]byte("\n"))
			buffer = buffer[ChunkSize:]
		}
	}
}

func TestGetDMLSQL(t *testing.T) {
	r := gin.Default()
	r.GET("/dify", sendDifyRequest)
	r.Run(":8080")
}

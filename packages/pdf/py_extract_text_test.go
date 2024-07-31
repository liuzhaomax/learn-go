package pdf

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"testing"
)

const scriptPath = "extract_text.py"

func TestPyExtractText3(t *testing.T) {
	// 创建 Gin 引擎
	r := gin.Default()

	// 设置上传路由
	r.POST("/upload", extractTextFromUploadedFile)

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}

func extractTextFromUploadedFile(c *gin.Context) {
	// 从请求中的表单文件读取文件头
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
		return
	}

	// 打开上传的文件
	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer file.Close()

	// 将文件内容读到缓冲区
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	// 调用 Python 脚本并传递 PDF 数据
	cmd := exec.Command("python", scriptPath)
	cmd.Stdin = buf

	// 获取 Python 脚本的输出
	out, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Python script error: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"extracted_text": string(out)})
}

// 输入是bytes
func TestPyExtractText2(t *testing.T) {
	pdfPath := "pdf17.pdf"
	scriptPath := "extract_text.py"
	file, err := os.Open(pdfPath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, file)
	if err != nil {
		fmt.Println(err)
	}

	// 调用 Python 脚本并传递 PDF 数据
	cmd := exec.Command("python", scriptPath)
	cmd.Stdin = buf

	// 获取 Python 脚本的输出
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(out))
}

// 输入是文件路径
func TestPyExtractText(t *testing.T) {
	// pdfPath := "D:\\workspace\\max\\learn-go\\packages\\pdf\\pdf17.pdf" // 绝对路径也可
	pdfPath := "pdf17.pdf"
	scriptPath := "extract_text.py"
	tempDir := "./" // 这里为空则默认放入系统临时文件文件夹

	file, err := os.Open(pdfPath)
	if err != nil {
		t.Fatalf("打开文件时出错: %v", err)
	}
	defer file.Close()

	// 看看tempdir是否存在
	if _, err := os.Stat(tempDir); os.IsNotExist(err) {
		err := os.MkdirAll(tempDir, 0755)
		if err != nil {
			fmt.Println(err)
		}
	}

	tempFile, err := os.CreateTemp(tempDir, "*.pdf")
	if err != nil {
		fmt.Println(err)
	}
	defer os.Remove(tempFile.Name())
	// 将内容写入临时文件
	_, err = io.Copy(tempFile, file)
	if err != nil {
		fmt.Println(err)
	}

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("python", scriptPath, tempFile.Name())
	} else {
		cmd = exec.Command("python3", scriptPath, tempFile.Name())
	}

	// 运行py文件，但go不拿系统输出
	// output, err := cmd.CombinedOutput()
	// if err != nil {
	// 	log.Fatalf("执行命令时出错: %v\n输出: %s", err, output)
	// }

	// 拿到系统输出
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 打印输出
	fmt.Printf("输出是：%s, 输出完毕", out.String())
}

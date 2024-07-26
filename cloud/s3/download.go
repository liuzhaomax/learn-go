package s3

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ledongthuc/pdf"
	"io"
	"net/http"
)

func download(c *gin.Context, urlStr string, fileName string) {
	// 下载文件
	response, err := http.Get(urlStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to download file: %s", err)})
		return
	}
	defer response.Body.Close()

	// 设置响应头，以便浏览器可以正确处理文件下载
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.Header("Content-Type", response.Header.Get("Content-Type"))

	// 将文件内容写入响应
	if _, err := io.Copy(c.Writer, response.Body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to write file to response: %s", err)})
		return
	}
}

func downloadAndConvertToTxt(urlStr string) (string, error) {
	// 下载文件
	response, err := http.Get(urlStr)
	if err != nil {
		return "", fmt.Errorf("failed to download file: %s", err)
	}
	defer response.Body.Close()

	text, err := extractTextFromPdfFromResponse(response)
	if err != nil {
		return "", err
	}
	return text, nil
}

func extractTextFromPdfFromResponse(resp *http.Response) (string, error) {
	// 读取响应体到内存中
	var buf bytes.Buffer
	_, err := io.Copy(&buf, resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %s", err)
	}

	// 使用pdf库解析内存中的PDF数据，保存到临时文件并读取
	reader, err := pdf.NewReader(bytes.NewReader(buf.Bytes()), int64(buf.Len()))
	if err != nil {
		return "", fmt.Errorf("failed to create PDF reader: %s", err)
	}

	var textBuffer bytes.Buffer

	// 遍历每一页
	numPages := reader.NumPage()
	for pageNum := 1; pageNum <= numPages; pageNum++ {
		page := reader.Page(pageNum)
		if &page == nil {
			return "", fmt.Errorf("page %d is nil", pageNum)
		}

		text, err := page.GetPlainText(nil)
		if err != nil {
			return "", err
		}

		textBuffer.WriteString(text)
		textBuffer.WriteString("\n")
	}

	return textBuffer.String(), nil
}

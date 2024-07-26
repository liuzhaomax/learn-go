package pdf

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/ledongthuc/pdf"
)

func ExtractTextFromPdfFromOS(pdfPath string) (string, error) {
	// 打开PDF文件
	f, r, err := pdf.Open(pdfPath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var textBuffer bytes.Buffer

	// 遍历每一页
	numPages := r.NumPage()
	for pageNum := 1; pageNum <= numPages; pageNum++ {
		page := r.Page(pageNum)
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

func ExtractTextFromPdfFromResponse(resp *http.Response) (string, error) {
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

func TestPdfToTxt(t *testing.T) {
	pdfPath := "./test.pdf"
	text, err := ExtractTextFromPdfFromOS(pdfPath)
	if err != nil {
		log.Fatalf("Error extracting text from PDF: %v", err)
	}
	fmt.Println(text)
}

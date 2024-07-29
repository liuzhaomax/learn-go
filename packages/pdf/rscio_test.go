package pdf

import (
	"fmt"
	"rsc.io/pdf"
	"strings"
	"testing"
)

func TestRscioPDF(t *testing.T) {
	// 要读取的PDF文件路径
	filePath := "pdf17.pdf"

	r, err := pdf.Open(filePath)
	if err != nil {
		t.Fatalf("无法打开PDF文件: %v", err)
	}

	var textBuilder strings.Builder
	numPages := r.NumPage()
	for i := 1; i <= numPages; i++ {
		page := r.Page(i)
		if &page == nil {
			t.Fatalf("无法获取第 %d 页", i)
		}
		text, err := extractPageText(&page)
		if err != nil {
			t.Fatalf("提取第 %d 页内容失败: %v", i, err)
		}
		textBuilder.WriteString(text)
		textBuilder.WriteString("\n")
	}

	fmt.Println(textBuilder.String())
}

func extractPageText(page *pdf.Page) (string, error) {
	var sb strings.Builder

	// 获取内容流
	content := page.Content()

	// 遍历内容流并提取文本
	texts := content.Text
	for _, text := range texts {
		sb.WriteString(text.S) // text.S 是实际的字符串内容
	}

	return sb.String(), nil
}

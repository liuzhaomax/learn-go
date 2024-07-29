package pdf

import (
	"fmt"
	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
	"log"
	"os"
	"strings"
	"testing"
)

func TestUnidoc(t *testing.T) {
	// 要读取的PDF文件路径
	filePath := "pdf17.pdf"

	// 打开PDF文件
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("无法打开文件: %v", err)
	}
	defer file.Close()

	// 加载PDF文档
	pdfReader, err := model.NewPdfReader(file)
	if err != nil {
		log.Fatalf("加载PDF文档失败: %v", err)
	}

	// 获取PDF文件的页数
	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		log.Fatalf("无法获取页数: %v", err)
	}

	// 创建一个字符串构建器来存储PDF内容
	var sb strings.Builder

	// 逐页提取内容
	for i := 1; i <= numPages; i++ {
		page, err := pdfReader.GetPage(i)
		if err != nil {
			log.Fatalf("无法获取第 %d 页: %v", i, err)
		}

		ex, err := extractor.New(page)
		if err != nil {
			log.Fatalf("无法创建提取器: %v", err)
		}

		text, err := ex.ExtractText()
		if err != nil {
			log.Fatalf("无法提取文本: %v", err)
		}

		sb.WriteString(text)
		sb.WriteString("\n")
	}

	// 打印PDF内容
	fmt.Println(sb.String())
}

package pdf

import (
	"fmt"
	"github.com/unidoc/unipdf/v3/extractor"
	"log"
	"os"
	"testing"

	"github.com/unidoc/unipdf/v3/model"
)

func TestUnipdf(t *testing.T) {
	// 要读取的PDF文件路径
	filePath := "./pdf17.pdf"

	// 打开PDF文件
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("无法打开PDF文件: %v", err)
	}
	defer f.Close()

	// 读取PDF文件
	pdfReader, err := model.NewPdfReader(f)
	if err != nil {
		log.Fatalf("无法读取PDF文件: %v", err)
	}

	// 获取页数
	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		log.Fatalf("无法获取PDF页数: %v", err)
	}

	// 打印PDF页面数
	for i := 1; i <= numPages; i++ {
		page, err := pdfReader.GetPage(i)
		if err != nil {
			log.Fatalf("无法获取第 %d 页: %v", i, err)
		}

		// 创建一个新的文本提取器
		ex, err := extractor.New(page)
		if err != nil {
			log.Fatalf("无法创建文本提取器: %v", err)
		}

		// 提取文本
		text, err := ex.ExtractText()
		if err != nil {
			log.Fatalf("无法提取第 %d 页的文本: %v", i, err)
		}

		fmt.Printf("第 %d 页内容:\n%s\n", i, text)
	}
}

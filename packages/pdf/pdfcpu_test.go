package pdf

import (
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"log"
	"os"
	"testing"
)

func TestPdfcpu(t *testing.T) {
	// 要读取的PDF文件路径
	filePath := "./pdf17.pdf"

	// 确保文件存在并且可以读取
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("无法打开文件: %v", err)
	}
	defer file.Close()

	// 使用pdfcpu读取PDF文件并检查有效性
	ctx, err := api.ReadContext(file, &model.Configuration{})
	if err != nil {
		log.Fatalf("无法读取PDF文件: %v", err)
	}

	fmt.Printf("PDF文件有效: %v\n", ctx)
}

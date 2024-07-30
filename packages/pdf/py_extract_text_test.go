package pdf

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"testing"
)

func TestPyExtractText(t *testing.T) {
	pdfPath := "pdf17.pdf"
	scriptPath := "extract_text.py"

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("python", scriptPath, pdfPath)
	} else {
		cmd = exec.Command("python3", scriptPath, pdfPath)
	}
	// 构建命令并输出
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("执行命令时出错: %v", err)
	}

	// 打印输出
	fmt.Println(string(output))
}

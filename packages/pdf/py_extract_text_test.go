package pdf

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
	"testing"
)

func TestPyExtractText(t *testing.T) {
	// pdfPath := "D:\\workspace\\max\\learn-go\\packages\\pdf\\pdf17.pdf" // 绝对路径也可
	pdfPath := "pdf17.pdf"
	scriptPath := "extract_text.py"

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("python", scriptPath, pdfPath)
	} else {
		cmd = exec.Command("python3", scriptPath, pdfPath)
	}

	// 运行py文件，但go不拿系统输出
	// output, err := cmd.CombinedOutput()
	// if err != nil {
	// 	log.Fatalf("执行命令时出错: %v\n输出: %s", err, output)
	// }

	// 拿到系统输出
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 打印输出
	fmt.Printf("输出是：%s, 输出完毕", out.String())
}

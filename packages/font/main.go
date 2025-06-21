package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// pip install fonttools
// pyftsubset input.ttf --text="ABC" --output-file=output.ttf

func main() {
	// 打印当前工作目录
	wd, _ := os.Getwd()
	log.Println("当前工作目录:", wd)

	// 1. 检查文件
	inputPath := filepath.Join(wd, "packages/font/input.ttf")
	if _, err := os.Stat(inputPath); err != nil {
		log.Fatalf("错误: %v\n文件路径: %s", err, inputPath)
	}
	log.Println("输入文件路径正确:", inputPath)

	outputPath := filepath.Join(wd, "packages/font/output.ttf")

	// 2. 执行命令
	cmd := exec.Command("pyftsubset", inputPath, "--text=南开钟楼血染无说书人版辅助器平民邪恶阵营胜利刘昭liuzhaomax@163.comLIUZHAOMAX", fmt.Sprintf("--output-file=%s", outputPath))
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("错误: %v\n输出: %s", err, string(output))
	}

	log.Println("成功生成 output.ttf")
}

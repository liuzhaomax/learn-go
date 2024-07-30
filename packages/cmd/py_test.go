package cmd

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestPython(t *testing.T) {
	// Python脚本文件名
	script := "hello.py"
	// 要传递给Python脚本的参数
	arg := "World"

	// 构建命令
	cmd := exec.Command("python", script, arg)

	// 获取命令输出
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 打印输出
	fmt.Println(string(output))
}

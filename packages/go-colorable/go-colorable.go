package main

import (
	"bytes"
	"github.com/dimiro1/banner"
	"github.com/mattn/go-colorable"
)

func main() {
	isEnabled := true
	isColorEnabled := true
	//banner.Init(os.Stdout, isEnabled, isColorEnabled, bytes.NewBufferString("My Custom Banner"))
	banner.Init(colorable.NewColorableStdout(), isEnabled, isColorEnabled, bytes.NewBufferString("My Custom Banner"))
	// 没区别啊？？？
}

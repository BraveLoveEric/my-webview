package main

import (
	"github.com/BraveLoveEric/my-webview/webview"
)

const (
	windowWidth  = 1920
	windowHeight = 1080
)

func main() {
	w := webview.New(webview.Settings{
		Width:     windowWidth,
		Height:    windowHeight,
		Title:     "Simple window demo",
		URL:       "https://www.baidu.com",
		Resizable: true,
	})
	defer w.Exit()
	w.Run()
}

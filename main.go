package main

import "github.com/zserge/webview"

func main() {
	// Open wikipedia in a 800x600 resizable window
	webview.Open("Minimal webview example",
		"http://192.168.39.152:8080/", 800, 600, true)
}

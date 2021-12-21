package main

import (
	"fmt"
	"github.com/webview/webview"
	"net/http"
	"os"
	"path"
)

var (
	AppPort      = "8998"
	ApiPort      = "8999"
	AppDirectory = "ui"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic("Cannot get current working directory")
	}
	AppDirectory = path.Join(cwd, AppDirectory)
	fmt.Println(AppDirectory)

	http.Handle("/", http.FileServer(http.Dir(AppDirectory)))
	go func() {
		_ = http.ListenAndServe(":"+AppPort, nil)
	}()

	debug := true
	wv := webview.New(debug)
	defer wv.Destroy()
	wv.Navigate("http://localhost:" + AppPort)
	wv.SetTitle("Ulaa Manager Setup")
	wv.SetSize(600, 400, webview.HintFixed)
	wv.Run()

}

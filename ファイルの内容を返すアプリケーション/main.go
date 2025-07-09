package main

import (
	"log"
	"net/http"
)

func main() {
	// パス/にHTTPリクエストがあった時、http.FileServer()が返す関数を渡す。
	// 引数に指定したhttp.Dir("static")により、staticという名前のディレクトリ配下からファイルを探して返す。
	// staticフォルダの中にはhello.htmlがあるので、その内容が返される。
	http.Handle("/", http.FileServer(http.Dir("static")))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("failed to start : ", err)
	}
}

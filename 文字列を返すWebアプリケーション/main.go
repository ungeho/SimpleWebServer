package main

// net/httpパッケージ は HTTP通信に関する様々な処理を提供する。
import (
	"fmt"
	"log"
	"net/http"
)

// パス/にHTTPリクエストを受けたらhello関数の処理が実行され、"Hello, Web application"の文字列を出力する。
func hello(w http.ResponseWriter, r *http.Request) {
	// Fprintfは指定した出力先(http.ResponseWriter)に文字列を書き込む。
	fmt.Fprintf(w, "Hello, Web application!")
}

func main() {
	// http.HandleFuncはHTTPリクエストに対する処理
	// HTTPサーバーを起動する前の準備としてhttp.HandleFunc()関数に登録する。
	// /というパスにHTTPリクエストが届いたら、hello関数の処理を実行するように登録。
	http.HandleFunc("/", hello)
	// ListenAndServeはHTTPサーバーを起動する関数
	// 8080番ポートで待ち受ける。
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("failed to start : ", err)
	}
}

package main

import (
	"html/template"
	"log"
	"net/http"
)

// .todo項目を格納する為の文字列型のスライスを定義
var todoList []string

func handleTodo(w http.ResponseWriter, r *http.Request) {
	// Goのテンプレート機能html/templateパッケージを使用
	// .todoList変数に保持するTodo項目をユーザーの操作に応じて変化。
	// HTMLの動的な生成が必要になる為、テンプレート機能を使用する。
	// .todo.htmlというファイルで用意されたテンプレートに
	// .todoList変数が保持するToDo項目を組み合わせてHTMLを生成し、レスポンスとして返す。
	// template.ParseFiles()関数は、todo.htmlという名前で用意された
	// テンプレートファイルを読み込み、テンプレートとして解析する。
	// 解析されたテンプレートは変数tに保存される。
	// Execute()関数には、HTTPレスポンスを出力するhttp.ResponseWriterとテンプレートに埋め込むtodoList変数を与え
	// 最終的なHTMLを生成してレスポンスとして返す。

	// ブランク識別子でエラーを無視
	// t, _ := template.ParseFiles("templates/todo.html")

	t, err := template.ParseFiles("templates/todo.html")
	if err != nil {
		http.Error(w, "テンプレート読み込みエラー: "+err.Error(), http.StatusInternalServerError)
		log.Println("テンプレート読み込み失敗:", err)
		return
	}
	t.Execute(w, todoList)
}

func main() {
	todoList = append(todoList, "顔を洗う", "朝食を食べる", "歯を磨く")

	// /static/に変更する事で、/static/~というパス配下を
	// 静的コンテンツ(html,css,javascript,画像など)に割り当てる。
	// StripePrefix()は、http.FileServer()に渡すパスを調整する為の関数
	// 何もしないと、`http://localhost:8080/static/todolist.css`というパスからファイルを取得した時
	// static/static/todolist.cssというパスからファイルを取得しようとしてしまう事になる。
	// 最初のstaticはhttp.Dir()関数で指定したもので、ファイルシステム上で静的ファイルを探す為の基点
	// 2番目のパスはURLに指定されていたもの。
	// その為、http.StripePrefix()関数で、http.FileServer()関数に渡すパスからstaticという文字列を取り除いている。
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// パス/todoが指定されたら、handleTodo()関数を呼び出す。
	http.HandleFunc("/todo", handleTodo)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("failed to start : ", err)
	}
}

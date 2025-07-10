package main

import (
	"html/template"
	"log"
	"net/http"
)

// .todo項目を格納する為の文字列型のスライスを定義
var todoList []string

func handleTodo(w http.ResponseWriter, r *http.Request) {
	// テンプレートとしてtemplates/todo.htmlを解析
	t, err := template.ParseFiles("templates/todo.html")
	if err != nil {
		http.Error(w, "テンプレート読み込みエラー: "+err.Error(), http.StatusInternalServerError)
		log.Println("テンプレート読み込み失敗:", err)
		return
	}
	// テンプレートに埋め込むtodoList変数の値を元に、最終的なHTMLを生成して返す。
	t.Execute(w, todoList)
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	// addパスへのHTTPリクエストを解析
	r.ParseForm()
	// .todoという名前のパラメーターから値を取り出してtodo変数に格納。
	todo := r.Form.Get("todo")
	// .todoListに追加
	todoList = append(todoList, todo)
	// handleTodo()関数を呼び出して、todoList変数の値を元にHTMLを再度生成する。
	// handleTodo(w, r)

	// しかし、handleTodo()関数でHTMLを再生成しても、パスは/addのままなので更新すると再度todoが追加されてしまう。
	// そこで、ステータスコード303(SeeOther)を返す事で、/todoへリクエストし直すようにブラウザに指示を行う。
	// （HTTPレスポンスでブラウザに300番台のステータスコードを返す事をリダイレクトと呼ぶ）
	// /todoであれば、更新しても再度追加される問題がなくなる。
	http.Redirect(w, r, "/todo", 303)
}

func main() {
	todoList = append(todoList, "顔を洗う", "朝食を食べる", "歯を磨く")

	// http.Dir("static")でファイスシステム上でstaticフォルダ内のファイルを参照するように指定している。
	// またhttp.Handleでは、/static/~というパス配下を静的コンテンツに割り当てる為に、/static/を指定している。
	// これにより、urlで/static/todolist.cssにアクセスした時、FileServerはstatic/static/todolist.cssを参照してしまう。
	// そこで、/static/が指定された時は、/static/を取り除いた状態で（つまりmain.goのある場所から見たstaticフォルダの配下）を探す
	// .todo.htmlで指定された href="/static/todo.css" がmain.goからみたstatic/todo.cssになるように
	// href="../static/todo.css"じゃないの？って思っていた部分は、ここで/static/が指定された時の挙動が指示されていた？ことで解決したかもしれない
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// パス/todoが指定されたら、handleTodo()関数を呼び出す。
	http.HandleFunc("/todo", handleTodo)

	// パス/todoが指定されたら、handleAdd()関数を呼び出す。
	// handleAdd()関数内の最後で、handleTodo()関数を呼び出しているので、追加処理後にHTMLが再生成される。
	http.HandleFunc("/add", handleAdd)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("failed to start : ", err)
	}

}

# Simple Web Server

Web技術入門の写経

## サーバーの起動

`main.go`が存在するフォルダで

`go run main.go`

## クライアント側からHTTPリクエストを送信

### 文字列を返すアプリケーションの場合

`curl http://localhost:8080`

### ファイルの内容を返すアプリケーションの場合

ブラウザで

`http://localhost:8080/hello.html`

### Tiny ToDo01 の場合

ブラウザで

`http://localhost:8080/todo`
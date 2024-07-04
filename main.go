package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	responseMessage := "Hello, World!"

	// 1. ハンドラの宣言
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		// ハンドラの処理
		io.WriteString(w, responseMessage) // wにレスポンスを書き込む
	}

	postArticleHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Article posted!")
	}

	articleListHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Article list!")
	}

	articleDetailHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Article detail!")
	}

	postNiceHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Nice posted!")
	}

	postCommentHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Comment posted!")
	}

	// ハンドラの登録
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/article", postArticleHandler)
	http.HandleFunc("/article/list", articleListHandler)
	http.HandleFunc("/article/1", articleDetailHandler)
	http.HandleFunc("/nice", postNiceHandler)
	http.HandleFunc("/comment", postCommentHandler)

	// サーバー起動時のログ出力
	log.Println("Listening on port 8080")

	//ListenAndServe関数：指定したポートでHTTPリクエストを待ち受ける
	log.Fatal(http.ListenAndServe(":8080", nil))
	// 2行で書くパターン
	// err := http.ListenAndServe(":8080", nil)
	// log.Fatal(err)
}

// インターフェース型

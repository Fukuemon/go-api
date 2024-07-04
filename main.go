package main

import (
	"log"
	"net/http"

	"github.com/Fukuemon/go-api/handlers"
)

func main() {

	// ハンドラの登録
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/article/list", handlers.ArticleListHandler)
	http.HandleFunc("/article/1", handlers.ArticleDetailHandler)
	http.HandleFunc("/nice", handlers.PostNiceHandler)
	http.HandleFunc("/comment", handlers.PostCommentHandler)

	// サーバー起動時のログ出力
	log.Println("Listening on port 8080")

	//ListenAndServe関数：指定したポートでHTTPリクエストを待ち受ける
	log.Fatal(http.ListenAndServe(":8080", nil))
	// 2行で書くパターン
	// err := http.ListenAndServe(":8080", nil)
	// log.Fatal(err)
}

// インターフェース型

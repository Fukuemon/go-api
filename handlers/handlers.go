package handlers

import (
	"fmt"
	"io"
	"net/http"
)

var responseMessage string = "Hello, World!"

// 1. ハンドラの宣言
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	// ハンドラの処理
	io.WriteString(w, responseMessage) // wにレスポンスを書き込む
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Article posted!")
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Article list!")
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleId := 1
	resString := fmt.Sprintf("Article detail! articleId: %d", articleId)
	io.WriteString(w, resString)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Nice posted!")
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Comment posted!")
}

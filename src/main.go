package main

import (
	"log"
	"net/http"

	"github.com/tamurakeito/tamurakeito-api-gateway/src/infrastructure"
	"github.com/tamurakeito/tamurakeito-api-gateway/src/presentation"
	"github.com/tamurakeito/tamurakeito-api-gateway/src/usecase"
)

func main() {
	// SqlHandlerのインスタンスを生成
	sqlHandler := infrastructure.NewSqlHandler()
	defer sqlHandler.Conn.Close()

	// リポジトリのインスタンスを生成
	repo := infrastructure.NewMySQLProxyConfigRepository(*sqlHandler)

	// サービスを使用してプロキシの設定
	service := usecase.NewProxyUsecase(repo)
	presentation.RegisterProxies(service)

	log.Println("Starting server on :80")
	log.Fatal(http.ListenAndServe(":80", nil))
}

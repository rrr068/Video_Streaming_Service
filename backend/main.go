package main

import (
    "github.com/gin-gonic/gin"
		"github.com/example/gin-backend/infra/db"
		"github.com/example/gin-backend/infra/router"
)

func main() {
	// DBの接続
	db.InitDB()
	// r := gin.Default()
	// routerの接続
	r := router.NewRouter()
	// サーバー接続
	r.Run(":8080")
}
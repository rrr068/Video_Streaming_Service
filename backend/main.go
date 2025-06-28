package main

import (
    "github.com/gin-gonic/gin"
		"github.com/example/gin-backend/database"
)

func main() {
	database.InitDB()

	r := gin.Default()

	// エンドポイント
	r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong,pong,pong"})
	})

	// Ginサーバー[localhost:8080/ping]にアクセスするとpongが帰ってくる
	r.Run(":8080")
}
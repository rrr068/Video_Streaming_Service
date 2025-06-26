package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // エンドポイント
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })

		// Ginサーバー[localhost:8080/ping]にアクセスするとpongが帰ってくる
    r.Run(":8080")
}
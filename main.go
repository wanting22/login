package main

import (
	"auth"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/login", auth.LoginPage)
	r.POST("/login", auth.LoginAuth)
	//讀取html的所在位置 ＊是代表目錄所有檔案 根目錄是什麼??html渲染是什麼
	r.LoadHTMLGlob("./html/*")
	//設定靜態資源的讀取
	//https://matthung0807.blogspot.com/2021/07/go-gin-static-files.html
	r.Static("/css", "./css")
	r.Run(":8000")
}

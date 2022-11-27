package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userdata = make(map[string]string)

// 在main之前執行
func init() {
	userdata["test"] = "sandy123"
}

func checkUserIsExist(username string) bool {
	_, exist := userdata[username]
	return exist
}

func checkPassword(username string, password string) bool {
	if userdata[username] == password {
		return true
	} else {
		return false
	}
}

func auth(username string, password string) bool {
	if exist := checkUserIsExist(username); exist {
		return checkPassword(username, password)
	} else {
		return false
	}
}

func loginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func loginAuth(c *gin.Context) {
	var (
		username string
		password string
	)
	if in, exist := c.GetPostForm("username"); exist && in != "" {
		username = in
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"error": errors.New("請輸入使用者名稱")})
		return
	}
	if in, exist := c.GetPostForm("password"); exist && in != "" {
		password = in
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"error": errors.New("請輸入密碼")})
		return
	}
	if in := auth(username, password); in {
		c.HTML(http.StatusOK, "login.html", gin.H{"success": "登入成功"})
	} else {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"error": "err"})
	}

}

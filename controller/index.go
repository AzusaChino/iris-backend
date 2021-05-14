package controller

import (
	"iris/common"
	"iris/model"
	"iris/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * URL控制器
 */
type Handler interface {
	Method() string
	Path() string
	HandlerFunc(ctx *gin.Context)
}

type LoginHandler struct{}

func (l *LoginHandler) Method() string {
	return http.MethodPost
}

func (l *LoginHandler) Path() string {
	return "/login"
}

func (l *LoginHandler) HandlerFunc(c *gin.Context) {
	json := model.User{}
	c.BindJSON(&json)
	db, cleanFunc, err := utils.GetDb()
	defer cleanFunc()
	if err != nil {
		log.Fatal(err)
	}
	var user model.User
	db.Where("username = ? AND password = ?", json.Username, json.Password).Find(&user)
	if &user == nil {
		c.JSON(http.StatusOK, common.Error(1001, "登陆失败"))
		return
	}

	c.JSON(http.StatusOK, common.Ok(user))
}

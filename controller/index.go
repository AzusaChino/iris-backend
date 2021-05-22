package controller

import (
	"iris/common"
	"iris/model"
	"iris/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	json := model.User{}

	// 通过 `json:""`，可以使用BindJSON获取参数
	c.BindJSON(&json)
	// 若通过`form:""`, 可以使用表单以及Bind获取参数
	db, cleanFunc, err := utils.GetDb()
	defer cleanFunc()
	if err != nil {
		log.Fatal(err)
	}
	var user model.User
	result := db.Where("username = ? AND password = ?", json.Username, json.Password).Find(&user)
	if result.Error != nil {
		c.JSON(http.StatusOK, common.Error(1001, "登陆失败"))
		return
	}

	c.JSON(http.StatusOK, common.Ok(user))
}

package controller

import (
	"github.com/azusachino/iris-content/common"
	"github.com/azusachino/iris-content/model"
	"github.com/azusachino/iris-content/pkg/orm"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	json := model.LoginParam{}

	// 通过 `json:""`，可以使用BindJSON获取参数
	c.BindJSON(&json)
	// 若通过`form:""`, 可以使用表单以及Bind获取参数
	db := orm.GetDb()

	var user model.User
	result := db.Where("username = ? AND password = ?", json.Username, json.Password).Find(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, common.Error(1001, "登陆失败"))
		return
	}

	c.JSON(http.StatusOK, common.Ok(user))
}

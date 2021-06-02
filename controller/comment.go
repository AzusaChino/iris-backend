package controller

import (
	"iris/common"
	"iris/model"
	"iris/pkg/orm"
	"net/http"

	"github.com/gin-gonic/gin"
)

func QueryComment(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, queryComment(id))
}

func queryComment(articleId string) common.RestResponse {
	db := orm.GetDb()

	var comments []model.Comment
	// search comments level 1
	result := db.Where("article_id = ? and pid = 0", articleId).Find(&comments)

	if result.Error != nil {
		return common.Error(1006, "获取文章评论失败")
	}

	var slice = []map[string]interface{}{}
	for _, c := range comments {
		var m = make(map[string]interface{})
		var user model.User
		db.Select("nickname", "avatar").Where("username = ?", c.CreateUser).Find(&user)
		m["comment"] = c
		m["user"] = user
		slice = append(slice, m)
	}

	return common.Ok(slice)
}

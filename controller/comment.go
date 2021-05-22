package controller

import (
	"iris/common"
	"iris/model"
	"iris/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func QueryComment(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, queryComment(id))
}

func queryComment(articleId string) common.RestResponse {
	db, clean, err := utils.GetDb()
	defer clean()
	if err != nil {
		log.Fatal(err)
	}
	var comments []model.Comment
	result := db.Where("article_id = ?", articleId).Find(&comments)
	if result.Error != nil {

		return common.Ok(comments)
	} else {
		return common.Error(1006, "获取文章评论失败")
	}
}

package controller

import (
	"iris/common"
	"iris/model"
	"iris/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentController struct{}

func (a *CommentController) Method() string {
	return http.MethodGet
}

func (a *CommentController) Path() string {
	return "/article/:id/comment"
}

func (a *CommentController) HandlerFunc(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, queryComments(id))
}

func queryComments(articleId string) common.RestResponse {
	db, clean, err := utils.GetDb()
	defer clean()
	if err != nil {
		log.Fatal(err)
	}
	var comments []model.Comment
	db.Where("article_id = ?", articleId).Find(&comments)

	return common.Ok(comments)
}

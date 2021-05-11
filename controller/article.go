package controller

import (
	"iris/common"
	"iris/model"
	"iris/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
}

func (a *ArticleController) Method() string {
	return http.MethodGet
}

func (a *ArticleController) Path() string {
	return "/articles"
}

func (a *ArticleController) HandlerFunc(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, queryArticles())
}

func queryArticles() common.RestResponse {
	db, cleanFunc, err := utils.GetDb()
	defer cleanFunc()
	if err != nil {
		log.Fatal(err)
	}

	var articles []model.Article
	db.Find(&articles)

	return common.Ok(articles)
}

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

type ArticleDetailController struct {
}

func (a *ArticleDetailController) Method() string {
	return http.MethodGet
}

func (a *ArticleDetailController) Path() string {
	return "/article/:id"
}

func (a *ArticleDetailController) HandlerFunc(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, queryArticle(id))
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

func queryArticle(id string) common.RestResponse {
	db, cleanFunc, err := utils.GetDb()
	defer cleanFunc()
	if err != nil {
		log.Fatal(err)
	}
	var article model.Article
	db.Where("id = ?", id).Find(&article)
	return common.Ok(article)
}

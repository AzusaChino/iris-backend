package controller

import (
	"iris/common"
	"iris/model"
	"iris/pkg/orm"
	"net/http"

	"github.com/gin-gonic/gin"
)

func QueryArticle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, queryArticles())
}

func QueryArticleDetail(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, queryArticleDetail(id))
}

// TODO 分页查询
func queryArticles() common.RestResponse {
	db := orm.GetDb()

	var articles []model.Article
	result := db.Find(&articles)

	if result.Error != nil {
		return common.Ok(articles)
	} else {
		return common.Error(1004, "没有任何文章")
	}
}

func queryArticleDetail(id string) common.RestResponse {
	db := orm.GetDb()

	var article model.Article
	result := db.Where("id = ?", id).Find(&article)
	if result.Error != nil {
		return common.Ok(article)
	} else {
		return common.Error(1005, "未找到文章详情")
	}
}

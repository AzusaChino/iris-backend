package controller

import (
	"github.com/azusachino/iris/common"
	"github.com/azusachino/iris/model"
	"github.com/azusachino/iris/pkg/orm"
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
		return common.Error(1004, "没有任何文章")
	}
	return common.Ok(articles)
}

func queryArticleDetail(aid string) common.RestResponse {
	db := orm.GetDb()

	var articleDetail model.ArticleDetail
	result := db.Where("article_id = ?", aid).Find(&articleDetail)
	if result.Error != nil {
		return common.Error(1005, "未找到文章详情")
	}
	return common.Ok(articleDetail)
}

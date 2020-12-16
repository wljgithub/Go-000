package api

import (
	"github.com/gin-gonic/gin"
	"week4/api/article"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	api:=g.Group("/api")
	{
		api.GET("/getArticle",article.GetArticle)
	}

	return g
}

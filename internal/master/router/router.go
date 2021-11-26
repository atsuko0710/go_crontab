package router

import (
	"master/internal/master/router/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	
	// 添加中间件
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(mw...)

	// 404 处理
	g.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "该路由不存在"})
	})

	// TODO 增加路由

	return g
}
package master

import (
	"master/internal/master/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitApiServer() (err error) {
	g := gin.New()

	// 配置路由
	router.Load(
		g,
	)

	http.ListenAndServe("config.Conf.Addr", g).Error()
	return
}
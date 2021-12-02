package master

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitApiServer() (err error) {
	
	gin.SetMode(viper.GetString("runmode"))
	g := gin.New()

	// 配置路由
	Load(
		g,
	)

	http.ListenAndServe(viper.GetString("addr"), g).Error()
	return
}
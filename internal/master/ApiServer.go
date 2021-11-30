package master

import (
	"master/internal/master/router"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitApiServer() (err error) {
	if err = InitConfig(); err != nil {
		return err
	}
	
	gin.SetMode(viper.GetString("runmode"))
	g := gin.New()

	// 配置路由
	router.Load(
		g,
	)

	http.ListenAndServe(":8888", g).Error()
	return
}
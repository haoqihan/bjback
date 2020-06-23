package initialiaze

import (
	"bjback/router"
	"github.com/gin-gonic/gin"
)

// 初始化总路由
func Routers() *gin.Engine {
	var Router = gin.Default()
	ApiGroup := Router.Group("")
	router.InitUserRouter(ApiGroup) // 注册用户路由
	return Router
}

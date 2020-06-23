package router

import "github.com/gin-gonic/gin"

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("changePassword")   // 修改密码
		UserRouter.POST("uploadHeaderImg")  // 上传头像
		UserRouter.POST("getUserList")      // 分页获取用户列表
		UserRouter.POST("setUserAuthority") // 设置用户权限
		UserRouter.DELETE("deleteUser")     // 删除用户
	}
}

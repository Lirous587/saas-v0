package user

import (
	"github.com/gin-gonic/gin"
	"saas/internal/common/middleware/auth"
	"saas/internal/user/handler"
)

func RegisterV1(r *gin.RouterGroup, handler *handler.HttpHandler) func() {
	userGroup := r.Group("/v1/user")

	{
		// 登录相关路由
		userGroup.POST("/auth/github", handler.GithubAuth)

		// 令牌管理
		userGroup.POST("/refresh_token", handler.RefreshToken)

		// 需要token的路由
		protected := userGroup.Group("")
		protected.Use(auth.Validate())
		{
			protected.POST("/auth", handler.ValidateAuth)
			protected.GET("/profile", handler.GetProfile)
		}
	}
	return nil
}

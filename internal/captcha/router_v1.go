package captcha

import (
	"github.com/gin-gonic/gin"
	"saas/internal/captcha/handler"
	"saas/internal/common/middleware/auth"
	"saas/internal/common/reskit/response"
)

func RegisterV1(r *gin.RouterGroup, handler *handler.HttpHandler) func() {
	g := r.Group("/v1/captcha")
	{
		g.POST("", handler.Gen)
		//验证端点
		g.POST("/verify", handler.Verify(), func(ctx *gin.Context) {
			response.Success(ctx)
		})

		// 测试路由：生成验证码并返回图片+验证答案
		g.POST("/with-answer", auth.Validate(), handler.GenWithAnswer)
	}

	return nil
}

package bootstrap

import (
	"Gohub/routes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// SetupRoute 路由初始化
func SetupRoute(r *gin.Engine) {

	// 注册全局中间件
	registerGlobalMiddleWare(r)

	// 注册 API 路由器
	routes.RegisterAPIRoutes(r)

	// 配置 404 路由
	setup404Handler(r)
}

// 注册全局中间件
func registerGlobalMiddleWare(r *gin.Engine) {
	r.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}

func setup404Handler(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 的话
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 默认返回 JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}

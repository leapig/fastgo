package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(r *gin.RouterGroup) {
	// 基础路由
	Core(r.Group("core"))

	// gin路由在此定义
	r.GET("hello", func(context *gin.Context) {
		content := "<script>location.href = 'https://www.baidu.com'</script>"
		context.Data(http.StatusOK, "text/html", []byte(content))
	})

	r.GET("vK3zk1LUQg.txt", func(context *gin.Context) {
		context.String(200, "358c7d5014b73a775083c15f21cded09")
	})
}

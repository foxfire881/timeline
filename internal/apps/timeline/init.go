package timeline

import (
	"aio/internal/apps/timeline/controller"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	web := r.Group("/web/timeline")
	api := r.Group("/api/timeline")
	controller.RouterGroup(api)
	controller.RouterSubgroup(api)
	controller.RouterItem(api, web)
}

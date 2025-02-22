package info

import (
	"github.com/gin-gonic/gin"

	"github.com/darianJmy/fortis-services/cmd/app/options"
)

type infoRouter struct {
}

func NewRouter(o *options.ServerRunOptions) {
	router := &infoRouter{
	
	}

	router.initRoutes(o.HttpEngine)
}

func (ir *infoRouter) initRoutes(httpEngine *gin.Engine) {
	infoRoute := httpEngine.Group("/info")
	{
		infoRoute.POST("", ir.createInfo)
		infoRoute.PUT("/:infoId", ir.updateInfo)
		infoRoute.DELETE("/:infoId", ir.deleteInfo)
		infoRoute.GET("/:infoId", ir.getInfo)
		infoRoute.GET("", ir.listInfos)
	}
}

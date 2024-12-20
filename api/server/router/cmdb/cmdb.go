package cmdb

import (
	"github.com/darianJmy/fortis-services/pkg/controller"
	"github.com/gin-gonic/gin"

	"github.com/darianJmy/fortis-services/cmd/app/options"
)

type cmdbRouter struct {
	control controller.FortisInterface
}

func NewRouter(o *options.ServerRunOptions) {
	router := &cmdbRouter{
		control: o.Control,
	}

	router.cmdbRoutes(o.HttpEngine)
}

func (cm *cmdbRouter) cmdbRoutes(httpEngine *gin.Engine) {
	cmRoute := httpEngine.Group("/cmdb")
	{
		cmRoute.POST("", cm.createCMDB)
		cmRoute.PUT("/:cmdbId", cm.updateCMDB)
		cmRoute.DELETE("/:cmdbId", cm.deleteCMDB)
		cmRoute.GET("/:cmdbId", cm.getCMDB)
		cmRoute.GET("", cm.listCMDBs)
		cmRoute.GET("/resource/count", cm.resourceCount)

	}
}

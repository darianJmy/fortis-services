package cmdb

import (
	"github.com/gin-gonic/gin"

	"github.com/darianJmy/fortis-services/cmd/app/options"
)

type cmdbRouter struct {
}

func NewRouter(o *options.ServerRunOptions) {
	router := &cmdbRouter{}

	router.cmdbRoutes(o.HttpEngine)
}

func (cm *cmdbRouter) cmdbRoutes(httpEngine *gin.Engine) {
	infoRoute := httpEngine.Group("/cmdb")
	{
		infoRoute.POST("", cm.createCMDB)
		infoRoute.PUT("/:cmdbId", cm.updateCMDB)
		infoRoute.DELETE("/:cmdbId", cm.deleteCMDB)
		infoRoute.GET("/:cmdbId", cm.getCMDB)
		infoRoute.GET("", cm.listCMDBs)
	}
}

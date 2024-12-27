package cmdb

import (
	"github.com/gin-gonic/gin"

	"github.com/darianJmy/fortis-services/cmd/app/options"
	"github.com/darianJmy/fortis-services/pkg/controller"
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
		cmRoute.POST("/create/objClassification", cm.createObjClassification)
		cmRoute.POST("/list/objClassification", cm.listObjClassification)
		cmRoute.POST("/create/object", cm.createObject)
		cmRoute.POST("/list/object", cm.listObject)
		cmRoute.POST("/create/objectAttr", cm.createObjectAttr)
		cmRoute.POST("/list/objectAttr", cm.listObjectAttr)
		cmRoute.POST("/create/instance/object/:objectId", cm.createObjectData)
		cmRoute.POST("/list/instance/object/:objectId", cm.listObjectData)
	}
}

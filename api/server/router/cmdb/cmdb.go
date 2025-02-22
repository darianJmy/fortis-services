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
		cmRoute.GET("/list/objClassification", cm.listObjClassification)
		cmRoute.PUT("/update/objClassification/:objClsId", cm.updateObjClassification)
		cmRoute.DELETE("/delete/objClassification/:objClsId", cm.deleteObjClassification)

		cmRoute.POST("/create/object", cm.createObject)
		cmRoute.GET("/list/object/:objClsId", cm.listObject)
		cmRoute.PUT("/update/object/:objId", cm.updateObject)
		cmRoute.DELETE("/delete/object/:objId", cm.deleteObject)

		cmRoute.POST("/create/objectAttr", cm.createObjectAttr)
		cmRoute.GET("/list/objectAttr/:objId", cm.listObjectAttr)
		cmRoute.PUT("/update/objectAttr/:objId/:objAttrId", cm.updateObjectAttr)
		cmRoute.DELETE("/delete/objectAttr/:objId/:objAttrId", cm.deleteObjectAttr)

		cmRoute.POST("/create/object/instance/:objId", cm.createInstData)
		cmRoute.GET("/list/object/instance/:objId", cm.listInstData)
		cmRoute.PUT("/update/object/instance/:objId/:instId", cm.updateInstData)
		cmRoute.DELETE("/delete/object/instance/:objId/:instId", cm.deleteInstData)

		cmRoute.POST("/create/associationType", cm.createAssociationType)
		cmRoute.GET("/list/associationType", cm.listAssociationType)
		cmRoute.DELETE("/delete/associationType/:associationId", cm.deleteAssociationType)

		cmRoute.POST("/create/objAssociation", cm.createObjectAssociation)
		cmRoute.GET("/list/objAssociation/:objId", cm.listObjectAssociation)
		cmRoute.DELETE("/delete/objAssociation/:objAsstId", cm.deleteObjectAssociation)

		cmRoute.POST("/create/instAssociation", cm.createInstAssociation)
		cmRoute.GET("/list/instAssociation/:objId", cm.listInstAssociation)
		cmRoute.DELETE("/delete/instAssociation/:instAsstId", cm.deleteInstAssociation)
	}
}

package cmdb

import (
	"context"
	"github.com/darianJmy/fortis-services/pkg/types"
	"github.com/gin-gonic/gin"

	"github.com/darianJmy/fortis-services/api/server/httputils"
)

func (cm *cmdbRouter) createCMDB(c *gin.Context) {
	r := httputils.NewResponse()

	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) updateCMDB(c *gin.Context) {
	r := httputils.NewResponse()

	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) deleteCMDB(c *gin.Context) {
	r := httputils.NewResponse()

	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) getCMDB(c *gin.Context) {
	r := httputils.NewResponse()

	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) listCMDBs(c *gin.Context) {
	r := httputils.NewResponse()

	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) resourceCount(c *gin.Context) {
	r := httputils.NewResponse()

	resources, err := cm.control.CMDB().GetResource(context.TODO())
	if err != nil {
		httputils.SetFailed(c, r, err)
	}

	r.Result = resources
	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) createObjClassification(c *gin.Context) {
	r := httputils.NewResponse()

	var (
		obj types.ObjClassification
		err error
	)

	if err = c.ShouldBindJSON(&obj); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	if err = cm.control.CMDB().CreateObjClassification(context.TODO(), &obj); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) createObject(c *gin.Context) {
	r := httputils.NewResponse()

	var (
		obj types.ObjectDes
		err error
	)

	if err = c.ShouldBindJSON(&obj); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	if err = cm.control.CMDB().CreateObject(context.TODO(), &obj); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) createObjectAttr(c *gin.Context) {
	r := httputils.NewResponse()

	var (
		obj types.ObjectAttr
		err error
	)

	if err = c.ShouldBindJSON(&obj); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	if err = cm.control.CMDB().CreateObjectAttr(context.TODO(), &obj); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) createObjectData(c *gin.Context) {
	r := httputils.NewResponse()

	var (
		obj map[string]string
		err error
	)

	objId := c.Param("objectId")

	if err = c.ShouldBindJSON(&obj); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	if err = cm.control.CMDB().CreateObjectData(context.TODO(), objId, obj); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	httputils.SetSuccess(c, r)
}

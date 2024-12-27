package cmdb

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/darianJmy/fortis-services/api/server/httputils"
	"github.com/darianJmy/fortis-services/pkg/types"
)

func (cm *cmdbRouter) createObjClassification(c *gin.Context) {
	r := httputils.NewResponse()

	var (
		objCls types.ObjClassification
		err    error
	)

	if err = c.ShouldBindJSON(&objCls); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	result, err := cm.control.CMDB().CreateObjClassification(context.TODO(), &objCls)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	r.Result = result
	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) listObjClassification(c *gin.Context) {
	r := httputils.NewResponse()

	objCls, err := cm.control.CMDB().ListObjClassification(context.TODO())
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	r.Result = objCls
	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) updateObjClassification(c *gin.Context) {
	r := httputils.NewResponse()

	var (
		objCls types.ObjClassification
		err    error
	)

	p := c.Param("objCls")

	if err = c.ShouldBindJSON(&objCls); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	result, err := cm.control.CMDB().UpdateObjClassification(context.TODO(), p, &objCls)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	r.Result = result
	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) deleteObjClassification(c *gin.Context) {
	r := httputils.NewResponse()

	p := c.Param("objCls")

	result, err := cm.control.CMDB().DeleteObjClassification(context.TODO(), p)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	r.Result = result
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

	result, err := cm.control.CMDB().CreateObject(context.TODO(), &obj)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	r.Result = result
	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) listObject(c *gin.Context) {
	r := httputils.NewResponse()

	var (
		obj types.ObjectDes
		err error
	)

	if err = c.ShouldBindJSON(&obj); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	result, err := cm.control.CMDB().ListObject(context.TODO(), obj.ClassificationId)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	r.Result = result
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

	result, err := cm.control.CMDB().CreateObjectAttr(context.TODO(), &obj)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	r.Result = result
	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) listObjectAttr(c *gin.Context) {
	r := httputils.NewResponse()

	var (
		obj types.ObjectAttr
		err error
	)

	if err = c.ShouldBindJSON(&obj); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	result, err := cm.control.CMDB().ListObjectAttr(context.TODO(), obj.ObjectId)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	r.Result = result
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

	result, err := cm.control.CMDB().CreateObjectData(context.TODO(), objId, obj)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	r.Result = result
	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) listObjectData(c *gin.Context) {
	r := httputils.NewResponse()

	objId := c.Param("objectId")

	result, err := cm.control.CMDB().ListObjectData(context.TODO(), objId)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	r.Result = result
	httputils.SetSuccess(c, r)
}

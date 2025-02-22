package cmdb

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/darianJmy/fortis-services/api/server/httputils"
	"github.com/darianJmy/fortis-services/pkg/types"
)

func (cm *cmdbRouter) createObjClassification(c *gin.Context) {
	r := httputils.NewResponse()

	var objCls types.ObjClassification
	if err := c.ShouldBindJSON(&objCls); err != nil {
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

	result, err := cm.control.CMDB().ListObjClassification(context.TODO())
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	r.Result = result
	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) updateObjClassification(c *gin.Context) {
	r := httputils.NewResponse()

	p := c.Param("objClsId")

	var objCls types.ObjClassification
	if err := c.ShouldBindJSON(&objCls); err != nil {
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

	p := c.Param("objClsId")

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

	var obj types.ObjectDes
	if err := c.ShouldBindJSON(&obj); err != nil {
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

	p := c.Param("objClsId")

	result, err := cm.control.CMDB().ListObject(context.TODO(), p)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	r.Result = result
	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) updateObject(c *gin.Context) {
	r := httputils.NewResponse()

	p := c.Param("objId")

	var obj types.ObjectDes
	if err := c.ShouldBindJSON(&obj); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	result, err := cm.control.CMDB().UpdateObject(context.TODO(), p, &obj)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	r.Result = result
	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) deleteObject(c *gin.Context) {
	r := httputils.NewResponse()

	p := c.Param("objId")

	result, err := cm.control.CMDB().DeleteObject(context.TODO(), p)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	r.Result = result
	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) createObjectAttr(c *gin.Context) {
	r := httputils.NewResponse()

	var objAttr types.ObjectAttr
	if err := c.ShouldBindJSON(&objAttr); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	result, err := cm.control.CMDB().CreateObjectAttr(context.TODO(), &objAttr)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	r.Result = result
	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) listObjectAttr(c *gin.Context) {
	r := httputils.NewResponse()

	p := c.Param("objId")

	result, err := cm.control.CMDB().ListObjectAttr(context.TODO(), p)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	r.Result = result
	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) updateObjectAttr(c *gin.Context) {
	r := httputils.NewResponse()

	p1 := c.Param("objId")
	p2 := c.Param("objAttrId")

	var objAttr types.ObjectAttr
	if err := c.ShouldBindJSON(&objAttr); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	result, err := cm.control.CMDB().UpdateObjectAttr(context.TODO(), p1, p2, &objAttr)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	r.Result = result
	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) deleteObjectAttr(c *gin.Context) {
	r := httputils.NewResponse()

	p1 := c.Param("objId")
	p2 := c.Param("objAttrId")

	result, err := cm.control.CMDB().DeleteObjectAttr(context.TODO(), p1, p2)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	r.Result = result
	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) createInstData(c *gin.Context) {
	r := httputils.NewResponse()

	objId := c.Param("objId")

	var inst map[string]string
	if err := c.ShouldBindJSON(&inst); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	result, err := cm.control.CMDB().CreateInstanceData(context.TODO(), objId, inst)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	r.Result = result
	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) listInstData(c *gin.Context) {
	r := httputils.NewResponse()

	objId := c.Param("objId")

	result, err := cm.control.CMDB().ListInstanceData(context.TODO(), objId)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	r.Result = result
	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) updateInstData(c *gin.Context) {
	r := httputils.NewResponse()

	p1 := c.Param("objId")
	p2 := c.Param("instId")

	var inst map[string]string
	if err := c.ShouldBindJSON(&inst); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	result, err := cm.control.CMDB().UpdateInstanceData(context.TODO(), p1, p2, inst)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	r.Result = result
	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) deleteInstData(c *gin.Context) {
	r := httputils.NewResponse()

	objId := c.Param("objId")
	instId := c.Param("instId")

	result, err := cm.control.CMDB().DeleteInstanceData(context.TODO(), objId, instId)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	r.Result = result
	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) createAssociationType(c *gin.Context) {
	r := httputils.NewResponse()

	var asstType types.AssociationType
	if err := c.ShouldBindJSON(&asstType); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	result, err := cm.control.CMDB().CreateAssociationType(context.TODO(), &asstType)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	r.Result = result
	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) listAssociationType(c *gin.Context) {
	r := httputils.NewResponse()

	result, err := cm.control.CMDB().ListAssociationType(context.TODO())
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	r.Result = result
	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) deleteAssociationType(c *gin.Context) {
	r := httputils.NewResponse()

	p := c.Param("associationId")

	result, err := cm.control.CMDB().DeleteAssociationType(context.TODO(), p)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	r.Result = result
	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) createObjectAssociation(c *gin.Context) {
	r := httputils.NewResponse()

	var objAsst types.ObjAsstDes
	if err := c.ShouldBindJSON(&objAsst); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	result, err := cm.control.CMDB().CreateObjAsst(context.TODO(), &objAsst)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	r.Result = result
	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) listObjectAssociation(c *gin.Context) {
	r := httputils.NewResponse()

	objId := c.Param("objId")

	result, err := cm.control.CMDB().ListObjAsst(context.TODO(), objId)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	r.Result = result
	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) deleteObjectAssociation(c *gin.Context) {
	r := httputils.NewResponse()

	objAsstId := c.Param("objAsstId")

	result, err := cm.control.CMDB().DeleteObjAsst(context.TODO(), objAsstId)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	r.Result = result
	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) createInstAssociation(c *gin.Context) {
	r := httputils.NewResponse()

	var instAsst types.InstAsstDes
	if err := c.ShouldBindJSON(&instAsst); err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	result, err := cm.control.CMDB().CreateInstAsst(context.TODO(), &instAsst)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	r.Result = result
	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) listInstAssociation(c *gin.Context) {
	r := httputils.NewResponse()

	objId := c.Param("objId")

	result, err := cm.control.CMDB().ListInstAsst(context.TODO(), objId)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	r.Result = result
	httputils.SetSuccess(c, r)
}

func (cm *cmdbRouter) deleteInstAssociation(c *gin.Context) {
	r := httputils.NewResponse()

	instAsstId := c.Param("instAsstId")

	result, err := cm.control.CMDB().DeleteInstAsst(context.TODO(), instAsstId)
	if err != nil {
		httputils.SetFailed(c, r, err)
		return
	}

	r.Result = result
	httputils.SetSuccess(c, r)
}

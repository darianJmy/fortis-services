package cmdb

import (
	"context"
	"github.com/darianJmy/fortis-services/pkg/db"
	"github.com/darianJmy/fortis-services/pkg/db/model"
	"github.com/darianJmy/fortis-services/pkg/types"
	"time"
)

type CmdbGetter interface {
	CMDB() Interface
}

type Interface interface {
	CreateObjClassification(ctx context.Context, objCls *types.ObjClassification) (interface{}, error)
	ListObjClassification(ctx context.Context) ([]types.ObjClassification, error)
	UpdateObjClassification(ctx context.Context, clsId string, objCls *types.ObjClassification) (interface{}, error)
	DeleteObjClassification(ctx context.Context, clsId string) (interface{}, error)
	CreateObject(ctx context.Context, obj *types.ObjectDes) (interface{}, error)
	ListObject(ctx context.Context, clsId string) ([]types.ObjectDes, error)
	UpdateObject(ctx context.Context, objId string, obj *types.ObjectDes) (interface{}, error)
	DeleteObject(ctx context.Context, objId string) (interface{}, error)
	CreateObjectAttr(ctx context.Context, obj *types.ObjectAttr) (interface{}, error)
	ListObjectAttr(ctx context.Context, objId string) ([]types.ObjectAttr, error)
	CreateObjectData(ctx context.Context, objId string, obj map[string]string) (interface{}, error)
	ListObjectData(ctx context.Context, objId string) (interface{}, error)
}

type cmdb struct {
	factory *db.ShareDaoFactory
}

func (c *cmdb) CreateObjClassification(ctx context.Context, objCls *types.ObjClassification) (interface{}, error) {
	objClassification := &model.ObjClassification{
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
		ClassificationId:   objCls.ClassificationId,
		ClassificationName: objCls.ClassificationName,
	}

	result, err := c.factory.Cmdb.Create(ctx, objClassification.TableName(), objClassification)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *cmdb) ListObjClassification(ctx context.Context) ([]types.ObjClassification, error) {
	cursor, err := c.factory.Cmdb.List(ctx, model.ObjClassification{}.TableName())
	if err != nil {
		return nil, err
	}

	var objClasses []types.ObjClassification
	for cursor.Next(ctx) {
		var objCls model.ObjClassification
		if err = cursor.Decode(&objCls); err != nil {
			return nil, err
		}
		objClasses = append(objClasses, types.ObjClassification{
			ClassificationId:   objCls.ClassificationId,
			ClassificationName: objCls.ClassificationName,
		})
	}

	return objClasses, nil
}

func (c *cmdb) UpdateObjClassification(ctx context.Context, clsId string, objCls *types.ObjClassification) (interface{}, error) {
	exists, err := c.factory.Cmdb.CheckIdExists(ctx, model.ObjClassification{}.TableName(), "classification_id", clsId)
	if err != nil {
		return nil, err
	}

	if exists {
		return c.factory.Cmdb.Update(ctx, clsId, "classification_id", clsId, &model.ObjClassification{
			ClassificationName: objCls.ClassificationName,
		})
	}

	return nil, nil
}

func (c *cmdb) DeleteObjClassification(ctx context.Context, clsId string) (interface{}, error) {
	exists, err := c.factory.Cmdb.CheckIdExists(ctx, model.ObjClassification{}.TableName(), "classification_id", clsId)
	if err != nil {
		return nil, err
	}

	if exists {
		return c.factory.Cmdb.Delete(ctx, model.ObjClassification{}.TableName(), "classification_id", clsId)
	}

	return nil, nil
}

func (c *cmdb) CreateObject(ctx context.Context, obj *types.ObjectDes) (interface{}, error) {
	objectDes := &model.ObjectDes{
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
		ObjectId:         obj.ObjectId,
		ObjectName:       obj.ObjectName,
		ClassificationId: obj.ClassificationId,
	}

	result, err := c.factory.Cmdb.Create(ctx, objectDes.TableName(), objectDes)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *cmdb) ListObject(ctx context.Context, clsId string) ([]types.ObjectDes, error) {
	cursor, err := c.factory.Cmdb.ListWithFilter(ctx, model.ObjectDes{}.TableName(), "classification_id", clsId)
	if err != nil {
		return nil, err
	}

	var objects []types.ObjectDes
	for cursor.Next(ctx) {
		var obj model.ObjectDes
		if err = cursor.Decode(&obj); err != nil {
			return nil, err
		}
		objects = append(objects, types.ObjectDes{
			ObjectId:   obj.ObjectId,
			ObjectName: obj.ObjectName,
		})
	}

	return objects, nil
}

func (c *cmdb) UpdateObject(ctx context.Context, objId string, obj *types.ObjectDes) (interface{}, error) {
	object := &model.ObjectDes{
		ObjectName:       obj.ObjectName,
		ClassificationId: obj.ClassificationId,
	}

	exists, err := c.factory.Cmdb.CheckIdExists(ctx, model.ObjectDes{}.TableName(), "object_id", objId)
	if err != nil {
		return nil, err
	}

	if exists {
		return c.factory.Cmdb.Update(ctx, model.ObjectDes{}.TableName(), "object_id", objId, object)
	}

	return nil, nil
}

func (c *cmdb) DeleteObject(ctx context.Context, objId string) (interface{}, error) {
	exists, err := c.factory.Cmdb.CheckIdExists(ctx, model.ObjectDes{}.TableName(), "object_id", objId)
	if err != nil {
		return nil, err
	}

	if exists {
		return c.factory.Cmdb.Delete(ctx, model.ObjectDes{}.TableName(), "object_id", objId)
	}

	return c.factory.Cmdb.DeleteObjDes(ctx, objId)
}

func (c *cmdb) CreateObjectAttr(ctx context.Context, obj *types.ObjectAttr) (interface{}, error) {
	objectAttr := &model.ObjectAttDes{
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		ObjectId:     obj.ObjectId,
		PropertyId:   obj.PropertyId,
		PropertyName: obj.PropertyName,
		PropertyType: obj.PropertyType,
	}

	result, err := c.factory.Cmdb.Create(ctx, objectAttr.TableName(), objectAttr)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *cmdb) ListObjectAttr(ctx context.Context, objId string) ([]types.ObjectAttr, error) {
	objAttrs, err := c.factory.Cmdb.ListObjAttr(ctx, objId)
	if err != nil {
		return nil, err
	}

	var typeObjAttrs []types.ObjectAttr
	for _, v := range objAttrs {
		typeObjAttr := types.ObjectAttr{
			PropertyId:   v.PropertyId,
			PropertyName: v.PropertyName,
			PropertyType: v.PropertyType,
			ObjectId:     v.ObjectId,
		}
		typeObjAttrs = append(typeObjAttrs, typeObjAttr)
	}

	return typeObjAttrs, nil
}

func (c *cmdb) CreateObjectData(ctx context.Context, objId string, obj map[string]string) (interface{}, error) {
	attrs, err := c.factory.Cmdb.ListObjAttr(ctx, objId)
	if err != nil {
		return nil, err
	}

	newData := make(map[string]interface{})
	for _, attr := range attrs {
		if value, ok := obj[attr.PropertyId]; ok {
			newData[attr.PropertyId] = value
		}
	}

	result, err := c.factory.Cmdb.CreateObjData(ctx, objId, newData)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *cmdb) ListObjectData(ctx context.Context, objId string) (interface{}, error) {
	objData, err := c.factory.Cmdb.ListObjData(ctx, objId)
	if err != nil {
		return nil, err
	}

	return objData, nil
}

func NewCmdb(f *db.ShareDaoFactory) *cmdb {
	return &cmdb{f}
}

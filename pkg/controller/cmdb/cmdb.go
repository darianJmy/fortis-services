package cmdb

import (
	"context"
	"errors"
	"fmt"
	"github.com/darianJmy/fortis-services/pkg/db"
	"github.com/darianJmy/fortis-services/pkg/db/model"
	"github.com/darianJmy/fortis-services/pkg/types"
	"go.mongodb.org/mongo-driver/bson"
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
	UpdateObjectAttr(ctx context.Context, objId, propertyId string, obj *types.ObjectAttr) (interface{}, error)
	DeleteObjectAttr(ctx context.Context, objId, propertyId string) (interface{}, error)

	CreateInstanceData(ctx context.Context, objId string, obj map[string]string) (interface{}, error)
	ListInstanceData(ctx context.Context, objId string) (interface{}, error)
	UpdateInstanceData(ctx context.Context, objId string, obj map[string]string) (interface{}, error)
	DeleteInstanceData(ctx context.Context, objId, instId string) (interface{}, error)
}

type cmdb struct {
	factory *db.ShareDaoFactory
}

func (c *cmdb) CreateObjClassification(ctx context.Context, objCls *types.ObjClassification) (interface{}, error) {
	if exists, _ := c.factory.Cmdb.CheckIdExists(ctx,
		model.ObjClassification{}.TableName(),
		model.ClassificationId,
		objCls.ClassificationId); exists {
		return nil, errors.New("this objCls id already exists")
	}

	objClassification := &model.ObjClassification{
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
		ClassificationId:   objCls.ClassificationId,
		ClassificationName: objCls.ClassificationName,
	}

	return c.factory.Cmdb.Create(ctx, objClassification.TableName(), objClassification)
}

func (c *cmdb) ListObjClassification(ctx context.Context) ([]types.ObjClassification, error) {
	cursor, err := c.factory.Cmdb.List(ctx, model.ObjClassification{}.TableName())
	if err != nil {
		return nil, err
	}

	var objClasses []types.ObjClassification
	for cursor.Next(ctx) {
		var v model.ObjClassification
		if err = cursor.Decode(&v); err != nil {
			return nil, err
		}
		objClasses = append(objClasses, types.ObjClassification{
			ClassificationId:   v.ClassificationId,
			ClassificationName: v.ClassificationName,
		})
	}

	return objClasses, nil
}

func (c *cmdb) UpdateObjClassification(ctx context.Context, clsId string, objCls *types.ObjClassification) (interface{}, error) {
	if exists, _ := c.factory.Cmdb.CheckIdExists(ctx,
		model.ObjClassification{}.TableName(),
		model.ClassificationId,
		clsId); !exists {
		return nil, errors.New("this objCls id does not exists")
	}

	return c.factory.Cmdb.Update(ctx, model.ObjClassification{}.TableName(), model.ClassificationId, clsId, &model.ObjClassification{
		UpdatedAt:          time.Now(),
		ClassificationName: objCls.ClassificationName,
	})
}

func (c *cmdb) DeleteObjClassification(ctx context.Context, clsId string) (interface{}, error) {
	if exists, _ := c.factory.Cmdb.CheckIdExists(ctx,
		model.ObjClassification{}.TableName(),
		model.ClassificationId,
		clsId); !exists {
		return nil, errors.New("this objCls id does not exists")
	}

	if exists, _ := c.factory.Cmdb.CheckIdExists(ctx,
		model.ObjectDes{}.TableName(),
		model.ClassificationId,
		clsId); exists {
		return nil, errors.New("the classification contains a model and cannot be deleted")
	}

	return c.factory.Cmdb.Delete(ctx, model.ObjClassification{}.TableName(), model.ClassificationId, clsId)
}

func (c *cmdb) CreateObject(ctx context.Context, obj *types.ObjectDes) (interface{}, error) {
	if exists, _ := c.factory.Cmdb.CheckIdExists(ctx,
		model.ObjClassification{}.TableName(),
		model.ClassificationId,
		obj.ClassificationId); !exists {
		return nil, errors.New("this objCls id does not exists")
	}

	if exists, _ := c.factory.Cmdb.CheckIdExists(ctx,
		model.ObjectDes{}.TableName(),
		model.ObjectId,
		obj.ObjectId); exists {
		return nil, errors.New("this object id already exists")
	}

	objectDes := &model.ObjectDes{
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
		ObjectId:         obj.ObjectId,
		ObjectName:       obj.ObjectName,
		ClassificationId: obj.ClassificationId,
	}

	if _, err := c.CreateObjectAttr(ctx, &types.ObjectAttr{
		ObjectId:     obj.ObjectId,
		PropertyId:   model.InstId,
		PropertyName: model.InstName,
		PropertyType: "string",
	}); err != nil {
		return nil, err
	}

	return c.factory.Cmdb.Create(ctx, objectDes.TableName(), objectDes)
}

func (c *cmdb) ListObject(ctx context.Context, clsId string) ([]types.ObjectDes, error) {
	cursor, err := c.factory.Cmdb.ListWithFilter(ctx, model.ObjectDes{}.TableName(), model.ClassificationId, clsId)
	if err != nil {
		return nil, err
	}

	var objects []types.ObjectDes
	for cursor.Next(ctx) {
		var v model.ObjectDes
		if err = cursor.Decode(&v); err != nil {
			return nil, err
		}
		objects = append(objects, types.ObjectDes{
			ObjectId:         v.ObjectId,
			ObjectName:       v.ObjectName,
			ClassificationId: v.ClassificationId,
		})
	}

	return objects, nil
}

func (c *cmdb) UpdateObject(ctx context.Context, objId string, obj *types.ObjectDes) (interface{}, error) {
	if exists, _ := c.factory.Cmdb.CheckIdExists(ctx,
		model.ObjectDes{}.TableName(),
		model.ObjectId,
		objId); !exists {
		return nil, errors.New("this object id does not exists")
	}

	if obj.ClassificationId != "" {
		if exists, _ := c.factory.Cmdb.CheckIdExists(ctx,
			model.ObjClassification{}.TableName(),
			model.ClassificationId,
			obj.ClassificationId); !exists {
			return nil, errors.New("this objCls id does not exists")
		}
	}

	object := &model.ObjectDes{
		UpdatedAt:        time.Now(),
		ObjectName:       obj.ObjectName,
		ClassificationId: obj.ClassificationId,
	}

	return c.factory.Cmdb.Update(ctx, model.ObjectDes{}.TableName(), model.ObjectId, objId, object)
}

func (c *cmdb) DeleteObject(ctx context.Context, objId string) (interface{}, error) {
	if exists, _ := c.factory.Cmdb.CheckIdExists(ctx,
		model.ObjectDes{}.TableName(),
		model.ObjectId,
		objId); !exists {
		return nil, errors.New("this object id does not exists")
	}

	return c.factory.Cmdb.Delete(ctx, model.ObjectDes{}.TableName(), model.ObjectId, objId)
}

func (c *cmdb) CreateObjectAttr(ctx context.Context, obj *types.ObjectAttr) (interface{}, error) {
	if exists, _ := c.factory.Cmdb.CheckIdExists(ctx,
		model.ObjectDes{}.TableName(),
		model.ObjectId,
		obj.ObjectId); !exists {
		return nil, errors.New("this object id does not exists")
	}

	conditions := map[string]string{
		model.ObjectId:   obj.ObjectId,
		model.PropertyId: obj.PropertyId,
	}
	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		model.ObjectAttDes{}.TableName(),
		conditions); exists {
		return nil, errors.New("this property id already exists")
	}

	objectAttr := &model.ObjectAttDes{
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		ObjectId:     obj.ObjectId,
		PropertyId:   obj.PropertyId,
		PropertyName: obj.PropertyName,
		PropertyType: obj.PropertyType,
	}

	return c.factory.Cmdb.Create(ctx, objectAttr.TableName(), objectAttr)
}

func (c *cmdb) ListObjectAttr(ctx context.Context, objId string) ([]types.ObjectAttr, error) {
	cursor, err := c.factory.Cmdb.ListWithFilter(ctx, model.ObjectAttDes{}.TableName(), model.ObjectId, objId)
	if err != nil {
		return nil, err
	}

	var objAttrs []types.ObjectAttr
	for cursor.Next(ctx) {
		var v model.ObjectAttDes
		if err = cursor.Decode(&v); err != nil {
			return nil, err
		}
		objAttrs = append(objAttrs, types.ObjectAttr{
			PropertyId:   v.PropertyId,
			PropertyName: v.PropertyName,
			PropertyType: v.PropertyType,
			ObjectId:     v.ObjectId,
		})
	}

	return objAttrs, nil
}

func (c *cmdb) UpdateObjectAttr(ctx context.Context, objId, propertyId string, obj *types.ObjectAttr) (interface{}, error) {
	conditions := map[string]string{
		model.ObjectId:   objId,
		model.PropertyId: propertyId,
	}
	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		model.ObjectAttDes{}.TableName(),
		conditions); !exists {
		return nil, errors.New("this property id does not exists")
	}

	object := &model.ObjectAttDes{
		UpdatedAt:    time.Now(),
		PropertyName: obj.PropertyName,
	}

	return c.factory.Cmdb.UpdateWithFilter(ctx, model.ObjectAttDes{}.TableName(), conditions, object)
}

func (c *cmdb) DeleteObjectAttr(ctx context.Context, objId, propertyId string) (interface{}, error) {
	conditions := map[string]string{
		model.ObjectId:   objId,
		model.PropertyId: propertyId,
	}
	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		model.ObjectAttDes{}.TableName(),
		conditions); !exists {
		return nil, errors.New("this object id does not exists")
	}

	update := bson.M{
		propertyId: "",
	}

	if _, err := c.factory.Cmdb.DeleteInstDataField(ctx, fmt.Sprintf("Asst_%s", objId), model.ObjectId, objId, update); err != nil {
		return nil, err
	}

	return c.factory.Cmdb.DeleteWithFilter(ctx, model.ObjectDes{}.TableName(), conditions)
}

func (c *cmdb) CreateInstanceData(ctx context.Context, objId string, obj map[string]string) (interface{}, error) {
	if exists, _ := c.factory.Cmdb.CheckIdExists(ctx,
		fmt.Sprintf("Asst_%s", objId),
		model.InstId,
		obj[model.InstId]); exists {
		return nil, errors.New("this inst id already exists")
	}

	objAttrs, err := c.ListObjectAttr(ctx, objId)
	if err != nil {
		return nil, err
	}

	newData := make(map[string]interface{})
	for _, attr := range objAttrs {
		if value, ok := obj[attr.PropertyId]; ok {
			newData[attr.PropertyId] = value
		}
	}

	result, err := c.factory.Cmdb.Create(ctx, fmt.Sprintf("Asst_%s", objId), newData)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *cmdb) ListInstanceData(ctx context.Context, objId string) (interface{}, error) {
	if exists, _ := c.factory.Cmdb.CheckIdExists(ctx,
		model.ObjectDes{}.TableName(),
		model.ObjectId,
		objId); !exists {
		return nil, errors.New("this inst id does not exists")
	}

	cursor, err := c.factory.Cmdb.List(ctx, fmt.Sprintf("Asst_%s", objId))
	if err != nil {
		return nil, err
	}

	var objData []map[string]interface{}
	for cursor.Next(ctx) {
		var v map[string]interface{}
		if err = cursor.Decode(&v); err != nil {
			return nil, err
		}
		objData = append(objData, v)
	}

	return objData, nil
}

func (c *cmdb) UpdateInstanceData(ctx context.Context, objId string, obj map[string]string) (interface{}, error) {
	if exists, _ := c.factory.Cmdb.CheckIdExists(ctx,
		model.ObjectDes{}.TableName(),
		model.ObjectId,
		objId); !exists {
		return nil, nil
	}

	objData, err := c.factory.Cmdb.Update(ctx, fmt.Sprintf("Asst_%s", objId), model.ObjectId, objId, obj)
	if err != nil {
		return nil, err
	}

	return objData, nil
}

func (c *cmdb) DeleteInstanceData(ctx context.Context, objId, instId string) (interface{}, error) {
	if exists, _ := c.factory.Cmdb.CheckIdExists(ctx,
		model.ObjectDes{}.TableName(),
		model.ObjectId,
		objId); !exists {
		return nil, nil
	}

	objData, err := c.factory.Cmdb.Delete(ctx, fmt.Sprintf("Asst_%s", objId), model.ObjectId, objId)
	if err != nil {
		return nil, err
	}

	return objData, nil
}

func NewCmdb(f *db.ShareDaoFactory) *cmdb {
	return &cmdb{f}
}

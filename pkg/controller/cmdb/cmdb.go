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
	condition := bson.M{
		model.ClassificationId: objCls.ClassificationId,
	}

	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		model.ObjClassification{}.TableName(),
		condition); exists {
		return nil, errors.New("this objCls id already exists")
	}

	objClassification := &model.ObjClassification{
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
		ClassificationId:   objCls.ClassificationId,
		ClassificationName: objCls.ClassificationName,
		ClassificationIcon: objCls.ClassificationIcon,
	}

	return c.factory.Cmdb.Create(ctx, objClassification.TableName(), objClassification)
}

func (c *cmdb) ListObjClassification(ctx context.Context) ([]types.ObjClassification, error) {
	cursor, err := c.factory.Cmdb.ListWithFilter(ctx, model.ObjClassification{}.TableName(), nil)
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
			ClassificationIcon: v.ClassificationIcon,
		})
	}

	if err = cursor.Close(ctx); err != nil {
		return nil, err
	}

	return objClasses, nil
}

func (c *cmdb) UpdateObjClassification(ctx context.Context, clsId string, objCls *types.ObjClassification) (interface{}, error) {
	condition := bson.M{
		model.ClassificationId: clsId,
	}

	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		model.ObjClassification{}.TableName(),
		condition); !exists {
		return nil, errors.New("this objCls id does not exists")
	}

	return c.factory.Cmdb.UpdateWithFilter(ctx, model.ObjClassification{}.TableName(), condition, &model.ObjClassification{
		UpdatedAt:          time.Now(),
		ClassificationName: objCls.ClassificationName,
		ClassificationIcon: objCls.ClassificationIcon,
	})
}

func (c *cmdb) DeleteObjClassification(ctx context.Context, clsId string) (interface{}, error) {
	condition := bson.M{
		model.ClassificationId: clsId,
	}

	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		model.ObjClassification{}.TableName(),
		condition); !exists {
		return nil, errors.New("this objCls id does not exists")
	}

	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		model.ObjectDes{}.TableName(),
		condition); exists {
		return nil, errors.New("the classification contains a model and cannot be deleted")
	}

	return c.factory.Cmdb.DeleteWithFilter(ctx, model.ObjClassification{}.TableName(), condition)
}

func (c *cmdb) CreateObject(ctx context.Context, obj *types.ObjectDes) (interface{}, error) {
	clsCondition := bson.M{
		model.ClassificationId: obj.ClassificationId,
	}
	objCondition := bson.M{
		model.ObjectId: obj.ObjectId,
	}

	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		model.ObjClassification{}.TableName(),
		clsCondition); !exists {
		return nil, errors.New("this objCls id does not exists")
	}
	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		model.ObjectDes{}.TableName(),
		objCondition); exists {
		return nil, errors.New("this object id already exists")
	}

	objectDes := &model.ObjectDes{
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
		ObjectId:         obj.ObjectId,
		ObjectName:       obj.ObjectName,
		ClassificationId: obj.ClassificationId,
	}
	objectAttDes := &model.ObjectAttDes{
		ObjectId:     obj.ObjectId,
		PropertyId:   model.InstId,
		PropertyName: model.InstName,
		PropertyType: "string",
	}

	return c.factory.Cmdb.CreateObjectWithTransaction(ctx, objectDes, objectAttDes)
}

func (c *cmdb) ListObject(ctx context.Context, clsId string) ([]types.ObjectDes, error) {
	condition := bson.M{
		model.ClassificationId: clsId,
	}

	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		model.ObjClassification{}.TableName(),
		condition); exists {
		return nil, errors.New("this objCls id already exists")
	}

	cursor, err := c.factory.Cmdb.ListWithFilter(ctx, model.ObjectDes{}.TableName(), condition)
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

	if err = cursor.Close(ctx); err != nil {
		return nil, err
	}

	return objects, nil
}

func (c *cmdb) UpdateObject(ctx context.Context, objId string, obj *types.ObjectDes) (interface{}, error) {
	objCondition := bson.M{
		model.ObjectId: objId,
	}
	clsCondition := bson.M{
		model.ClassificationId: obj.ClassificationId,
	}

	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		model.ObjectDes{}.TableName(),
		objCondition); !exists {
		return nil, errors.New("this object id does not exists")
	}
	if obj.ClassificationId != "" {
		if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
			model.ObjClassification{}.TableName(),
			clsCondition); !exists {
			return nil, errors.New("this objCls id does not exists")
		}
	}

	object := &model.ObjectDes{
		UpdatedAt:        time.Now(),
		ObjectName:       obj.ObjectName,
		ClassificationId: obj.ClassificationId,
	}

	return c.factory.Cmdb.UpdateWithFilter(ctx, model.ObjectDes{}.TableName(), objCondition, object)
}

func (c *cmdb) DeleteObject(ctx context.Context, objId string) (interface{}, error) {
	condition := bson.M{
		model.ObjectId: objId,
	}

	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		model.ObjectDes{}.TableName(),
		condition); !exists {
		return nil, errors.New("this object id does not exists")
	}
	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		fmt.Sprintf("Asst_%s", objId),
		nil); exists {
		return nil, errors.New("the object contains a data and cannot be deleted")
	}

	return c.factory.Cmdb.DeleteWithFilter(ctx, model.ObjectDes{}.TableName(), condition)
}

func (c *cmdb) CreateObjectAttr(ctx context.Context, obj *types.ObjectAttr) (interface{}, error) {
	objCondition := bson.M{
		model.ObjectId: obj.ObjectId,
	}
	objAttCondition := bson.M{
		model.ObjectId:   obj.ObjectId,
		model.PropertyId: obj.PropertyId,
	}

	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		model.ObjectDes{}.TableName(),
		objCondition); !exists {
		return nil, errors.New("this object id does not exists")
	}
	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		model.ObjectAttDes{}.TableName(),
		objAttCondition); exists {
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
	condition := bson.M{
		model.ObjectId: objId,
	}

	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		model.ObjectDes{}.TableName(),
		condition); !exists {
		return nil, errors.New("this object id does not exists")
	}

	cursor, err := c.factory.Cmdb.ListWithFilter(ctx, model.ObjectAttDes{}.TableName(), condition)
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

	if err = cursor.Close(ctx); err != nil {
		return nil, err
	}

	return objAttrs, nil
}

func (c *cmdb) UpdateObjectAttr(ctx context.Context, objId, propertyId string, obj *types.ObjectAttr) (interface{}, error) {
	objCondition := bson.M{
		model.ObjectId: obj,
	}
	objAttDesCondition := bson.M{
		model.ObjectId:   objId,
		model.PropertyId: propertyId,
	}

	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		model.ObjectDes{}.TableName(),
		objCondition); !exists {
		return nil, errors.New("this object id does not exists")
	}
	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		model.ObjectAttDes{}.TableName(),
		objAttDesCondition); !exists {
		return nil, errors.New("this property id or object id does not exists")
	}

	object := &model.ObjectAttDes{
		UpdatedAt:    time.Now(),
		PropertyName: obj.PropertyName,
	}

	return c.factory.Cmdb.UpdateWithFilter(ctx, model.ObjectAttDes{}.TableName(), objAttDesCondition, object)
}

func (c *cmdb) DeleteObjectAttr(ctx context.Context, objId, propertyId string) (interface{}, error) {
	objCondition := bson.M{
		model.ObjectId: objId,
	}
	objAttDesCondition := bson.M{
		model.ObjectId:   objId,
		model.PropertyId: propertyId,
	}

	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		model.ObjectDes{}.TableName(),
		objCondition); !exists {
		return nil, errors.New("this object id does not exists")
	}
	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		model.ObjectAttDes{}.TableName(),
		objAttDesCondition); !exists {
		return nil, errors.New("this property id or object id does not exists")
	}

	return c.factory.Cmdb.DeleteObjectAttrWithTransaction(ctx, objId, propertyId, objAttDesCondition)
}

func (c *cmdb) CreateInstanceData(ctx context.Context, objId string, obj map[string]string) (interface{}, error) {
	objCondition := bson.M{model.ObjectId: objId}
	instCondition := bson.M{model.InstId: obj[model.InstId]}

	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		model.ObjectDes{}.TableName(),
		objCondition); !exists {
		return nil, errors.New("this object id does not exists")
	}
	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		fmt.Sprintf("Asst_%s", objId),
		instCondition); exists {
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

	return c.factory.Cmdb.Create(ctx, fmt.Sprintf("Asst_%s", objId), newData)
}

func (c *cmdb) ListInstanceData(ctx context.Context, objId string) (interface{}, error) {
	condition := bson.M{
		model.InstId: objId,
	}

	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		model.ObjectDes{}.TableName(),
		condition); !exists {
		return nil, errors.New("this inst id does not exists")
	}

	cursor, err := c.factory.Cmdb.ListWithFilter(ctx, fmt.Sprintf("Asst_%s", objId), nil)
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
	condition := bson.M{
		model.ObjectId: objId,
	}

	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		model.ObjectDes{}.TableName(),
		condition); !exists {
		return nil, nil
	}

	var updateData map[string]string
	for k, v := range obj {
		if k != model.InstId {
			updateData[k] = v
		}
	}

	return c.factory.Cmdb.UpdateWithFilter(ctx, fmt.Sprintf("Asst_%s", objId), condition, updateData)
}

func (c *cmdb) DeleteInstanceData(ctx context.Context, objId, instId string) (interface{}, error) {
	objCondition := bson.M{
		model.ObjectId: objId,
	}
	instCondition := bson.M{
		model.InstId: instId,
	}

	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		model.ObjectDes{}.TableName(),
		objCondition); !exists {
		return nil, nil
	}

	objData, err := c.factory.Cmdb.DeleteWithFilter(ctx, fmt.Sprintf("Asst_%s", objId), instCondition)
	if err != nil {
		return nil, err
	}

	return objData, nil
}

func NewCmdb(f *db.ShareDaoFactory) *cmdb {
	return &cmdb{f}
}

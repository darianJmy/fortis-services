package cmdb

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/darianJmy/fortis-services/pkg/db"
	"github.com/darianJmy/fortis-services/pkg/db/model"
	"github.com/darianJmy/fortis-services/pkg/types"
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
	UpdateInstanceData(ctx context.Context, objId, instId string, obj map[string]string) (interface{}, error)
	DeleteInstanceData(ctx context.Context, objId, instId string) (interface{}, error)

	CreateAssociationType(ctx context.Context, asstType *types.AssociationType) (interface{}, error)
	ListAssociationType(ctx context.Context) ([]types.AssociationType, error)
	DeleteAssociationType(ctx context.Context, asstId string) (interface{}, error)

	CreateObjAsst(ctx context.Context, objAsst *types.ObjAsstDes) (interface{}, error)
	ListObjAsst(ctx context.Context, objId string) (interface{}, error)
	DeleteObjAsst(ctx context.Context, objAsstId string) (interface{}, error)

	CreateInstAsst(ctx context.Context, instAsst *types.InstAsstDes) (interface{}, error)
	ListInstAsst(ctx context.Context, objId string) (interface{}, error)
	DeleteInstAsst(ctx context.Context, instAsstId string) (interface{}, error)
}

type cmdb struct {
	factory *db.ShareDaoFactory
}

func (c *cmdb) CreateObjClassification(ctx context.Context, objCls *types.ObjClassification) (interface{}, error) {
	condition := bson.M{model.ClassificationId: objCls.ClassificationId}

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
	condition := bson.M{model.ClassificationId: clsId}

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
	condition := bson.M{model.ClassificationId: clsId}

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
	clsCondition := bson.M{model.ClassificationId: obj.ClassificationId}
	objCondition := bson.M{model.ObjectId: obj.ObjectId}

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
		ObjectId:         obj.ObjectId,
		ObjectName:       obj.ObjectName,
		Description:      obj.Description,
		ClassificationId: obj.ClassificationId,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
	objectAttDes := &model.ObjectAttDes{
		ObjectId:     obj.ObjectId,
		PropertyId:   model.InstId,
		PropertyName: model.InstName,
		PropertyType: "string",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	return c.factory.Cmdb.CreateObjectWithTransaction(ctx, objectDes, objectAttDes)
}

func (c *cmdb) ListObject(ctx context.Context, clsId string) ([]types.ObjectDes, error) {
	condition := bson.M{model.ClassificationId: clsId}

	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		model.ObjClassification{}.TableName(),
		condition); !exists {
		return nil, errors.New("this objCls id does not exists")
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
			Description:      v.Description,
		})
	}

	if err = cursor.Close(ctx); err != nil {
		return nil, err
	}

	return objects, nil
}

func (c *cmdb) UpdateObject(ctx context.Context, objId string, obj *types.ObjectDes) (interface{}, error) {
	objCondition := bson.M{model.ObjectId: objId}
	clsCondition := bson.M{model.ClassificationId: obj.ClassificationId}

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
	condition := bson.M{model.ObjectId: objId}

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
	objCondition := bson.M{model.ObjectId: obj.ObjectId}
	objAttCondition := bson.M{model.ObjectId: obj.ObjectId, model.PropertyId: obj.PropertyId}

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
		ObjectId:     obj.ObjectId,
		PropertyId:   obj.PropertyId,
		PropertyName: obj.PropertyName,
		PropertyType: obj.PropertyType,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	return c.factory.Cmdb.Create(ctx, objectAttr.TableName(), objectAttr)
}

func (c *cmdb) ListObjectAttr(ctx context.Context, objId string) ([]types.ObjectAttr, error) {
	condition := bson.M{model.ObjectId: objId}

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
			ObjectId:     v.ObjectId,
			PropertyId:   v.PropertyId,
			PropertyName: v.PropertyName,
			PropertyType: v.PropertyType,
		})
	}

	if err = cursor.Close(ctx); err != nil {
		return nil, err
	}

	return objAttrs, nil
}

func (c *cmdb) UpdateObjectAttr(ctx context.Context, objId, propertyId string, obj *types.ObjectAttr) (interface{}, error) {
	objCondition := bson.M{model.ObjectId: objId}
	objAttCondition := bson.M{model.ObjectId: objId, model.PropertyId: propertyId}

	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		model.ObjectDes{}.TableName(),
		objCondition); !exists {
		return nil, errors.New("this object id does not exists")
	}
	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		model.ObjectAttDes{}.TableName(),
		objAttCondition); !exists {
		return nil, errors.New("this property id or object id does not exists")
	}

	object := &model.ObjectAttDes{
		UpdatedAt:    time.Now(),
		PropertyName: obj.PropertyName,
	}

	return c.factory.Cmdb.UpdateWithFilter(ctx, model.ObjectAttDes{}.TableName(), objAttCondition, object)
}

func (c *cmdb) DeleteObjectAttr(ctx context.Context, objId, propertyId string) (interface{}, error) {
	objCondition := bson.M{model.ObjectId: objId}
	objAttCondition := bson.M{model.ObjectId: objId, model.PropertyId: propertyId}

	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		model.ObjectDes{}.TableName(),
		objCondition); !exists {
		return nil, errors.New("this object id does not exists")
	}
	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		model.ObjectAttDes{}.TableName(),
		objAttCondition); !exists {
		return nil, errors.New("this property id or object id does not exists")
	}

	return c.factory.Cmdb.DeleteObjectAttrWithTransaction(ctx, objId, propertyId, objAttCondition)
}

func (c *cmdb) CreateInstanceData(ctx context.Context, objId string, obj map[string]string) (interface{}, error) {
	objCondition := bson.M{model.ObjectId: objId}
	instCondition := bson.M{model.InstId: obj[model.InstId]}

	if obj[model.InstId] == "" {
		return nil, errors.New("this inst id is empty")
	}
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
	condition := bson.M{model.ObjectId: objId}

	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		model.ObjectDes{}.TableName(),
		condition); !exists {
		return nil, errors.New("this object id does not exists")
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

	if err = cursor.Close(ctx); err != nil {
		return nil, err
	}

	return objData, nil
}

func (c *cmdb) UpdateInstanceData(ctx context.Context, objId, instId string, obj map[string]string) (interface{}, error) {
	objCondition := bson.M{model.ObjectId: objId}
	instCondition := bson.M{model.InstId: instId}

	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		model.ObjectDes{}.TableName(),
		objCondition); !exists {
		return nil, errors.New("this object id does not exists")
	}
	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		fmt.Sprintf("Asst_%s", objId),
		instCondition); !exists {
		return nil, errors.New("this inst id does not exists")
	}

	var updateData = make(map[string]string)
	for k, v := range obj {
		if k != model.InstId {
			updateData[k] = v
		}
	}

	return c.factory.Cmdb.UpdateWithFilter(ctx, fmt.Sprintf("Asst_%s", objId), instCondition, updateData)
}

func (c *cmdb) DeleteInstanceData(ctx context.Context, objId, instId string) (interface{}, error) {
	objCondition := bson.M{model.ObjectId: objId}
	instCondition := bson.M{model.InstId: instId}

	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		model.ObjectDes{}.TableName(),
		objCondition); !exists {
		return nil, errors.New("this object id does not exists")
	}
	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		fmt.Sprintf("Asst_%s", objId),
		instCondition); !exists {
		return nil, errors.New("this inst id does not exists")
	}

	objData, err := c.factory.Cmdb.DeleteWithFilter(ctx, fmt.Sprintf("Asst_%s", objId), instCondition)
	if err != nil {
		return nil, err
	}

	return objData, nil
}

func (c *cmdb) CreateAssociationType(ctx context.Context, asstType *types.AssociationType) (interface{}, error) {
	condition := bson.M{model.AsstId: asstType.AsstId}

	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		model.AsstDes{}.TableName(),
		condition); exists {
		return nil, errors.New("this asst id already exists")
	}

	asstDes := &model.AsstDes{
		AsstId:    asstType.AsstId,
		AsstName:  asstType.AsstName,
		SrcDes:    asstType.SrcDes,
		DestDes:   asstType.DestDes,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return c.factory.Cmdb.Create(ctx, asstDes.TableName(), asstDes)
}

func (c *cmdb) ListAssociationType(ctx context.Context) ([]types.AssociationType, error) {
	cursor, err := c.factory.Cmdb.ListWithFilter(ctx, model.AsstDes{}.TableName(), nil)
	if err != nil {
		return nil, err
	}

	var asstTypes []types.AssociationType
	for cursor.Next(ctx) {
		var v model.AsstDes
		if err = cursor.Decode(&v); err != nil {
			return nil, err
		}
		asstTypes = append(asstTypes, types.AssociationType{
			AsstId:  v.AsstId,
			SrcDes:  v.SrcDes,
			DestDes: v.DestDes,
		})
	}

	if err = cursor.Close(ctx); err != nil {
		return nil, err
	}

	return asstTypes, nil
}

func (c *cmdb) DeleteAssociationType(ctx context.Context, asstId string) (interface{}, error) {
	condition := bson.M{model.AsstId: asstId}

	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		model.AsstDes{}.TableName(),
		condition); !exists {
		return nil, errors.New("this objCls id does not exists")
	}

	return c.factory.Cmdb.DeleteWithFilter(ctx, model.AsstDes{}.TableName(), condition)
}

func (c *cmdb) CreateObjAsst(ctx context.Context, objAsst *types.ObjAsstDes) (interface{}, error) {
	condition := bson.M{model.ObjAsstId: fmt.Sprintf("%s_%s_%s", objAsst.SrcObjId, objAsst.ObjAsstId, objAsst.DestObjId)}

	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		model.ObjAsstDes{}.TableName(),
		condition); exists {
		return nil, errors.New("this asst id already exists")
	}

	asstDes := &model.ObjAsstDes{
		ObjAsstId: fmt.Sprintf("%s_%s_%s", objAsst.SrcObjId, objAsst.AsstId, objAsst.DestObjId),
		SrcObjId:  objAsst.SrcObjId,
		DestObjId: objAsst.DestObjId,
		AsstId:    objAsst.AsstId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return c.factory.Cmdb.Create(ctx, asstDes.TableName(), asstDes)
}

func (c *cmdb) ListObjAsst(ctx context.Context, objId string) (interface{}, error) {
	conditions := []bson.M{{model.SrcObjId: objId}, {model.DestObjId: objId}}

	var objAsstDes []types.ObjAsstDes
	for _, condition := range conditions {
		cursor, err := c.factory.Cmdb.ListWithFilter(ctx, model.ObjAsstDes{}.TableName(), condition)
		if err != nil {
			return nil, err
		}

		for cursor.Next(ctx) {
			var v model.ObjAsstDes
			if err = cursor.Decode(&v); err != nil {
				return nil, err
			}
			objAsstDes = append(objAsstDes, types.ObjAsstDes{
				ObjAsstId: v.ObjAsstId,
				SrcObjId:  v.SrcObjId,
				DestObjId: v.DestObjId,
				AsstId:    v.AsstId,
			})
		}

		if err = cursor.Close(ctx); err != nil {
			return nil, err
		}
	}

	return objAsstDes, nil
}

func (c *cmdb) DeleteObjAsst(ctx context.Context, objAsstId string) (interface{}, error) {
	condition := bson.M{model.ObjAsstId: objAsstId}

	if exists, _ := c.factory.Cmdb.CheckIdExistsWithFilter(ctx,
		model.ObjAsstDes{}.TableName(),
		condition); !exists {
		return nil, errors.New("this obj asst id does not exists")
	}

	return c.factory.Cmdb.DeleteWithFilter(ctx, model.ObjAsstDes{}.TableName(), condition)
}

func (c *cmdb) CreateInstAsst(ctx context.Context, instAsst *types.InstAsstDes) (interface{}, error) {
	return nil, nil
}

func (c *cmdb) ListInstAsst(ctx context.Context, objId string) (interface{}, error) { return nil, nil }

func (c *cmdb) DeleteInstAsst(ctx context.Context, instAsstId string) (interface{}, error) {
	return nil, nil
}

func NewCmdb(f *db.ShareDaoFactory) *cmdb {
	return &cmdb{f}
}

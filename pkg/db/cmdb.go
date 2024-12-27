package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/darianJmy/fortis-services/pkg/db/model"
)

type cmdb struct {
	db *mongo.Database
}

func (c *cmdb) CreateObjClassification(ctx context.Context, ojbCls *model.ObjClassification) (interface{}, error) {
	result, err := c.db.Collection(ojbCls.TableName()).InsertOne(ctx, ojbCls)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

func (c *cmdb) ListObjClassification(ctx context.Context) ([]model.ObjClassification, error) {
	var (
		objCls     model.ObjClassification
		objClasses []model.ObjClassification
	)

	cursor, err := c.db.Collection(objCls.TableName()).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var r model.ObjClassification
		if err = cursor.Decode(&r); err != nil {
			return nil, err
		}
		objClasses = append(objClasses, r)
	}

	return objClasses, nil
}

func (c *cmdb) UpdateObjClassification(ctx context.Context, clsId string, objCls *model.ObjClassification) (interface{}, error) {
	result, err := c.db.Collection(objCls.TableName()).UpdateOne(ctx, bson.M{"classification_id": clsId}, objCls)
	if err != nil {
		return nil, err
	}

	return result.UpsertedID, nil
}

func (c *cmdb) DeleteObjClassification(ctx context.Context, clsId string) (interface{}, error) {
	var objCls model.ObjClassification

	result, err := c.db.Collection(objCls.TableName()).DeleteOne(ctx, bson.M{"classification_id": clsId})
	if err != nil {
		return nil, err
	}

	return result.DeletedCount, nil
}

func (c *cmdb) CreateObjDes(ctx context.Context, obj *model.ObjectDes) (interface{}, error) {
	result, err := c.db.Collection(obj.TableName()).InsertOne(ctx, obj)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

func (c *cmdb) ListObjDes(ctx context.Context, clsId string) ([]model.ObjectDes, error) {
	var (
		obj  model.ObjectDes
		objs []model.ObjectDes
	)

	cursor, err := c.db.Collection(obj.TableName()).Find(ctx, bson.M{"classification_id": clsId})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var r model.ObjectDes
		if err = cursor.Decode(&r); err != nil {
			return nil, err
		}
		objs = append(objs, r)
	}

	return objs, nil
}

func (c *cmdb) UpdateObjDes(ctx context.Context, objId string, obj *model.ObjectDes) (interface{}, error) {
	result, err := c.db.Collection(obj.TableName()).UpdateOne(ctx, bson.M{"object_id": objId}, obj)
	if err != nil {
		return nil, err
	}

	return result.UpsertedID, nil
}

func (c *cmdb) DeleteObjDes(ctx context.Context, objId string) (interface{}, error) {
	var obj model.ObjectDes

	result, err := c.db.Collection(obj.TableName()).DeleteOne(ctx, bson.M{"object_id": objId})
	if err != nil {
		return nil, err
	}

	return result.DeletedCount, nil
}

func (c *cmdb) CreateObjAttr(ctx context.Context, obj *model.ObjectAttDes) (interface{}, error) {
	result, err := c.db.Collection(obj.TableName()).InsertOne(ctx, obj)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

func (c *cmdb) ListObjAttr(ctx context.Context, objId string) ([]model.ObjectAttDes, error) {
	var (
		objAttr  model.ObjectAttDes
		objAttrs []model.ObjectAttDes
	)

	cursor, err := c.db.Collection(objAttr.TableName()).Find(ctx, bson.M{"object_id": objId})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var r model.ObjectAttDes
		if err = cursor.Decode(&r); err != nil {
			return nil, err
		}
		objAttrs = append(objAttrs, r)
	}

	return objAttrs, nil
}

func (c *cmdb) UpdateObjAttr(ctx context.Context, propertyId string, objAttr *model.ObjectAttDes) (interface{}, error) {
	result, err := c.db.Collection(objAttr.TableName()).UpdateOne(ctx, bson.M{"property_id": propertyId}, objAttr)
	if err != nil {
		return nil, err
	}

	return result.UpsertedID, nil
}

func (c *cmdb) DeleteObjAttr(ctx context.Context, propertyId string) (interface{}, error) {
	var objAttr model.ObjClassification

	result, err := c.db.Collection(objAttr.TableName()).DeleteOne(ctx, bson.M{"property_id": propertyId})
	if err != nil {
		return nil, err
	}

	return result.DeletedCount, nil
}

func (c *cmdb) CreateObjData(ctx context.Context, objId string, object interface{}) (interface{}, error) {
	result, err := c.db.Collection(fmt.Sprintf("Asst_%s", objId)).InsertOne(ctx, object)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (c *cmdb) ListObjData(ctx context.Context, objId string) (interface{}, error) {
	cursor, err := c.db.Collection(fmt.Sprintf("Asst_%s", objId)).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var objData []interface{}
	for cursor.Next(ctx) {
		var r interface{}
		if err = cursor.Decode(&r); err != nil {
			return nil, err
		}
		objData = append(objData, r)
	}

	return objData, err
}

func newCmdb(db *mongo.Database) *cmdb {
	return &cmdb{db}
}

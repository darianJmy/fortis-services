package db

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/darianJmy/fortis-services/pkg/db/model"
)

type cmdb struct {
	db *mongo.Database
}

func (c *cmdb) Create(ctx context.Context, tableName string, obj interface{}) (interface{}, error) {
	result, err := c.db.Collection(tableName).InsertOne(ctx, obj)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

func (c *cmdb) List(ctx context.Context, tableName string) (*mongo.Cursor, error) {
	cursor, err := c.db.Collection(tableName).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	return cursor, nil
}

func (c *cmdb) ListWithFilter(ctx context.Context, tableName string, key, value string) (*mongo.Cursor, error) {
	cursor, err := c.db.Collection(tableName).Find(ctx, bson.M{key: value})
	if err != nil {
		return nil, err
	}

	return cursor, nil
}

func (c *cmdb) Update(ctx context.Context, tableName, key, value string, obj interface{}) (interface{}, error) {
	result, err := c.db.Collection(tableName).UpdateOne(ctx, bson.M{key: value}, bson.M{"$set": obj})
	if err != nil {
		return nil, err
	}

	return result.UpsertedID, nil
}

func (c *cmdb) Delete(ctx context.Context, tableName, key, value string) (interface{}, error) {
	result, err := c.db.Collection(tableName).DeleteOne(ctx, bson.M{key: value})
	if err != nil {
		return nil, err
	}

	return result.DeletedCount, nil
}

func (c *cmdb) CheckIdExists(ctx context.Context, tableName string, key, value string) (bool, error) {
	if err := c.db.Collection(tableName).FindOne(ctx, bson.M{key: value}).Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}

	return true, nil
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

func (c *cmdb) CheckObjDesExists(ctx context.Context, objId string) (bool, error) {
	var obj model.ObjectDes

	_, err := c.db.Collection(obj.TableName()).Find(ctx, bson.M{"object_id": objId})
	if err != nil {
		return false, err
	}

	return true, nil
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

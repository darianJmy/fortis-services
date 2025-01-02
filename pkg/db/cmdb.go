package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/darianJmy/fortis-services/pkg/db/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

func (c *cmdb) ListWithFilter(ctx context.Context, tableName string, condition bson.M) (*mongo.Cursor, error) {
	return c.db.Collection(tableName).Find(ctx, condition)
}

func (c *cmdb) UpdateWithFilter(ctx context.Context, tableName string, condition bson.M, obj interface{}) (interface{}, error) {
	result, err := c.db.Collection(tableName).UpdateOne(ctx, condition, bson.M{"$set": obj})
	if err != nil {
		return nil, err
	}

	return result.UpsertedID, nil
}

func (c *cmdb) DeleteWithFilter(ctx context.Context, tableName string, condition bson.M) (interface{}, error) {
	result, err := c.db.Collection(tableName).DeleteOne(ctx, condition)
	if err != nil {
		return nil, err
	}

	return result.DeletedCount, nil
}

func (c *cmdb) CheckIdExistsWithFilter(ctx context.Context, tableName string, condition bson.M) (bool, error) {
	if err := c.db.Collection(tableName).FindOne(ctx, condition).Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (c *cmdb) CreateObjectWithTransaction(ctx context.Context, obj, objAttDes interface{}) (interface{}, error) {
	session, err := c.db.Client().StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(ctx)

	callback := func(ctx mongo.SessionContext) (interface{}, error) {
		if _, err = c.db.Collection(model.ObjectDes{}.TableName()).
			InsertOne(ctx, obj); err != nil {
			return nil, err
		}

		if _, err = c.db.Collection(model.ObjectAttDes{}.TableName()).
			InsertOne(ctx, objAttDes); err != nil {
			return nil, err
		}
		return nil, nil
	}

	return session.WithTransaction(ctx, callback)
}

func (c *cmdb) DeleteObjectAttrWithTransaction(ctx context.Context, objId, propertyId string, condition bson.M) (interface{}, error) {
	session, err := c.db.Client().StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(ctx)

	callback := func(ctx mongo.SessionContext) (interface{}, error) {
		if _, err = c.db.Collection(model.ObjectAttDes{}.TableName()).DeleteOne(ctx, condition); err != nil {
			return nil, err
		}

		if _, err = c.db.Collection(fmt.Sprintf("Asst_%s", objId)).
			DeleteMany(ctx, bson.M{"$unset": bson.M{propertyId: ""}}); err != nil {
			return nil, err
		}
		return nil, nil
	}

	return session.WithTransaction(ctx, callback)
}

func newCmdb(db *mongo.Database) *cmdb {
	return &cmdb{db}
}

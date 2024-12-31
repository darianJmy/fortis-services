package db

import (
	"context"
	"errors"
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

func (c *cmdb) UpdateWithFilter(ctx context.Context, tableName string, conditions map[string]string, obj interface{}) (interface{}, error) {
	filter := bson.M{}
	for k, v := range conditions {
		filter[k] = v
	}

	result, err := c.db.Collection(tableName).UpdateOne(ctx, filter, bson.M{"$set": obj})
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

func (c *cmdb) DeleteWithFilter(ctx context.Context, tableName string, conditions map[string]string) (interface{}, error) {
	filter := bson.M{}
	for k, v := range conditions {
		filter[k] = v
	}

	result, err := c.db.Collection(tableName).DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return result.DeletedCount, nil
}

func (c *cmdb) DeleteInstDataField(ctx context.Context, tableName, key, value string, obj interface{}) (interface{}, error) {
	result, err := c.db.Collection(tableName).UpdateOne(ctx, bson.M{key: value}, bson.M{"$unset": obj})
	if err != nil {
		return nil, err
	}

	return result.UpsertedID, nil
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

func (c *cmdb) CheckIdExistsWithFilter(ctx context.Context, tableName string, conditions map[string]string) (bool, error) {
	filter := bson.M{}
	for k, v := range conditions {
		filter[k] = v
	}

	if err := c.db.Collection(tableName).FindOne(ctx, filter).Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func newCmdb(db *mongo.Database) *cmdb {
	return &cmdb{db}
}

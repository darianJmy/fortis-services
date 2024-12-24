package db

import (
	"context"
	"github.com/darianJmy/fortis-services/pkg/db/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type cmdb struct {
	db *mongo.Database
}

func (c *cmdb) CreateCollection(ctx context.Context, m *model.ObjectDes) error {
	if _, err := c.db.Collection(m.TableName()).InsertOne(ctx, &model.ObjectDes{
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		ObjectName: m.ObjectName,
		ObjectId:   m.ObjectId,
	}); err != nil {
		return err
	}

	return nil
}

func (c *cmdb) UpdateCollection(ctx context.Context, mId string, m *model.ObjectDes) error {
	if _, err := c.db.Collection(m.TableName()).UpdateByID(ctx, mId, m); err != nil {
		return err
	}

	return nil
}

func (c *cmdb) DeleteCollection(ctx context.Context, mId string) error {
	filter := bson.M{"_id": mId}

	des := model.ObjectDes{}

	if _, err := c.db.Collection(des.TableName()).DeleteOne(ctx, filter); err != nil {
		return err
	}

	return nil
}

func (c *cmdb) GetCollection(ctx context.Context, mId string) (*model.ObjectDes, error) {
	var des model.ObjectDes

	filter := bson.M{"object_id": mId}

	cursor := c.db.Collection(des.TableName()).FindOne(ctx, filter)

	if err := cursor.Decode(&des); err != nil {
		return nil, err
	}

	return &des, nil

}

func (c *cmdb) ListCollection(ctx context.Context, object interface{}) error {
	return nil
}

func (c *cmdb) CreateField(ctx context.Context, object interface{}) error {
	return nil
}

func (c *cmdb) CreateDefaultField(ctx context.Context, mId string) error {
	var (
		attDes model.ObjectAttDes
		err    error
	)

	for _, v := range defaultFields() {
		_, err = c.db.Collection(attDes.TableName()).InsertOne(ctx, &model.ObjectAttDes{
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			ObjectId:     mId,
			PropertyId:   v.PropertyId,
			PropertyType: v.PropertyType,
			Editable:     v.Editable,
			IsRequired:   v.IsRequired,
		})
	}

	if err != nil {
		return err
	}

	return nil
}

func (c *cmdb) UpdateField(ctx context.Context, object interface{}) error {
	return nil
}

func (c *cmdb) DeleteField(ctx context.Context, object interface{}) error {
	return nil
}

func (c *cmdb) GetField(ctx context.Context, object interface{}) error {
	return nil
}

func (c *cmdb) ListField(ctx context.Context, object interface{}) error {
	return nil
}

func (c *cmdb) CreateData(ctx context.Context, object interface{}) error {
	return nil
}

func (c *cmdb) UpdateData(ctx context.Context, object interface{}) error {
	return nil
}

func (c *cmdb) DeleteData(ctx context.Context, object interface{}) error {
	return nil
}

func (c *cmdb) GetData(ctx context.Context, object interface{}) error {
	return nil
}

func (c *cmdb) ListData(ctx context.Context, object interface{}) error {
	return nil
}

func (c *cmdb) CreateObjClassification(ctx context.Context, obj model.ObjClassification) error {
	_, err := c.db.Collection(obj.TableName()).InsertOne(ctx, obj)
	if err != nil {
		return err
	}

	return nil
}

func (c *cmdb) CreateObjDes(ctx context.Context, obj model.ObjectDes) error {
	_, err := c.db.Collection(obj.TableName()).InsertOne(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

func (c *cmdb) CreateObjAttr(ctx context.Context, obj model.ObjectAttDes) error {
	_, err := c.db.Collection(obj.TableName()).InsertOne(ctx, obj)
	if err != nil {
		return err
	}
	return nil
}

type fields struct {
	PropertyId   string `bson:"property_id"`
	PropertyType string `bson:"property_type"`
	Editable     bool   `bson:"editable"`
	IsRequired   bool   `bson:"is_required"`
}

func defaultFields() []fields {
	return []fields{
		{PropertyId: "inst_name", PropertyType: "singlechar", Editable: true, IsRequired: true},
		{PropertyId: "create_at", PropertyType: "time", Editable: false, IsRequired: false},
		{PropertyId: "create_by", PropertyType: "objuser", Editable: false, IsRequired: false},
		{PropertyId: "update_at", PropertyType: "time", Editable: false, IsRequired: false},
		{PropertyId: "update_by", PropertyType: "objuser", Editable: false, IsRequired: false},
	}
}

func newCmdb(db *mongo.Database) *cmdb {
	return &cmdb{db}
}

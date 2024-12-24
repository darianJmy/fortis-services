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
	CreateModel(ctx context.Context, m *types.ModelDes) (*types.ModelDes, error)
	GetModel(ctx context.Context, mId string) (*types.ModelDes, error)
	GetResource(ctx context.Context) ([]types.Resource, error)
	CreateObjClassification(ctx context.Context, obj *types.ObjClassification) error
	CreateObject(ctx context.Context, obj *types.ObjectDes) error
	CreateObjectAttr(ctx context.Context, obj *types.ObjectAttr) error
}

type cmdb struct {
	factory *db.ShareDaoFactory
}

func (c *cmdb) CreateModel(ctx context.Context, m *types.ModelDes) (*types.ModelDes, error) {

	if err := c.factory.Cmdb.CreateCollection(ctx, &model.ObjectDes{ObjectId: m.ObjectId, ObjectName: m.ObjectName}); err != nil {
		return nil, err
	}

	if err := c.factory.Cmdb.CreateDefaultField(ctx, m.ObjectId); err != nil {
		return nil, err
	}

	return m, nil
}

func (c *cmdb) GetModel(ctx context.Context, mId string) (*types.ModelDes, error) {

	des, err := c.factory.Cmdb.GetCollection(ctx, mId)
	if err != nil {
		return nil, err
	}

	return &types.ModelDes{
		ObjectId:   des.ObjectId,
		ObjectName: des.ObjectName,
	}, nil
}

func (c *cmdb) CreateObjClassification(ctx context.Context, obj *types.ObjClassification) error {
	objClassification := &model.ObjClassification{
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
		ClassificationId:   obj.ClassificationId,
		ClassificationName: obj.ClassificationName,
	}

	if err := c.factory.Cmdb.CreateObjClassification(ctx, *objClassification); err != nil {
		return err
	}

	return nil
}

func (c *cmdb) CreateObject(ctx context.Context, obj *types.ObjectDes) error {
	objectDes := &model.ObjectDes{
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
		ObjectId:         obj.ObjectId,
		ObjectName:       obj.ObjectName,
		ClassificationId: obj.ClassificationId,
	}

	if err := c.factory.Cmdb.CreateObjDes(ctx, *objectDes); err != nil {
		return err
	}

	return nil
}

func (c *cmdb) CreateObjectAttr(ctx context.Context, obj *types.ObjectAttr) error {
	objectAttr := &model.ObjectAttDes{
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		ObjectId:     obj.ObjectId,
		PropertyId:   obj.PropertyId,
		PropertyType: obj.PropertyType,
	}

	if err := c.factory.Cmdb.CreateObjAttr(ctx, *objectAttr); err != nil {
		return err
	}

	return nil
}

func (c *cmdb) GetResource(ctx context.Context) ([]types.Resource, error) {
	return []types.Resource{
		{Name: "test", Number: 1},
		{Name: "test2", Number: 0},
	}, nil
}

func NewCmdb(f *db.ShareDaoFactory) *cmdb {
	return &cmdb{f}
}

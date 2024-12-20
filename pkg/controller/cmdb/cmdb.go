package cmdb

import (
	"context"
	"github.com/darianJmy/fortis-services/pkg/db"
	"github.com/darianJmy/fortis-services/pkg/db/model"
	"github.com/darianJmy/fortis-services/pkg/types"
)

type CmdbGetter interface {
	CMDB() Interface
}

type Interface interface {
	CreateModel(ctx context.Context, m *types.ModelDes) (*types.ModelDes, error)
	GetModel(ctx context.Context, mId string) (*types.ModelDes, error)
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

func NewCmdb(f *db.ShareDaoFactory) *cmdb {
	return &cmdb{f}
}

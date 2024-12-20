package controller

import (
	"github.com/darianJmy/fortis-services/cmd/app/config"
	"github.com/darianJmy/fortis-services/pkg/controller/cmdb"
	"github.com/darianJmy/fortis-services/pkg/db"
)

type FortisInterface interface {
	cmdb.CmdbGetter
}

type sample struct {
	cc      config.Config
	factory *db.ShareDaoFactory
}

func (s *sample) CMDB() cmdb.Interface { return cmdb.NewCmdb(s.factory) }

func New(cfg config.Config, f *db.ShareDaoFactory) FortisInterface {
	return &sample{
		cc:      cfg,
		factory: f,
	}
}

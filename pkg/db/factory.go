package db

import (
	"github.com/darianJmy/fortis-services/pkg/db/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type ShareDaoFactory struct {
	User     *Curd[*model.User]
	Role     *Curd[*model.Role]
	UserRole *Curd[*model.UserRole]
	Menu     *Curd[*model.Menu]
	RoleMenu *Curd[*model.RoleMenu]

	Enforcer *csEnforcer
	Cmdb     *cmdb
}

func NewDaoFactory(mongoCN *mongo.Database, migrate bool) (*ShareDaoFactory, error) {

	return &ShareDaoFactory{

		Cmdb:     newCmdb(mongoCN),
	}, nil
}

package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	ClassificationId = "classification_id"
	ObjectId         = "object_id"
	PropertyId       = "property_id"
	InstId           = "inst_id"
	InstName         = "实例名称"
	CreatedAt        = "created_at"
	CreatedAtName    = "创建时间"
	UpdatedAt        = "updated_at"
	UpdatedAtName    = "更新时间"
)

type ObjClassification struct {
	ID                 primitive.ObjectID `bson:"id,omitempty"`
	CreatedAt          time.Time          `bson:"created_at,omitempty"`
	UpdatedAt          time.Time          `bson:"updated_at,omitempty"`
	ClassificationId   string             `bson:"classification_id,omitempty"`
	ClassificationName string             `bson:"classification_name,omitempty"`
	ClassificationIcon string             `bson:"classification_icon,omitempty"`
}

func (o ObjClassification) TableName() string { return "ObjClassification" }

type ObjectDes struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt        time.Time          `bson:"created_at,omitempty"`
	UpdatedAt        time.Time          `bson:"updated_at,omitempty"`
	ObjectId         string             `bson:"object_id,omitempty"`
	ObjectName       string             `bson:"object_name,omitempty"`
	Description      string             `bson:"description,omitempty"`
	ClassificationId string             `bson:"classification_id,omitempty"`
}

func (o ObjectDes) TableName() string { return "ObjDes" }

type ObjectAttDes struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
	ObjectId     string             `bson:"object_id"`
	PropertyId   string             `bson:"property_id"`
	PropertyName string             `bson:"property_name"`
	PropertyType string             `bson:"property_type"`
	Editable     bool               `bson:"editable"`
	IsRequired   bool               `bson:"is_required"`
}

func (o ObjectAttDes) TableName() string { return "ObjAttDes" }

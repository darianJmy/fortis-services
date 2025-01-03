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
	AsstId           = "asst_id"
	ObjAsstId        = "obj_asst_id"
	SrcObjId         = "src_obj_id"
)

type ObjClassification struct {
	ID                 primitive.ObjectID `bson:"id,omitempty"`
	ClassificationId   string             `bson:"classification_id,omitempty"`
	ClassificationName string             `bson:"classification_name,omitempty"`
	ClassificationIcon string             `bson:"classification_icon,omitempty"`
	CreatedAt          time.Time          `bson:"created_at,omitempty"`
	UpdatedAt          time.Time          `bson:"updated_at,omitempty"`
}

func (o ObjClassification) TableName() string { return "ObjClassification" }

type ObjectDes struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	ObjectId         string             `bson:"object_id,omitempty"`
	ObjectName       string             `bson:"object_name,omitempty"`
	ClassificationId string             `bson:"classification_id,omitempty"`
	Description      string             `bson:"description,omitempty"`
	CreatedAt        time.Time          `bson:"created_at,omitempty"`
	UpdatedAt        time.Time          `bson:"updated_at,omitempty"`
}

func (o ObjectDes) TableName() string { return "ObjDes" }

type ObjectAttDes struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	ObjectId     string             `bson:"object_id,omitempty"`
	PropertyId   string             `bson:"property_id,omitempty"`
	PropertyName string             `bson:"property_name,omitempty"`
	PropertyType string             `bson:"property_type,omitempty"`
	Editable     bool               `bson:"editable,omitempty"`
	IsRequired   bool               `bson:"is_required,omitempty"`
	CreatedAt    time.Time          `bson:"created_at,omitempty"`
	UpdatedAt    time.Time          `bson:"updated_at,omitempty"`
}

func (o ObjectAttDes) TableName() string { return "ObjAttDes" }

type AsstDes struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	AsstId    string             `bson:"asst_id,omitempty"`
	SrcDes    string             `bson:"src_des,omitempty"`
	DestDes   string             `bson:"dest_des,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
}

func (o AsstDes) TableName() string { return "AsstDes" }

type ObjAsstDes struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ObjAsstId string             `bson:"obj_asst_id,omitempty"`
	SrcObjId  string             `bson:"src_obj_id,omitempty"`
	DestObjId string             `bson:"dest_obj_id,omitempty"`
	AsstId    string             `bson:"asst_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
}

func (o ObjAsstDes) TableName() string { return "ObjAsstDes" }

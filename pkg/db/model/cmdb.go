package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ObjClassification struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt          time.Time          `bson:"created_at"`
	UpdatedAt          time.Time          `bson:"updated_at"`
	ClassificationId   string             `bson:"classification_id,omitempty"`
	ClassificationName string             `bson:"classification_name"`
}

func (o *ObjClassification) TableName() string { return "ObjClassification" }

type ObjectDes struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt        time.Time          `bson:"created_at"`
	UpdatedAt        time.Time          `bson:"updated_at"`
	ObjectId         string             `bson:"object_id"`
	ObjectName       string             `bson:"object_name"`
	Description      string             `bson:"description"`
	ClassificationId string             `bson:"classification_id"`
}

func (o *ObjectDes) TableName() string { return "ObjDes" }

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

func (o *ObjectAttDes) TableName() string { return "ObjAttDes" }

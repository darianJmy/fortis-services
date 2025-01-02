package types

type ObjClassification struct {
	ClassificationId   string `json:"classification_id"`
	ClassificationName string `json:"classification_name"`
	ClassificationIcon string `json:"classification_icon"`
}

type ObjectDes struct {
	ObjectId         string `json:"object_id"`
	ObjectName       string `json:"object_name"`
	ClassificationId string `json:"classification_id"`
}

type ObjectAttr struct {
	PropertyId   string `json:"property_id"`
	PropertyName string `json:"property_name"`
	PropertyType string `json:"property_type"`
	ObjectId     string `json:"object_id"`
}

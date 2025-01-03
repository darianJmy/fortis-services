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
	Description      string `json:"description"`
}

type ObjectAttr struct {
	PropertyId   string `json:"property_id"`
	PropertyName string `json:"property_name"`
	PropertyType string `json:"property_type"`
	ObjectId     string `json:"object_id"`
}

type AssociationType struct {
	AsstId   string `json:"asst_id"`
	AsstName string `json:"asst_name"`
	SrcDes   string `json:"src_des"`
	DestDes  string `json:"dest_des"`
}

type ObjAsstDes struct {
	ObjAsstId string `json:"obj_asst_id"`
	SrcObjId  string `json:"src_obj_id"`
	DestObjId string `json:"dest_obj_id"`
	AsstId    string `json:"asst_id"`
}

type InstAsstDes struct {
	ObjAsstId string `json:"obj_asst_id"`
	SrcObjId  string `json:"src_obj_id"`
	DestObjId string `json:"dest_obj_id"`
	AsstId    string `json:"asst_id"`
}

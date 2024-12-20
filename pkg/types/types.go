package types

type User struct {
	Id          int64  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Password    string `json:"password,omitempty"`
	Status      int64  `json:"status,omitempty"`
	Email       string `json:"email,omitempty"`
	Description string `json:"description,omitempty"`
}

type Password struct {
	OriginPassword  string `json:"origin_password"`
	CurrentPassword string `json:"current_password"`
}

type Role struct {
	Id          int64  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type Roles struct {
	RoleIds []int64 `json:"role_ids"`
}

type Menu struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	Method      string `json:"method"`
	MenuType    int64  `json:"menu_type"`
	Description string `json:"description,omitempty"`
}

type Menus struct {
	MenuIds []int64 `json:"menu_ids"`
}

type ModelDes struct {
	ObjectId   string `json:"object_id"`
	ObjectName string `json:"object_name"`
}

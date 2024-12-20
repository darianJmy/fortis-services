package model

import (
	"gorm.io/gorm"
)

// User 用户
type User struct {
	gorm.Model

	Name        string `gorm:"column:name;index:idx_name,unique;not null;" json:"name" form:"name"`         // 用户名
	Password    string `gorm:"column:password;type:varchar(256);not null;" json:"password" form:"password"` // 密码
	Status      int8   `gorm:"column:status;type:tinyint(1);not null;" json:"status" form:"status"`         // 状态 1 启用,2 不启用
	Email       string `gorm:"column:email;type:varchar(128);not null;" json:"email" form:"email"`          // 邮件
	Description string `gorm:"column:description;type:text" json:"description" form:"description"`          // 描述
}

func (user *User) TableName() string { return "users" }

// Role 角色
type Role struct {
	gorm.Model

	Name        string `gorm:"column:name;index:idx_role_name,unique;not null" json:"name" form:"name"` // 名称
	Description string `gorm:"column:description;type:text" json:"description" form:"description"`      // 描述
}

func (r *Role) TableName() string { return "roles" }

// Menu 菜单
type Menu struct {
	gorm.Model

	Name        string `gorm:"column:name;size:128;not null;" json:"name" form:"name"`                       // 菜单名称
	URL         string `gorm:"column:url;size:128;not null" json:"url,omitempty" form:"url"`                 // 菜单URL
	Method      string `gorm:"column:method;size:32;not null;" json:"method,omitempty" form:"method"`        // 操作类型 none/GET/POST/PUT/DELETE
	MenuType    int8   `gorm:"column:menu_type;type:tinyint(1);not null;" json:"menu_type" form:"menu_type"` // 菜单类型 1 左侧菜单,2 按钮, 3 非展示权限
	Description string `gorm:"column:description;type:text" json:"description" form:"description"`           // 描述

	Children []Menu `gorm:"-" json:"children"`
}

func (m *Menu) TableName() string {
	return "menus"
}

// UserRole 用户绑定角色
type UserRole struct {
	gorm.Model

	UserID int64 `gorm:"column:user_id;unique_index:uk_user_role_user_id;not null;" json:"user_id"` // 管理员ID
	RoleID int64 `gorm:"column:role_id;unique_index:uk_user_role_user_id;not null;" json:"role_id"` // 角色ID
}

func (u *UserRole) TableName() string { return "user_roles" }

// RoleMenu 角色绑定菜单
type RoleMenu struct {
	gorm.Model

	RoleID int64 `gorm:"column:role_id;unique_index:uk_role_menu_role_id;not null;" json:"role_id"`  // 角色ID
	MenuID int64 `gorm:"column:menu_id;unique_index:uk_role_menu_role_id;not null;" json:"menu_id'"` // 菜单ID
}

func (m *RoleMenu) TableName() string {
	return "role_menus"
}

// Rule 规则，由 casbin 控制
type Rule struct {
	gorm.Model

	PType  string `json:"ptype" gorm:"column:ptype;size:100" description:"策略类型"`
	Role   string `json:"role" gorm:"column:v0;size:100" description:"角色"`
	Path   string `json:"path" gorm:"column:v1;size:100" description:"api路径"`
	Method string `json:"method" gorm:"column:v2;size:100" description:"访问方法"`
	V3     string `gorm:"column:v3;size:100"`
	V4     string `gorm:"column:v4;size:100"`
	V5     string `gorm:"column:v5;size:100"`
}

func (r *Rule) TableName() string { return "rules" }

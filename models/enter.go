package models

import "time"

type Model struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at" comment:"创建时间"`
	UpdatedAt time.Time `json:"updated_at" comment:"更新时间"`
}

type UserModel struct {
	Model
	Username string      `gorm:"size:64;unique" json:"username" comment:"用户名"`
	Nickname string      `gorm:"size:64;" json:"nickname" comment:"昵称"`
	Avatar   string      `gorm:"size:256;" json:"avatar" comment:"头像"`
	Email    string      `gorm:"size:128;" json:"email" comment:"邮箱"`
	Password string      `gorm:"size:64" json:"password" comment:"密码"`
	IsAdmin  bool        `gorm:"default:false" json:"is_admin" comment:"是否管理员"`
	RoleList []RoleModel `gorm:"many2many:user_role_models; joinForeignKey:UserID; joinReferences:RoleID;" json:"roleList" comment:"角色列表"`
}

func (u *UserModel) GetRoleList() []uint {
	var roleList []uint
	for _, model := range u.RoleList {
		roleList = append(roleList, model.ID)
	}
	return roleList
}

type RoleModel struct {
	Model
	Title       string      `gorm:"size:16;unique" json:"title" comment:"角色名称"`
	Description string      `gorm:"size:256;" json:"description" comment:"角色描述"`
	UserList    []UserModel `gorm:"many2many:user_role_models; joinForeignKey:RoleID; joinReferences:UserID;" json:"userList" comment:"用户列表"`
	MenuList    []MenuModel `gorm:"many2many:role_menu_models; joinForeignKey:RoleID; joinReferences:MenuID;" json:"menuList" comment:"菜单列表"`
}

type UserRoleModel struct {
	Model
	UserID    uint      `gorm:"index" json:"user_id" comment:"用户ID"`
	UserModel UserModel `gorm:"foreignKey:UserID;references:ID" json:"user" comment:"用户"`
	RoleID    uint      `gorm:"index" json:"role_id" comment:"角色ID"`
	RoleModel RoleModel `gorm:"foreignKey:RoleID;references:ID" json:"role" comment:"角色"`
}

type Meta struct {
	Icon  string `gorm:"size:128" json:"icon" comment:"图标"`
	Title string `gorm:"size:64" json:"title" comment:"标题"`
}

type MenuModel struct {
	Model
	Name            string `gorm:"size:64;unique" json:"name" comment:"菜单名称"`
	Path            string `gorm:"size:128" json:"path" comment:"菜单路径"`
	Component       string `gorm:"size:128" json:"component" comment:"组件路径"`
	Meta            `gorm:"embedded" json:"meta" comment:"元数据"`
	ParentMenuID    *uint        `gorm:"index" json:"ParentMenuID" comment:"父菜单ID"`
	ParentMenuModel *MenuModel   `gorm:"foreignKey:ParentMenuID;references:ID" json:"-" comment:"父菜单"`
	Children        []*MenuModel `gorm:"foreignKey:ParentMenuID;references:ID" json:"children" comment:"子菜单"`
	Sort            int          `json:"sort" comment:"排序"`
}

type APIModel struct {
	Model
	Name        string `gorm:"size:64;unique" json:"name" comment:"API名称"`
	Path        string `gorm:"size:128" json:"path" comment:"API路径"`
	Method      string `gorm:"size:16" json:"method" comment:"请求方法"`
	Group       string `gorm:"size:64" json:"group" comment:"API分组"`
	Description string `gorm:"size:256" json:"description" comment:"API描述"`
}

type RoleMenuModel struct {
	Model
	RoleID    uint      `gorm:"index" json:"role_id" comment:"角色ID"`
	RoleModel RoleModel `gorm:"foreignKey:RoleID;references:ID" json:"role" comment:"角色"`
	MenuID    uint      `gorm:"index" json:"menu_id" comment:"菜单ID"`
	MenuModel MenuModel `gorm:"foreignKey:MenuID;references:ID" json:"menu" comment:"菜单"`
}

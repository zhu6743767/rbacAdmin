package user_api

import (
	"rbacAdmin/common/query"
	"rbacAdmin/common/resp"
	"rbacAdmin/middleware"
	"rbacAdmin/models"

	"github.com/gin-gonic/gin"
)

type UserListRequest struct {
	models.Page
	Role     uint   `form:"role" json:"role" comment:"角色ID"`
	Username string `form:"username" json:"username" comment:"用户名"`
	Email    string `form:"email" json:"email" comment:"邮箱"`
}

type UserListResponse struct {
	models.UserModel
}

func (UserApi) UserListView(c *gin.Context) {
	cr := middleware.GetBind[UserListRequest](c)

	// var list = make([]models.UserModel, 0)

	// // 分页
	// offset := (cr.Page.Page - 1) * cr.Limit

	// global.DB.Preload("RoleList").Where(models.UserModel{
	// 	Username: cr.Username,
	// 	Email:    cr.Email,
	// }).Limit(cr.Limit).Offset(offset).Find(&list)

	// var count int64
	// global.DB.Model(&models.UserModel{}).Where(models.UserModel{
	// 	Username: cr.Username,
	// 	Email:    cr.Email,
	// }).Count(&count)

	list, count, _ := query.List(models.UserModel{
		Username: cr.Username,
		Email:    cr.Email,
	}, query.Option{
		Page:  cr.Page,
		Debug: true,
		Likes: []string{
			"Username",
			"Nickname",
			"Email",
		},
		Preloads: []string{
			"RoleList",
		},
	})

	resp.OkWithList("用户列表", list, count, c)
}

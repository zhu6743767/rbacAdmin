package user_api

import (
	"rbacAdmin/common/resp"
	"rbacAdmin/global"
	"rbacAdmin/middleware"
	"rbacAdmin/models" // 添加这行导入
	"rbacAdmin/utils/jwts"
	"rbacAdmin/utils/pwd"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 定义登录请求结构体
type LogingRequest struct {
	Username string `json:"username" binding:"required" label:"用户名"`
	Password string `json:"password" binding:"required" label:"密码"`
}

// 定义登录响应结构体
type LoginResponse struct {
	Token string `json:"token"`
}

// LoginView 登录视图
func (u *UserApi) LoginView(c *gin.Context) {
	// var cr LogingRequest
	// err := c.ShouldBindJSON(&cr)
	// if err != nil {
	// 	//c.JSON(200, gin.H{"code": 1001, "msg": err.Error(), "data": nil})
	// 	resp.FailWithBindingError(err, c)
	// 	return
	// }

	// cr, err := BindJson[LogingRequest](c)
	// if err != nil {
	// 	return
	// }

	cr := middleware.GetBind[LogingRequest](c)

	var user models.UserModel
	err := global.DB.Preload("RoleList").Take(&user, "username = ?", cr.Username).Error
	if err != nil {
		//c.JSON(200, gin.H{"code": 1001, "msg": err.Error(), "data": nil})
		resp.FailWithError(err, c)
		return
	}
	if !pwd.ComparePassword(user.Password, cr.Password) {
		//c.JSON(200, gin.H{"code": 1001, "msg": "用户名或密码错误", "data": nil})
		resp.FailWithMsg("用户名或密码错误", c)
		return
	}

	var roleList []uint
	for _, model := range user.RoleList {
		roleList = append(roleList, model.ID)
	}

	token, err := jwts.GetToken(jwts.ClaimsUserInfo{
		UserID:   user.ID,
		Username: user.Username,
		RoleList: roleList,
	})
	if err != nil {
		logrus.Error("jwt生成token失败: ", err)
		//c.JSON(200, gin.H{"code": 1001, "msg": "用户登录失败", "data": nil})
		resp.FailWithError(err, c)
		return
	}
	//c.JSON(200, gin.H{"code": 200, "msg": "登录成功", "data": LoginResponse{
	//resp.OkWithMsg("登录成功", c)
	resp.OkWithData(LoginResponse{
		Token: token,
	}, c)
}

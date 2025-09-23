package api

import (
	"rbacAdmin/api/captcha_api"
	"rbacAdmin/api/email_api"
	"rbacAdmin/api/user_api"
)

type Api struct {
	UserApi    *user_api.UserApi
	CaptchaApi *captcha_api.CaptchaApi
	EmailApi   *email_api.EmailApi
}

var App = new(Api)

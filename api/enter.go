package api

import (
	"rbacAdmin/api/captcha_api"
	"rbacAdmin/api/email_api"
	"rbacAdmin/api/image_api"
	"rbacAdmin/api/user_api"
)

type Api struct {
	UserApi    *user_api.UserApi
	CaptchaApi *captcha_api.CaptchaApi
	EmailApi   *email_api.EmailApi
	ImageApi   *image_api.ImageApi
}

var App = &Api{
	UserApi:    &user_api.UserApi{},
	CaptchaApi: &captcha_api.CaptchaApi{},
	EmailApi:   &email_api.EmailApi{},
	ImageApi:   &image_api.ImageApi{},
}

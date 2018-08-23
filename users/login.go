package users

import "github.com/aosfather/bingo/mvc"

/**
  登录
 */
 type loginRequest struct {

 }
type LoginController struct {
	mvc.SimpleController
}

func (this *LoginController)GetUrl() string {
	return "/login"
}

func (this *LoginController) GetParameType(method string) interface{} {
	return &loginRequest{}

}

func (this *LoginController) Post(c mvc.Context, p interface{}) (interface{}, mvc.BingoError) {

	return nil,nil
}



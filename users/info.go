package users

import "github.com/aosfather/bingo/mvc"

//基本信息


//注册用户
type registerRequest  struct {
  Name string   //名称   ---当使用其它登录系统时候可以不填
  Mobile string //手机号  ---当使用其它登录系统时候可以不填
  OtherLoginName string //其它系统的登录用户名
  OtherLoginSystem string//其它登录系统名称
  Source string //来源
  InvitationCode string //邀请码
}

type RegisterController struct {
	mvc.SimpleController
}

func (this *RegisterController)GetUrl() string {
	return "/register"
}

func (this *RegisterController) GetParameType(method string) interface{} {
	return &registerRequest{}

}

func (this *RegisterController) Post(c mvc.Context, p interface{}) (interface{}, mvc.BingoError) {

	return nil,nil
}
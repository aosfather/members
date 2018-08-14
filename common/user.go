package common
/**
  用户信息
 */
type User struct {
	baseInfo
	Id string  //用户编码
	Name string
	Male bool //男性 true，女性false
	BirthDay string //2001-01-01的格式
	LoginName string //登录用户名
	LoginPwd string //登录密码
	OtherLoginName string //其它系统的登录用户名
	OtherLoginSystem string//其它登录系统名称
	Source string //来源
	InvitationCode string //邀请码

}

/**
  证件
 */
type credential struct{
	baseInfo
	UserId string //用户编码
	Type string //证件类型
	Name string //名称
	LastName string //姓
	FirstName string //名
	No string //证件号
	BeginDate string //有效期开始时间
	EndDate string //有效期结束时间
	Authority string //发证机关
	Validation bool //是否验证，true 验证通过
}

/*
  通讯录信息
 */
type Contact struct {
	baseInfo
	UserId string //用户编码
	Catalog string //通讯录类型
	Tags string //标签
	Name string //名称
	Mobile string //手机
	Email string //邮箱
	Address string //联系地址
	PostCode string //邮编
	QQ string //QQ号
	Webchat string //微信号
	Webo string //微博
	Facebook string //
	Twitter string

}

/*
  地址信息

 */
type Address struct {
	baseInfo
	UserId string //用户编码
	Index int64  //序号
	Default bool //默认地址
	PostCode string //邮编
	AreaCode string //区域码
	Detail   string //详细地址
}
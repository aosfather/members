package common

/**
会员卡
1、发卡机构
2、卡名称
3、卡归属
4、卡额度  -1表示不限制额度
5、卡有效期
6、卡适用的品类
7、卡使用的规则
8、卡状态  有效、冻结、过期、作废、未激活

 */

type ConfigData map[string]interface{}  //配置数据
type CardStatus byte //卡状态
const (
	CS_INIT CardStatus=0 //未激活
	CS_ENABLED CardStatus=1 //有效
	CS_FROZEN CardStatus=3  //冻结
	CS_EXPIRED CardStatus=5 //过期
	CS_CANCELED CardStatus=6 //作废
)

//会员卡
type MemberCard struct{
	baseInfo
	No string //卡编号
	Official string //发卡机构
	Name string //卡名称
	UserId string//归属
    Amount int64 //卡额度
    Expire int64 //过期时间  -1表示永不过期
	Categorys string  //卡能适用的品类
    Status CardStatus //卡状态
}


//卡使用规则
type CardUseRule struct {
	baseInfo
	CardNo string//卡编号
	Categorys string //适用的品类  其中 *号表示适用所有品类 x*表示适用于x开头的品类。多个品类适用","号分割 !表示禁止品类
	RuleId string//规则编号
}

type Rule struct {
	baseInfo
	Id string //规则编号
	Name string //规则名称
	Detail string //规则说明
	Catalog string //规则分类
	Creater string //创建者
	ReferCode string //实现引用的规则模板编码
	Data ConfigData   //规则配置数据
}
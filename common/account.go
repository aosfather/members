package common

//基本信息
type baseInfo struct {
	CreateDate string  //创建日期
	LastUpdate string  //最后修改日期

}

//基本账户
type _account struct {
	baseInfo
	Code string //账户编号
	Owner string //账户归属
	Catalog string //账户分类
	Amount int64   //账户余额


}



//积分账户
type PointAccount struct {
	_account

}

//账户流水
type FlowWater struct {

}
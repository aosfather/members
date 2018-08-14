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

//账户流水
type FlowWater struct {
	Id     string `Table:"t_flow_details" Option:"auto" json:"id"` //流水id
	Account   string `json:"account"`                                    //账户id
	Detail string `json:"detail"`                                        //描述
	Amount int64  `json:"amount"`                                        //发生数
	Date   string `Field:"last_update" json:"date"`                      //时间
	Type   byte `Field:"op_type" json:"type"`                          //操作类型
	Begin  int64  `Field:"the_begin" json:"begin"`                       //期初
	End    int64  `Field:"the_end" json:"end"`                           //期末
	Extra  string `json:"extra"`                                       //额外信息，比如发起系统等
}

//积分账户
type PointAccount struct {
	_account
	CountAmount     int64  `json:"count" Field:"sum_amount"`   //获取的总的积分
	ExchangedAmount int64  `json:"exchanged"`                  //已经兑换的积分
}

/**
优惠券的种类：

    满减券：如满100-10，满2件减10元
    折扣券：如满100打9折

优惠券相关概念：

    规则定义：涉及计算的部分，作用范围：商品、商户、类目，计算方式：是否累加，是否互斥等
    优惠券定义、模板：发放有效期，使用有效期，使用渠道，发放渠道，券类型：商户券，平台券
    优惠券：发放到用户手中的优惠券

优惠券用法：

1.    发券：业务方调发券接口发券，运营后台为指定用户批量导券，发券常用于价值比较高的券

    领券：由运营创建一个优惠券模板， 配置在活动页或者商品详情页上，让用户主动领券，领券常用于价值比较低的券，用户自己就能领到

3.    用券：消费者收到优惠券后，在结算页从优惠券列表中选择优惠券并使用.

优惠券的使用流程：

    配置券定义（券模板）
    运营给用户发券/用户主动领券
    用户使用券
    核销，退券等
    优惠券数据统计：通过读取线上优惠券发放与使用的记录，数据分析团队创建数据报表，提供给大数据团队进行分析，通过分析，优惠券发放也可以做到千人千面

优惠券的发放限制：

    具体业务方，比如激活中心，新手券，下单券，用户生日券
    具体的发放时间：如2017-09-01到2017-09-10
    具体活动，如双十一
    具体地点，如上海
    具体的人群，如男性

优惠券的使用限制：

    门槛
    状态
    有效期
    业务渠道

有效期的三种类型：

    模板固定有效期
    模板动态有效期
    发券时指定有效期

发券方式:

    同步
    异步

发放的数量限制:

    券库存限制
    用户总数限制
    用户每天限制

通知方式:

    短信通知
    邮件通知
    app消息通知

报警：

    优惠券库存告警
    优惠券使用时间告警

主要业务方：

1.    用户端：商详页领券，店铺页领券，活动页领券

2.    后台系统：批量发放优惠券

3.    抽奖发券

4.    新用户注册返券

5.    下单返券

    评价返券

非功能特性：

    幂等性：发券，使用券要保证幂等性，领券不需要。
    高性能：活动页领券等业务场景必须拥有高性能，服务可水平扩展
    可扩展性：对于新的优惠券方式实现快速扩展，对于外部渠道的优惠券可以快速对接


 */

//优惠劵模板账户
type CouponTemplateAccount struct {
	_account
	CouponTemplate

}
//优惠券模板
type CouponTemplate struct {
	Code string
	//发放有效期，使用有效期，使用渠道，发放渠道，券类型：商户券，平台券


}

//用户领到的优惠劵
type Coupon struct {

}
package trade

import "github.com/aosfather/members/common"

/**
  订单结算--面向用户
 */
type Settle struct {
 Id string //结算单号
 TradeIds string //订单号使用;分割
 User string //用户
 Amount common.Money       //票面金额
 DiscountAmount common.Money //折扣金额
 CouponAmount common.Money //优惠劵优惠抵扣金额
 RealAmount common.Money   //实付金额
 Status byte //待支付，已支付，已取消

}
 //结算条目
 type SettleItem struct {

 }

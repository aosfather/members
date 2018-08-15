package trade


/**
  交易模型
  1、进入购物车
  2、下订单
  3、生成结算单->收费部分，优惠劵，打折卡等
  4、支付结算单
  5、交易成功，进入出库后续流程
  ----------------------------------
  退货、换货等为独立流程，只对订单状态产生影响
 */



 //订单
 type TradeOrder struct {
 	Id string //订单编号


 	Status OrderStatus
 	PayStatus string

 }



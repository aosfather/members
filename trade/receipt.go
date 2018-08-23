package trade

import (
	"github.com/aosfather/members/common"
	"github.com/aosfather/bingo/mvc"
)

//小票
type Receipt struct {
	Merchant string  //商家id
	Customer string  //客户
	Terminal string  //终端地点
    Date string //打印日期
    Items []ReceiptItem //条目
    Amount common.Money //总金额
    Number int64 //总件数
}



/**
  小票条目
 */
type ReceiptItem struct {
	Code string //商品编码
	Name string //商品名称
	Price common.Money //商品价格
	Number common.UnitNumber//商品数量
	Style string //规格描述
	Amount common.Money //总金额

}

//获取小票列表
type QueryReceiptController struct {
  mvc.SimpleController
}
func (this *QueryReceiptController)GetUrl() string {
	return "/query_receipt"
}

func (this *QueryReceiptController) GetParameType(method string) interface{} {
	return nil

}

func (this *QueryReceiptController) Get(c mvc.Context, p interface{}) (interface{}, mvc.BingoError) {

	return nil,nil
}

//小票
type ReceiptController struct {
   mvc.SimpleController
}

func (this *ReceiptController)GetUrl() string {
	return "/receipt"
}

func (this *ReceiptController) GetParameType(method string) interface{} {
	return nil

}

func (this *ReceiptController) Post(c mvc.Context, p interface{}) (interface{}, mvc.BingoError) {

	return nil,nil
}

func (this *ReceiptController) Get(c mvc.Context, p interface{}) (interface{}, mvc.BingoError) {

	return nil,nil
}
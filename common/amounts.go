package common

import "fmt"

//数量相关的类型
type Money int64         //单位分
type UnitNumber int64    //单位数量


func (this Money)ToYuan() string{
	return fmt.Sprintf("%.2f",float64(this)/float64(100))
}

//单价*件数
func (this *Money)SumByNumber(n UnitNumber){
	*this=Money(int64(*this)*int64(n))
}

func (this Money)ToSumByNumber(n UnitNumber) Money{
	return Money(int64(this)*int64(n))
}

func (this Money)ToDiscount(discounts... Money) Money{
	result:=int64(this)
	for _,d:=range discounts {
		result-=int64(d)
	}
	return Money(result)
}

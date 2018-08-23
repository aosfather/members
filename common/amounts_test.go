package common

import "testing"

func TestMoney_ToYuan(t *testing.T) {
	var m Money=12
	t.Log(m.ToYuan())

	m.SumByNumber(100)
	t.Log(m.ToYuan())
	t.Log(m)
	t.Log(m.ToSumByNumber(10))
	t.Log(m)
	t.Log(m.ToDiscount(10,1,2,8))
}

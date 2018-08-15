package common


type DetailType map[string]string //详情类型

/**
 商品
 */
type baseCodeObject struct {
	Id string
	Name string //名称
	Code string //编码
}
 //商品分类
 type SPCatalog struct {
 	baseCodeObject

 }

 //商品规格
 type SPStand struct {

 }

 //商品品牌
 type SPBand struct {

 }

 //商品
 type SPU struct {
	 baseCodeObject
	 Catalog string //分类id
	 Band string //品牌id
	 //简介
	 Description string
	 //简介图
	 CoverImageUrl  string //图片地址
	 //详情
     Detail DetailType
 }



 //属性
 type Property struct{
 	baseCodeObject
 	Catalog string

 }

 //属性取值
 type PropertyValue struct {
 	Id int64
 	PropertyId string //属性ID
 	Value string //选项值

 }

 //商品库存规格
 type SKU struct {
	 baseCodeObject
	 Price int64 //价格
     StoreSnap int64 //库存快照
	 SPUId string   //商品编码
 }

//规格属性取值，一个规格多个取值。如 红色、电信版、IPhone6
 type SKUPropertyValue struct {
 	 Id int64
     SKUId string  //规格编码
	 ProertyValueId int64 //属性取值ID
 }





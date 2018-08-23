package product

import "github.com/aosfather/bingo/mvc"

/**
 商品库
 1、给定的查询条件对指定的商品库进行查询
 2、
 */

 type searchRequest struct {
 	RequestNo string //查询号
 	Catalog string //品类
 	Param map[string]string
 	Filter queryFilterOption
 	OrderModel string
 }

type queryFilterOption struct {
	Type   int      `json:"type"`
	Values []string `json:"values"`
}

 type SearchController struct {
 	mvc.SimpleController

 }

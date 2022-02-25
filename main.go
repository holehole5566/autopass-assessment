package main

import (
	"test/model"
	"test/service"
)

var Calculator service.Calculator
var Promotion1 service.Promotion1
var Promotion2 service.Promotion2
var Promotion3 service.Promotion3
var Promotion4 service.Promotion4

func init() {
	Calculator = service.NewCalculator()
	//設定優惠參數
	//訂單滿 X 元折 Z %,每人只能折N
	Promotion1 = service.NewPromotion1(100, 10, 80)
	//特定商品滿 X 件折 Y 元
	Promotion2 = service.NewPromotion2(2, 100, 1)
	//訂單滿 X 元贈送特定商品
	Promotion3 = service.NewPromotion3(500, 4)
	//訂單滿 X 元折 Y 元,此折扣在全站總共只能套用 N1 次,元,在全站每個月折扣上限為 N2 元
	Promotion4 = service.NewPromotion4(300, 150, 2, 200)
	
	//要用的優惠丟進Calculator
	Calculator.SetPromotion(4)

}

func main() {

	//新增User
	user1 := model.NewUser(1)
	user2 := model.NewUser(2)
	//建立Order
	order1 := model.NewOrder(user1)
	order2 := model.NewOrder(user1)
	order3 := model.NewOrder(user2)
	//新增商品
	product1 := model.NewProduct(1, 500)

	//組合Order
	order1.ProductList = append(order1.ProductList, product1)
	order2.ProductList = append(order2.ProductList, product1)
	order3.ProductList = append(order3.ProductList, product1)

	//檢查折扣&顯示結果1
	Calculator.CheckPromotion(order1, Promotion1, Promotion2, Promotion3, Promotion4)

	//檢查折扣&顯示結果2
	Calculator.CheckPromotion(order2, Promotion1, Promotion2, Promotion3, Promotion4)

	//檢查折扣&顯示結果3
	Calculator.CheckPromotion(order3, Promotion1, Promotion2, Promotion3, Promotion4)

}

package service

import (
	"fmt"
	"strconv"
	"test/model"
)

type Calculator struct {
	PromotionList []int
}

func NewCalculator() Calculator {
	calculator := Calculator{}
	return calculator
}

type Promotion1 struct {
	X int
	Z int
	N int
}

type Promotion2 struct {
	X   int
	Y   int
	PID int
}

type Promotion3 struct {
	X   int
	PID int
}

type Promotion4 struct {
	X  int
	Y  int
	N1 int
	N2 int
}

var Promotion4UsedCount int
var Promotion4UsedAmount int

func NewPromotion1(x int, z int, n int) Promotion1 {
	promotion1 := Promotion1{
		X: x,
		Z: z,
		N: n,
	}
	return promotion1
}

func NewPromotion2(x int, y int, pid int) Promotion2 {
	promotion2 := Promotion2{
		X:   x,
		Y:   y,
		PID: pid,
	}
	return promotion2
}

func NewPromotion3(x int, pid int) Promotion3 {
	promotion3 := Promotion3{
		X:   x,
		PID: pid,
	}
	return promotion3
}

func NewPromotion4(x int, y int, n1 int, n2 int) Promotion4 {
	promotion4 := Promotion4{
		X:  x,
		Y:  y,
		N1: n1,
		N2: n2,
	}
	return promotion4
}

func (p *Promotion1) Discount1(order *model.Order, c *Calculator) *model.Order {

	sum := c.GetToTalPrice(order)
	order.OriginMoney = sum
	if sum >= p.X && order.User.Discounted < p.N {
		fmt.Printf("適用優惠1:滿%d元折%d%% 每人最多折扣%d元\n", p.X, p.Z, p.N)
		if (sum*p.Z)/100 > p.N-order.User.Discounted {
			order.DiscountMoney += p.N - order.User.Discounted
			order.User.Discounted += p.N - order.User.Discounted
		} else {
			order.DiscountMoney += (sum * p.Z) / 100
			order.User.Discounted += (sum * p.Z) / 100
		}
	}

	return order
}

func (p *Promotion2) Discount2(order *model.Order, c *Calculator) *model.Order {

	count := 0
	for _, v := range order.ProductList {
		if v.ID == p.PID {
			count++
		}
	}
	sum := c.GetToTalPrice(order)
	order.OriginMoney = sum
	if count >= p.X {
		fmt.Printf("適用優惠2:特定商品%d滿%d個折%d元\n", p.PID, p.X, p.Y)

		order.DiscountMoney += p.Y
	}
	return order
}

func (p *Promotion3) Discount3(order *model.Order, c *Calculator) *model.Order {

	sum := c.GetToTalPrice(order)
	order.OriginMoney = sum
	if sum >= p.X {
		fmt.Printf("適用優惠3:訂單滿%d元送特定商品%d\n", p.X, p.PID)

		order.DiscountMoney += 0
		product := model.NewProduct(p.PID, 0)
		order.ProductList = append(order.ProductList, product)
	}
	return order
}

func (p *Promotion4) Discount4(order *model.Order, c *Calculator) *model.Order {

	sum := c.GetToTalPrice(order)
	order.OriginMoney = sum
	if sum >= p.X && Promotion4UsedAmount < p.N2 && Promotion4UsedCount <= p.N1 {
		fmt.Printf("適用優惠4:訂單滿%d元折%d元,此折扣在全站總共只能套用%d次,在全站每個月折扣上限為%d元\n", p.X, p.Y, p.N1, p.N2)
		Promotion4UsedCount++
		if p.Y > p.N2-Promotion4UsedAmount {
			order.DiscountMoney += p.N2 - Promotion4UsedAmount
			Promotion4UsedAmount += order.DiscountMoney
		} else {
			order.DiscountMoney += p.Y
			Promotion4UsedAmount += order.DiscountMoney
		}

	}
	return order
}

func (c *Calculator) CheckPromotion(order *model.Order, p1 Promotion1, p2 Promotion2, p3 Promotion3, p4 Promotion4) *model.Order {

	if contains(c.PromotionList, 1) {
		p1.Discount1(order, c)
	}

	if contains(c.PromotionList, 2) {

		p2.Discount2(order, c)
	}

	if contains(c.PromotionList, 3) {

		p3.Discount3(order, c)
	}

	if contains(c.PromotionList, 4) {
		p4.Discount4(order, c)
	}

	allproduct := c.GetAllProducts(order)
	fmt.Println("得到商品:")
	fmt.Println(allproduct)
	fmt.Println("原始金額:")
	fmt.Println(order.OriginMoney)
	fmt.Println("折扣金額:")
	fmt.Println(order.DiscountMoney)
	fmt.Println("最後金額:")
	fmt.Println(order.OriginMoney - order.DiscountMoney)

	return order
}

func (c *Calculator) SetPromotion(i int) {

	c.PromotionList = append(c.PromotionList, i)
}

func (c *Calculator) GetToTalPrice(order *model.Order) int {

	sum := 0
	for _, v := range order.ProductList {
		sum += v.Price
	}
	return sum

}

func (c *Calculator) GetAllProducts(order *model.Order) string {

	all := ""
	for _, v := range order.ProductList {
		all += strconv.Itoa(v.ID) + ";"
	}
	return all

}

func contains(list []int, p int) bool {
	for _, a := range list {
		if a == p {
			return true
		}
	}
	return false
}

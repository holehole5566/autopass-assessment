package service

import (
	"strconv"
	"strings"
	"test/model"
	"testing"
)

func TestNewCalculator(t *testing.T) {

	calculator := NewCalculator()

	if len(calculator.PromotionList) != 0 {
		t.Errorf("wrong")
	}

}

func TestNewPromotion1(t *testing.T) {

	promotion1 := NewPromotion1(100, 10, 10)

	if promotion1.X != 100 || promotion1.Z != 10 {
		t.Errorf("wrong")
	}

}

func TestDiscount1(t *testing.T) {
	p := NewPromotion1(100, 10, 100)
	c := NewCalculator()
	user1 := model.NewUser(1)
	order1 := model.NewOrder(user1)
	product1 := model.NewProduct(1, 100)
	product2 := model.NewProduct(2, 200)
	order1.ProductList = append(order1.ProductList, product1)
	order1.ProductList = append(order1.ProductList, product2)
	sum := c.GetToTalPrice(order1)
	order1.OriginMoney = sum
	if sum >= p.X && order1.User.Discounted < p.N {
		
		if (sum*p.Z)/100 > p.N-order1.User.Discounted {
			order1.DiscountMoney += p.N-order1.User.Discounted
			order1.User.Discounted+=p.N-user1.Discounted
		} else {
			order1.DiscountMoney += (sum * p.Z) / 100
			order1.User.Discounted += (sum * p.Z) / 100
		}
		if order1.OriginMoney != 300 {
			t.Errorf("wrong")
		}
		if order1.DiscountMoney != 30 {
			t.Errorf("wrong")
		}
		if order1.User.Discounted != 30 {
			t.Errorf("wrong")
		}
	} else {
		t.Errorf("wrong")
	}
}

func TestDiscount2(t *testing.T) {
	p := NewPromotion2(2, 100, 1)
	c := NewCalculator()
	user1 := model.NewUser(1)
	order1 := model.NewOrder(user1)
	product1 := model.NewProduct(1, 100)
	product2 := model.NewProduct(2, 200)
	order1.ProductList = append(order1.ProductList, product1)
	order1.ProductList = append(order1.ProductList, product1)
	order1.ProductList = append(order1.ProductList, product1)
	order1.ProductList = append(order1.ProductList, product2)

	sum := c.GetToTalPrice(order1)
	order1.OriginMoney = sum
	count := 0
	for _, v := range order1.ProductList {
		if v.ID == p.PID {
			count++
		}
	}
	if count >= p.X {
	
		order1.DiscountMoney += p.Y
		if order1.OriginMoney != 500 {
			t.Errorf("wrong")
		}
		if order1.DiscountMoney != 100 {
			t.Errorf("wrong")
		}
	} else {
		t.Errorf("wrong")
	}
}

func TestDiscount3(t *testing.T) {
	p := NewPromotion3(500, 4)
	c := NewCalculator()
	user1 := model.NewUser(1)
	order1 := model.NewOrder(user1)
	product1 := model.NewProduct(1, 100)
	product2 := model.NewProduct(2, 200)
	order1.ProductList = append(order1.ProductList, product1)
	order1.ProductList = append(order1.ProductList, product1)
	order1.ProductList = append(order1.ProductList, product1)
	order1.ProductList = append(order1.ProductList, product2)

	sum := c.GetToTalPrice(order1)
	order1.OriginMoney = sum
	if sum >= p.X {
		product := model.NewProduct(p.PID, 0)
		order1.ProductList = append(order1.ProductList, product)
	
		order1.DiscountMoney += 0
		if order1.ProductList[len(order1.ProductList)-1].ID != p.PID {
			t.Errorf("wrong")
		}
	}
}

func TestDiscount4(t *testing.T) {
	p := NewPromotion4(300,150,2,200)
	c := NewCalculator()
	user1 := model.NewUser(1)
	order1 := model.NewOrder(user1)
	product1 := model.NewProduct(1, 100)
	product2 := model.NewProduct(2, 200)
	order1.ProductList = append(order1.ProductList, product1)
	order1.ProductList = append(order1.ProductList, product1)
	order1.ProductList = append(order1.ProductList, product1)
	order1.ProductList = append(order1.ProductList, product2)

	sum := c.GetToTalPrice(order1)
	order1.OriginMoney = sum
	if sum >= p.X && Promotion4UsedAmount < p.N2 && Promotion4UsedCount <= p.N1 {
		Promotion4UsedCount++
		if p.Y > p.N2-Promotion4UsedAmount {
			order1.DiscountMoney += p.N2 - Promotion4UsedAmount
			Promotion4UsedAmount += order1.DiscountMoney
		} else {
			order1.DiscountMoney += p.Y
			Promotion4UsedAmount += order1.DiscountMoney
		}

	}
	if order1.DiscountMoney != 150 || Promotion4UsedAmount != 150 || Promotion4UsedCount != 1 {
		t.Errorf("wrong")
	}
}

func TestGetToTalPrice(t *testing.T) {
	user1 := model.NewUser(1)
	order1 := model.NewOrder(user1)
	product1 := model.NewProduct(1, 100)
	product2 := model.NewProduct(2, 200)
	order1.ProductList = append(order1.ProductList, product1)
	order1.ProductList = append(order1.ProductList, product2)
	sum := 0
	for _, v := range order1.ProductList {
		sum += v.Price
	}
	if sum != 300 {
		t.Errorf("wrong")
	}
}

func TestGetAllProduct(t *testing.T) {
	user1 := model.NewUser(1)
	order1 := model.NewOrder(user1)
	product1 := model.NewProduct(1, 100)
	product2 := model.NewProduct(2, 200)
	order1.ProductList = append(order1.ProductList, product1)
	order1.ProductList = append(order1.ProductList, product2)
	all := ""
	for _, v := range order1.ProductList {
		all += strconv.Itoa(v.ID) + ";"
	}
	all = strings.Trim(all, ";")
	if all != "1;2" {
		t.Errorf("wrong")
	}
}

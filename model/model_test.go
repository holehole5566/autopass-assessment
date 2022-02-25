package model

import (
	"testing"
)

func TestNewUser(t *testing.T) {
    
	var user1 = NewUser(1)
    if user1.ID != 1 || user1.Discounted !=0 {
        t.Errorf("wrong")
    }

}
func TestNewProduct(t *testing.T) {
    
	var product1 = NewProduct(1,100)
    productID:=1
    productPrice:=100
 
    if product1.ID != productID && product1.Price != productPrice {
        t.Errorf("wrong")
    }

}

func TestNewOrder(t *testing.T) {
    
	var product1 = NewProduct(1,100)
    var product2 = NewProduct(2,200)
    var user1=NewUser(1)
    var order1=NewOrder(user1)
 
    order1.ProductList=append(order1.ProductList,product1)
    order1.ProductList=append(order1.ProductList,product2)

    if order1.ProductList[0].ID != product1.ID && order1.ProductList[1].ID != product2.ID {
        t.Errorf("wrong")
    }

}


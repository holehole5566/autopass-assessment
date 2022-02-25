package model

type Order struct {
	User          *User
	OriginMoney   int
	DiscountMoney int
	ProductList   []Product
}

type Product struct {
	ID    int
	Price int
}

type User struct {
	ID         int
	Discounted int
}

func NewUser(id int) *User {
	user := &User{
		ID:         id,
		Discounted: 0,
	}
	return user
}

func NewProduct(id int, price int) Product {
	product := Product{
		ID:    id,
		Price: price,
	}
	return product
}

func NewOrder(user *User) *Order {

	productlist := []Product{}
	order := &Order{
		User:          user,
		ProductList:   productlist,
		OriginMoney:   0,
		DiscountMoney: 0,
	}

	return order
}

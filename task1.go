package main

import (
	"errors"
	"fmt"
)

type Product struct {
	art      string
	title    string
	price    float64
	quantity int
}

func FindProducts(products [5]Product, article string) (*Product, error) {
	for _, product := range products {
		if product.art == article {
			return &product, nil
		}
	}
	return &Product{}, errors.New("not found")
}

func NewProduct(art, title string, price float64, quantity int) Product {
	return Product{
		art:      art,
		title:    title,
		price:    price,
		quantity: quantity,
	}
}

func QuantityCheck(product *Product, quantity int) bool {
	return product.quantity >= quantity
}
func BuyProduct(product *Product, quantity int) error {
	if quantity <= 0 {
		return errors.New("0 нельзя")
	}
	if !QuantityCheck(product, quantity) {
		return errors.New("много")
	}
	product.quantity -= quantity
	fmt.Println(product.price * float64(quantity))
	fmt.Println(product)
	return nil
}

func main() {
	products := [5]Product{
		NewProduct("a1", "keyboard1", 2340, 10),
		NewProduct("a3", "keyboard2", 3400, 9),
		NewProduct("a4", "keyboard3", 25400, 8),
		NewProduct("a45", "keyboard4", 23400, 7),
		NewProduct("a6", "keyboard5", 25010, 6),
	}
	var article string
	var quantity int
	fmt.Println("Enter article numver: ")
	if _, err := fmt.Scanln(&article); err != nil {
		fmt.Println("Wrong num")

	}
	fmt.Println("Enter quantity: ")
	if _, err := fmt.Scanln(&quantity); err != nil {
		fmt.Println("Wrong quantity")
	}
	product, err := FindProducts(products, article)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := BuyProduct(product, quantity); err != nil {
		fmt.Println(err)
		return
	}
}

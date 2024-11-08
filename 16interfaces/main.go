package main

import "fmt"

type Item interface {
	GetPrice() float64
	GetProductName() string
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Book struct {
	ProductName  string
	ProductPrice float64
}

type Rectangular struct {
	Width  float64
	Height float64
}

func (rc *Rectangular) Area() float64 {
	res := rc.Height * rc.Width
	return res
}

func (rc *Rectangular) Perimeter() float64 {
	res := 2 * (rc.Height + rc.Width)
	return res
}

func (b *Book) GetProductName() string {
	return b.ProductName
}

func (b *Book) GetPrice() float64 {
	return b.ProductPrice
}

type Electronic struct {
	ProductName  string
	ProductPrice float64
}

func (e *Electronic) GetProductName() string {
	return e.ProductName
}

func (e *Electronic) GetPrice() float64 {
	return e.ProductPrice
}

func CalculateTotalPrices(items []Item) float64 {
	totalPrice := 0.0

	for _, item := range items {
		totalPrice += item.GetPrice()
	}
	return totalPrice
}

func main() {
	fmt.Println("Learning interface in golang")
	book := Book{"javascript", 90.00}
	electronic := Electronic{"bulb", 45.00}

	items := []Item{&book, &electronic}

	res := CalculateTotalPrices(items)

	fmt.Printf("total price %v \n ", res)
	var shape Shape
	r := Rectangular{
		Width:  5,
		Height: 10,
	}

	shape = &r

	fmt.Printf("The value of area is %f\n", shape.Area())

}

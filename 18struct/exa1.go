package main

import (
	"fmt"
)

// problem is that design a book system and track borrow

type Book struct {
	Id         int
	Author     string
	Name       string
	IsBorrowed bool
}

func (b Book) Borrow() error {
	if b.IsBorrowed {
		return fmt.Errorf("the book %s is already borrowed \n", b.Name)
	}
	b.IsBorrowed = true
	fmt.Printf("book %s is borrowed \n", b.Name)
	return nil
}

func (b Book) Return() error {
	if b.IsBorrowed {
		fmt.Printf("The Book %s is  borrowed\n", b.Name)
	}
	b.IsBorrowed = false
	fmt.Printf("the book %s is returned\n", b.Name)
	return nil
}

func Exa1() error {
	fmt.Println("Practicing on struct")

	book := Book{
		Id:     1,
		Name:   "the deep work",
		Author: "i don't know",
	}

	err := book.Borrow()
	if err != nil {
		return err
	}
	err = book.Return()
	if err != nil {
		return err
	}

	return nil
}

package main

import (
	"fmt"

	"gopractice.com/model"
)

func main() {
	var first *model.Book = CreateBook("One Piece", "Oda", 1030, 1000)
	var second *model.Book = CreateBook("Naruto", "Kishimoto", 1030, 600)

	first.Display()
	second.Display()

	first.Borrow()

	second.ReturnBook()

	SwapTitles(first, second)

}
func CreateBook(a string, b string, c int, d int) *model.Book {

	book := model.Book{Title: a, Author: b, Pages: c, CopiesAvailable: d}
	return &book
}

func SwapTitles(a *model.Book, b *model.Book) {
	var c string
	c = a.Title
	a.Title = b.Title
	b.Title = c
	fmt.Println("Swapped...")
	fmt.Println(a, b)

}

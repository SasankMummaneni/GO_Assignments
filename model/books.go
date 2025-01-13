package model

import "fmt"

type Book struct {
	Title           string
	Author          string
	Pages           int
	CopiesAvailable int
}

func (book Book) Display() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages:%d\nAvail;able Copies: %d\n", book.Title, book.Author, book.Pages, book.CopiesAvailable)
}
func (book *Book) Borrow() {
	if book.CopiesAvailable > 0 {
		book.CopiesAvailable--
		fmt.Println("Borrowed the book....")
	} else {
		fmt.Println("Not availble")
	}
}
func (book *Book) ReturnBook() {
	book.CopiesAvailable += 1
}

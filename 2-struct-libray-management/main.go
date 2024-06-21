package main

import (
	"fmt"

	"github.com/Soyaib10/practice-go-again/tree/main/2-struct-libray-management/books"
	"github.com/Soyaib10/practice-go-again/tree/main/2-struct-libray-management/library"
	"github.com/Soyaib10/practice-go-again/tree/main/2-struct-libray-management/users"
)

func main() {
	lib := &library.Library{}

	// Add books
    lib.AddBook(books.Book{ID: 1, Title: "1984", Author: "George Orwell", Status: "Available"})
    lib.AddBook(books.Book{ID: 2, Title: "To Kill a Mockingbird", Author: "Harper Lee", Status: "Available"})

    // Register users
    lib.RegisterUser(users.User{ID: 1, Name: "Alice", Email: "alice@example.com"})
    lib.RegisterUser(users.User{ID: 2, Name: "Bob", Email: "bob@example.com"})

    // Display books and users
    lib.DisplayBooks()
    lib.DisplayUsers()

	// Lend a book
	err := lib.LendBook(1, 1) 
	if err != nil {
		fmt.Println(err)
	}

	// Display books and users after lending a book
	lib.DisplayBooks()
	lib.DisplayUsers()
}
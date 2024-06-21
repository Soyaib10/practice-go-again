package library

import (
	"fmt"

	"github.com/Soyaib10/practice-go-again/tree/main/2-struct-libray-management/books"
	"github.com/Soyaib10/practice-go-again/tree/main/2-struct-libray-management/users"
)

type Library struct {
	Books []books.Book
	Users []users.User
}

func (lib *Library) AddBook(book books.Book) {
	lib.Books = append(lib.Books, book)
}

func (lib *Library) RegisterUser(user users.User) {
	lib.Users = append(lib.Users, user)
}

func (lib *Library) LendBook(bookID int, userID int) error {
	var bookIndex int
    var userIndex int
    var bookFound bool
    var userFound bool

	for i, book := range lib.Books {
		if book.ID == bookID {
			bookIndex = i
			bookFound = true
			break
		}
	}

	for i, user := range lib.Users {
		if user.ID == userID {
			userIndex = i
			userFound = true
			break
		}
	}

	if !bookFound {
		return fmt.Errorf("book not found")
	}

	if !userFound {
        return fmt.Errorf("user not found")
    }

	if lib.Books[bookIndex].Status == "Lent" {
		return fmt.Errorf("book is already lent")
	}

	lib.Books[bookIndex].Status = "Lent"
	lib.Users[userIndex].Books = append(lib.Users[userIndex].Books, lib.Books[bookIndex])
	return nil
}

func (lib *Library) DisplayBooks() {
	fmt.Println("Books in Library:")
	for _, book := range lib.Books {
		fmt.Printf("ID: %d, Title: %s, Author: %s, Status: %s\n", book.ID, book.Title, book.Author, book.Status)
	}
}

func (lib *Library) DisplayUsers() {
    fmt.Println("Registered Users:")
    for _, user := range lib.Users {
        fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
        fmt.Println("Borrowed Books:")
        for _, book := range user.Books {
            fmt.Printf("    ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
        }
    }
}


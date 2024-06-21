package users

import "github.com/Soyaib10/practice-go-again/tree/main/2-struct-libray-management/books"

type User struct {
	ID int
	Name string
	Email string
	Books []books.Book // a slice of Book structs
}
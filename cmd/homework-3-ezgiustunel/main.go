package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-ezgiustunel/helper"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-ezgiustunel/service/domain/book"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-3-ezgiustunel/service/infrastructure"
)

var bookList []book.Book
var repository *book.BookRepository

const list = "list"
const search = "search"
const buy = "buy"
const delete = "delete"

func init() {
	db := infrastructure.ConnectDB("postgres://ezgiustunel:pass@localhost:5432/library")
	repository = book.NewBookRepository(db)
	repository.Migration()

	bookList, _ = helper.ReadCsv("book.csv")
	repository.InsertData(bookList)
}

func main() {
	args := os.Args

	if len(args) == 1 {
		helper.PrintMessagesToConsole()
		return
	}

	firstInput := args[1]

	switch firstInput {
	case list:
		listBooks()
	case search:
		searchBook(args)
	case buy:
		buyBook(args)
	case delete:
		deleteBook(args)
	default:
		helper.PrintMessagesToConsole()
	}
}

// listBooks: list all books
func listBooks() {
	books := repository.FindAll()

	for _, book := range books {
		fmt.Println(book.Name)
	}
}

// searchBook: searches the books by given input
func searchBook(args []string) {
	if len(args) < 3 {
		helper.PrintMessagesToConsole()
		return
	}

	searchedBook := strings.Join(args[2:], " ")
	books, err := repository.FindByBookName(searchedBook)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, book := range books {
		fmt.Printf("name: %s, author: %s\n", book.Name, book.AuthorName)
	}
}

// buyBook: updates stock number
func buyBook(args []string) {
	if len(args) < 4 {
		helper.PrintMessagesToConsole()
		return
	}

	id, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Please give a number")
		return
	}

	bookNumber, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Println("Please give a number")
		return
	}

	book, err := repository.FindById(id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, errStock := book.DecreaseStockNumber(bookNumber)

	if errStock != nil {
		fmt.Println(errStock.Error())
		return
	}
	repository.Update(book)
}

// deleteBook: delete book from db
func deleteBook(args []string) {
	if len(args) != 3 {
		helper.PrintMessagesToConsole()
		return
	}

	id, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Please give a number")
		return
	}

	repository.DeleteById(id)
}

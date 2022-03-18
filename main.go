package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Books struct has a field called arr.
type Books struct {
	names          []string
	pageNumber     []int
	numberOfStocks []int
	price          []float64
	stockCode      []string
	isbn           []string
	authorInfo     []Authors
}

type Authors struct {
	id         int
	authorName string
}

// b variable constructs a books struct which contains the fields of the books.
var b = Books{
	names:          []string{"The Giver", "Bleak House", "The Lord of The Rings: Fellowship Of The Ring", "White Fang", "Lord of The Flies", "Wiseguy"},
	pageNumber:     []int{},
	numberOfStocks: []int{10, 20, 30, 40, 50, 60},
	price:          []float64{},
	stockCode:      []string{},
	isbn:           []string{"182-124-3423-35765", "223-4575-6788-466445"},
	authorInfo:     []Authors{{id: 1, authorName: "George Orwell"}},
}

//Search function prints the name of the book given as input to the screen.
func (boo *Books) search(bookName []string) string {

	for _, value := range boo.names {

		for i := range bookName {
			if strings.ToLower(strings.Join(bookName[i:], " ")) == strings.ToLower(value) {
				return "Book found: " + value
			}
		}
	}
	return "No book found under the name you entered."
}

//list function prints the book list
func list() {

	fmt.Println("\n-----Booklist-----")

	for i, books := range b.names {
		fmt.Printf("ID : %v, Book: %v\n", i, books)
	}
}

func get(id int) string {

	for i, books := range b.names {
		if i == id && id <= len(b.names)-1 && id >= 0 {
			return "Book: " + books
		}
	}
	return "No books found with this id number."
}

//delete function removes the name of the book from booklist by index number.
func delete(boo *[]string, id int) {

	fmt.Printf("\nThe book %v has been deleted.", (*&b).names[id])
	if id >= 0 && id <= len(b.names)-1 {
		*boo = append(b.names[:id], b.names[id+1:]...)
		fmt.Println()
	}
	b.names = *boo
	list()
}

//buy function implements purchasing the book.
func buy(id int, quantity int) string {
	for i, value := range b.numberOfStocks {
		if i == id && id >= 0 && id <= len(b.numberOfStocks)-1 && value > 0 && quantity <= value && quantity > 0 {
			value -= quantity
			return "You have purchased " + b.names[i]
		}
	}
	return "We ran out of the book " + b.names[id]
}

func main() {

	//args variable returns the command line arguments as a slice from terminal.
	args := os.Args[1:]

	for i := range args {

		if len(args) >= 2 && args[i] == "search" {

			fmt.Println(b.search(os.Args[2:]))
			break

		} else if len(args) == 1 && args[i] == "list" {

			list()
			break

		} else if len(args) >= 2 && args[i] == "get" {
			id, _ := strconv.Atoi(os.Args[2])
			fmt.Println(get(id))
			break

		} else if len(args) >= 2 && args[i] == "delete" {
			id, _ := strconv.Atoi(os.Args[2])
			delete(&b.names, id)
			break

		} else if len(args) >= 2 && args[i] == "buy" {
			id, _ := strconv.Atoi(os.Args[2])
			quantity, _ := strconv.Atoi(os.Args[3])
			fmt.Println(buy(id, quantity))
			break

		} else {
			fmt.Println("Wrong input.")
			break
		}
	}
}

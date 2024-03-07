package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// "errors"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "A", Author: "ABC", Quantity: 5},
	{ID: "2", Title: "B", Author: "EECSD", Quantity: 532},
	{ID: "3", Title: "C", Author: "GFCBC", Quantity: 623},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	// fmt.Println(id,"khk")

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "book id not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*book, error) {

	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("Book not found")
}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Id not passed"})
		return
	}
	fmt.Println("sdfs")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not found"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available"})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)

}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Id not passed"})
		return
	}
	fmt.Println("sdfs")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not found"})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)

}

func multBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Id not passed"})
		return
	}
	fmt.Println("sdfs")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not found"})
		return
	}

	book.Quantity *=2
	c.IndentedJSON(http.StatusOK, book)

}


func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/createbooks", createBook)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)
	router.PATCH("/mult", multBook)
	router.Run("localhost:8080")

}

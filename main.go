package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

var books = []book{
	{ID: "1", Title: "Blood Meridian", Author: "Cormac McCarthy", Price: 30.00},
	{ID: "2", Title: "Hyperion", Author: "Dan Simmons", Price: 12.99},
	{ID: "3", Title: "Gravity's Rainbow", Author: "Thomas Pynchon", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookById)
	router.POST("/books", postAlbums)

	router.Run("localhost:8080")
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func postAlbums(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func getBook(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func getBookById(c *gin.Context) {
	id := c.Param("id")

	for _, b := range books {
		if b.ID == id {
			c.IndentedJSON(http.StatusOK, b)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
}

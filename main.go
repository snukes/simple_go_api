package main

import (
	"net/http"
	"strconv"

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

// simple self incremented counter. Potential for this to run into to synchronization issues
// refactor for higher request volume solutions
var curIdCount = 3

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookById)
	router.POST("/books", postBooks)
	router.OPTIONS("/books", options)

	router.Run("localhost:8080")
}

func getBooks(c *gin.Context) {
	JSON(c, http.StatusOK, books)
}

func postBooks(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	if newBook.ID == "" {
		curIdCount++
		newBook.ID = strconv.Itoa(curIdCount)
	}

	books = append(books, newBook)
	JSON(c, http.StatusCreated, newBook)
}

func getBookById(c *gin.Context) {
	id := c.Param("id")

	for _, b := range books {
		if b.ID == id {
			c.IndentedJSON(http.StatusOK, b)
			return
		}
	}

	JSON(c, http.StatusNotFound, gin.H{"message": "not found"})
}

func options(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, accept, origin, Cache-Control, X-Requested-With")
	c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	c.Status(http.StatusNoContent)
}

func JSON(c *gin.Context, code int, obj interface{}) {
	// local only, don't use this in production
	c.Header("Access-Control-Allow-Origin", "*")
	c.IndentedJSON(code, obj)
}

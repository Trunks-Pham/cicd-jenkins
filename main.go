package main

import (
	"github.com/gin-gonic/gin"
)

type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var Products []Product

func main() {
	r := gin.Default()

	r.GET("/products", func(c *gin.Context) {
		c.JSON(200, Products)
	})

	r.POST("/products", func(c *gin.Context) {
		var newProduct Product
		if err := c.ShouldBindJSON(&newProduct); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		Products = append(Products, newProduct)
		c.JSON(201, gin.H{"message": "Product created successfully"})
	})

	r.Run(":3000")
}

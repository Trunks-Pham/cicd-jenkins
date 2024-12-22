package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetProducts(t *testing.T) {
	// Setup
	r := gin.Default()
	r.GET("/products", func(c *gin.Context) {
		c.JSON(200, Products)
	})

	// Test
	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Verify
	if w.Code != 200 {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	var products []Product
	err = json.Unmarshal(w.Body.Bytes(), &products)
	if err != nil {
		t.Fatal(err)
	}

	if len(products) != len(Products) {
		t.Errorf("Expected %d products, got %d", len(Products), len(products))
	}
}

func TestCreateProduct(t *testing.T) {
	// Setup
	r := gin.Default()
	r.POST("/products", func(c *gin.Context) {
		var newProduct Product
		if err := c.ShouldBindJSON(&newProduct); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Use a new slice for each test
		var testProducts []Product
		testProducts = append(testProducts, newProduct)
		c.JSON(201, gin.H{"products": testProducts})
	})

	// Test
	testProduct := Product{
		ID:    "1",
		Name:  "Product 1",
		Price: 100,
	}

	jsonData, err := json.Marshal(testProduct)
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Verify
	if w.Code != 201 {
		t.Errorf("Expected status code 201, got %d", w.Code)
	}

	var response gin.H
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Error(err)
	}

	if response["message"] != "Product created successfully" {
		t.Errorf("Expected message 'Product created successfully', got '%s'", response["message"])
	}
}

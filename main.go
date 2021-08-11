package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type product struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Price       float32            `json:"price,omitempty" bson:"price,omitempty"`
	Rate        float32            `json:"rate,omitempty" bson:"rate,omitempty"`
	Image       string             `json:"image,omitempty" bson:"image,omitempty"`
}

const MONGO_URL = "SOMETHING HERE"

var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
var client, _ = mongo.Connect(ctx, options.Client().ApplyURI(MONGO_URL))

func main() {

	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/products", func(c *gin.Context) {
		initCollection()
		c.IndentedJSON(http.StatusCreated, gin.H{"message": "Initializing DB is complete"})
	})

	router.GET("/products", func(c *gin.Context) {
		products := getProducts()
		c.IndentedJSON(http.StatusOK, products)
	})

	router.GET("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		product := getProductById(id)

		c.IndentedJSON(http.StatusCreated, product)
	})

	router.DELETE("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		product := deleteProductById(id)

		c.IndentedJSON(http.StatusCreated, product)
	})

	router.PUT("/products/:id", func(c *gin.Context) {
		id := c.Param("id")

		var replacementProduct product
		putError := c.BindJSON(&replacementProduct)

		if putError != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "wrong information"})
			return
		}

		product := updateProduct(id, replacementProduct)
		c.IndentedJSON(http.StatusOK, product)
	})

	router.Run("localhost:8000")
}

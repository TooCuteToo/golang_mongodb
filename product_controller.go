package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func initCollection() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// client, _ := mongo.Connect(ctx, options.Client().ApplyURI(MONGO_URL))
	collection := client.Database("QLProducts").Collection("Products")

	for i := 1; i < 10; i++ {
		collection.InsertOne(ctx, product{
			Name:        faker.Word(),
			Description: faker.Paragraph(),
			Price:       10 + rand.Float32()*(100-10),
			Rate:        0 + rand.Float32()*(5-0),
			Image:       fmt.Sprintf("http://lorempixel.com/200/200?%s", faker.UUIDDigit()),
		})
	}
}

func getProducts() []product {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// client, _ := mongo.Connect(ctx, options.Client().ApplyURI(MONGO_URL))
	collection := client.Database("QLProducts").Collection("Products")

	var products []product

	cursor, _ := collection.Find(ctx, bson.D{})
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product product
		cursor.Decode(&product)
		products = append(products, product)
	}

	return products
}

func getProductById(id string) product {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// client, _ := mongo.Connect(ctx, options.Client().ApplyURI(MONGO_URL))
	collection := client.Database("QLProducts").Collection("Products")

	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objectId}}
	cursor := collection.FindOne(ctx, filter)

	var product product
	cursor.Decode(&product)

	return product
}

func deleteProductById(id string) product {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// client, _ := mongo.Connect(ctx, options.Client().ApplyURI(MONGO_URL))
	collection := client.Database("QLProducts").Collection("Products")

	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objectId}}
	cursor := collection.FindOneAndDelete(ctx, filter)

	var product product
	cursor.Decode(&product)

	return product
}

func updateProduct(id string, newProduct product) product {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// client, _ := mongo.Connect(ctx, options.Client().ApplyURI(MONGO_URL))
	collection := client.Database("QLProducts").Collection("Products")

	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", objectId}}
	cursor := collection.FindOneAndReplace(ctx, filter, newProduct)

	var product product
	cursor.Decode(&cursor)

	return product
}

package controllers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nimi-io/Golang-Ecommerce-Essential/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	prodCollection *mongo.Collection
	UserCollection *mongo.Collection
}

func NewApplication(prodCollection *mongo.Collection, UserCollection *mongo.Collection) *Application {
	return &Application{prodCollection: prodCollection, UserCollection: UserCollection}
}
func (app *Application) AddToCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryId := c.Query("id")
		if productQueryId == "" {
			log.Println("No product id provided")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("No product id provided"))
			return
		}

		userQueryId := c.Query("id")
		if userQueryId == "" {
			log.Println("No user id provided")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("No user id provided"))
			return
		}

		productId, err := primitive.ObjectIDFromHex(productQueryId)
		if err != nil {
			log.Println("Invalid product id")
			_ = c.AbortWithError(http.StatusInternalServerError, errors.New("Invalid product id"))
			return
		}

		userId, err := primitive.ObjectIDFromHex(userQueryId)
		if err != nil {
			log.Println("Invalid user id")
			_ = c.AbortWithError(http.StatusInternalServerError, errors.New("Invalid user id"))
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err = database.AddProductToCart(ctx, app.UserCollection, productId, userId)
		if err != nil {
			log.Println("Error adding product to cart")
			_ = c.AbortWithError(http.StatusInternalServerError, errors.New("Error adding product to cart"))
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Product added to cart"})
		return

	}

}

func (app *Application) RemoveItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryId := c.Query("id")
		if productQueryId == "" {
			log.Println("No product id provided")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("No product id provided"))
			return
		}

		userQueryId := c.Query("id")
		if userQueryId == "" {
			log.Println("No user id provided")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("No user id provided"))
			return
		}

		productId, err := primitive.ObjectIDFromHex(productQueryId)
		if err != nil {
			log.Println("Invalid product id")
			_ = c.AbortWithError(http.StatusInternalServerError, errors.New("Invalid product id"))
			return
		}

		userId, err := primitive.ObjectIDFromHex(userQueryId)
		if err != nil {
			log.Println("Invalid user id")
			_ = c.AbortWithError(http.StatusInternalServerError, errors.New("Invalid user id"))
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err = database.RemoveItemFromCart(ctx, app.UserCollection, productId, userId)
		if err != nil {
			log.Println("Error removing product from cart")
			_ = c.AbortWithError(http.StatusInternalServerError, errors.New("Error removing product from cart"))
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Product removed from cart"})
		return
	}
}

// func (app *Application) GetItemFromCart() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		userQueryId := c.Query("id")
// 		if userQueryId == "" {
// 			log.Println("No user id provided")
// 			_ = c.AbortWithError(http.StatusBadRequest, errors.New("No user id provided"))
// 			return
// 		}

// 		userId, err := primitive.ObjectIDFromHex(userQueryId)
// 		if err != nil {
// 			log.Println("Invalid user id")
// 			_ = c.AbortWithError(http.StatusInternalServerError, errors.New("Invalid user id"))
// 			return
// 		}

// 		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
// 		defer cancel()

// 		cart, err := database.GetCart(ctx, app.UserCollection, userId)
// 		if err != nil {
// 			log.Println("Error getting cart")
// 			_ = c.AbortWithError(http.StatusInternalServerError, errors.New("Error getting cart"))
// 			return
// 		}

// 		c.JSON(http.StatusOK, cart)
// 		return
// 	}
// }

func(app *Application) BuyFromCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		userQueryId := c.Query("id")
		if userQueryId == "" {
			log.Println("No user id provided")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("No user id provided"))
			return
		}

		userId, err := primitive.ObjectIDFromHex(userQueryId)
		if err != nil {
			log.Println("Invalid user id")
			_ = c.AbortWithError(http.StatusInternalServerError, errors.New("Invalid user id"))
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		

		err = database.BuyItemFromCart(ctx, app.UserCollection, userId)
		if err != nil {
			log.Println("Error buying from cart")
			_ = c.AbortWithError(http.StatusInternalServerError, errors.New("Error buying from cart"))
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Product bought from cart"})
		return
	}
}

func (app *Application) InstantBuy() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryId := c.Query("id")
		if productQueryId == "" {
			log.Println("No product id provided")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("No product id provided"))
			return
		}

		productId, err := primitive.ObjectIDFromHex(productQueryId)
		if err != nil {
			log.Println("Invalid product id")
			_ = c.AbortWithError(http.StatusInternalServerError, errors.New("Invalid product id"))
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err = database.InstantBuyer(ctx, app.prodCollection, productId)
		if err != nil {
			log.Println("Error buying product")
			_ = c.AbortWithError(http.StatusInternalServerError, errors.New("Error buying product"))
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Product bought"})
		return
	}
}

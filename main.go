package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nimi-io/Golang-Ecommerce-Essential/controllers"
	"github.com/nimi-io/Golang-Ecommerce-Essential/database"
	"github.com/nimi-io/Golang-Ecommerce-Essential/middleware"
	"github.com/nimi-io/Golang-Ecommerce-Essential/routes"
)

func main(){
	port := os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}

	app:= controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))
	router := gin.New()

	router.Use(gin.Logger())
	routes.UserRoutes(router)

	router.Use(middleware.Authentication())
	routes.CartRoutes(router)
	router.GET("/addtocart", app.AddToCart())
	router.GET("/cart", app.GetCartProducts())
	router.GET("/removeitem", app.RemoveItem())
	router.POST("/checkout", app.Checkout())

	log.Fatal(router.Run(":" + port))
}
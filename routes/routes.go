package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nimi-io/Golang-Ecommerce-Essential/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/admin/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/productview", controllers.SearchProduct())

}

func CartRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/addtocart", controllers.AddToCart())
	incomingRoutes.GET("/cart", controllers.GetCartProducts())
	incomingRoutes.GET("/removeitem", controllers.RemoveItem())
	incomingRoutes.POST("/checkout", controllers.Checkout())
}
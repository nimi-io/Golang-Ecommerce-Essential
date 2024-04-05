package controllers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/nimi-io/Golang-Ecommerce-Essential/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HashPassword(password string) string {

}

func VerifyPassword(password string, hash string) bool {

}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		validate := validator.New()

		defer cancel()

		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(400, gin.H{"error": validationErr.Error()})
			return
		}

		count, err := UserCollection.CountDocuments(ctx, bson.H{"email": user.Email})

		defer cancel()

		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if count > 0 {
			c.JSON(400, gin.H{"error": "Email already exists"})
			return
		}

		hashedPassword := HashPassword(*user.Password)
		user.Password = &hashedPassword

		user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_Id = user.ID.Hex()

		token, refreshToken, _ := generate.TokenGenerator(*user)
		user.Token = token
		user.Refresh_Token = refreshToken
		user.UserCart = make([]models.ProductUser, 0)
		user.Address_Details = make([]models.Address, 0)
		user.Order_Status = make([]models.Order, 0)

		_, err = UserCollection.InsertOne(ctx, user)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		defer cancel()

		c.JSON(200, gin.H{"message": "User created successfully"})

	}
}

func SignIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		validate := validator.New()

		defer cancel()

		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(400, gin.H{"error": validationErr.Error()})
			return
		}

		count, err := UserCollection.CountDocuments(ctx, bson.H{"email": user.Email})

		defer cancel()

		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if count == 0 {
			c.JSON(400, gin.H{"error": "Email does not exist"})
			return
		}

		var dbUser models.User
		err = UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&dbUser)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if !VerifyPassword(*user.Password, *dbUser.Password) {
			c.JSON(400, gin.H{"error": "Invalid password"})
			return
		}

		token, refreshToken, _ := generate.TokenGenerator(*dbUser)
		defer cancel()

		dbUser.Token = token
		dbUser.Refresh_Token = refreshToken

		generate.UpdateAllTokens(token, refreshToken)

		c.JSON(200, gin.H{"message": "User signed in successfully", "user": dbUser, "token": token})
	}

}
func ProductViewerAdmin() gin.HandlerFunc {}

func ProductSearch() gin.HandlerFunc {}

func SearchByQuery() gin.HandlerFunc {}

package main

import (
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	First_Name      *string            `bson:"first_name,omitempty"`
	Last_Name       *string            `bson:"last_name,omitempty"`
	Email           *string            `bson:"email,omitempty"`
	Password        *string            `bson:"password,omitempty"`
	Phone           *string            `bson:"phone,omitempty"`
	Token           *string            `bson:"token,omitempty"`
	Refresh_Token   *string            `bson:"refresh_token,omitempty"`
	CreatedAt       time.Time          `bson:"created_at,omitempty"`
	UpdatedAt       time.Time          `bson:"updated_at,omitempty"`
	User_Id         string             `bson:"user_id,omitempty"`
	Address_Details *string            `bson:"address_details,omitempty"`
	Order_Status    *string            `bson:"order_status,omitempty"`
	UserCart        []ProductUser      `bson:"user_cart,omitempty"`
}

type ProductUser struct {
	// Define your ProductUser struct here
}

func main() {
	// Create a sample user
	user := User{
		ID:              primitive.NewObjectID(),
		First_Name:      StringPtr("John"),
		Last_Name:       StringPtr("Doe"),
		Email:           StringPtr("john.doe@example.com"),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		User_Id:         "12345",
		Address_Details: StringPtr("123 Main St, City, Country"),
		Order_Status:    StringPtr("Pending"),
		UserCart:        []ProductUser{},
	}

	// Marshal the user struct to JSON

	jsonBytes, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println("JSON representation:")
	fmt.Println(string(jsonBytes))

	// Marshal the user struct to BSON
	bsonBytes, err := bson.Marshal(user)
	if err != nil {
		fmt.Println("Error marshaling BSON:", err)
		return
	}
	fmt.Println("BSON representation:")
	fmt.Println(bsonBytes)
}

// Utility function to create a pointer to a string
func StringPtr(s string) *string {
	return &s
}

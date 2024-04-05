package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{
	ID primitive.ObjectID 				`json: "id" bson:"_id"`
	First_Name *string 					`bson:"first_name"					validate: "required,min=3,max=20" json:"first_name"` 
	Last_Name *string 					`bson:"last_name" 					validate: "required,min=3,max=20" json:"last_name"`
	Email *string 						`json:"email" 							validate: "required,email"`
	Password *string 					`json:"password" 						validate: "required,min=6,max=20"`
	Phone *string 						`json:"phone" 							validate: "required,min=10,max=10"`
	Token *string 						`json:"token"`
	Refresh_Token *string 				`json:"refresh_token"`
	CreatedAt time.Time					`json:"created_at"`
	UpdatedAt time.Time					`json:"updated_at"`
	User_Id string 						`json:"user_id"`
	Order_Status []Order				`json:"orders" bson:"order_status"`
	UserCart []ProductUser 				`json:"usercart" bson:"usercart"`
	Address_Details []Address 			`json:"address_details" bson:"address_details"`

}

type Product struct{
	ID primitive.ObjectID `bson:"_id"`
	Image *string `json:"image"`
	Rating *float64 `json:"rating"`
	Product_Name *string `json:"name"`
	Price *float64 `json:"price"`
}

type ProductUser struct{
	Product_Id primitive.ObjectID `bson:"_id"`
	Product_Name *string `json:"name" bson:product_name"`
	Price int `json:"price" bson:"price"`
	Rating *uint `json:"rating" bson:"rating"`
	Image *string `json:"image" bson:"image"`
}
type Address struct{
	Address_Id primitive.ObjectID `bson:"_id"`
	House *string `json: "house_name" bson:"house_name"`
	Street *string `json:"street_name" bson:"street_name"`
	City *string `json:"city_name" bson:"city_name"`	
	Pincode *string `json: "pin_code" bson:"pin_code"`
}
type Order struct{
	Order_Id primitive.ObjectID `bson:"_id"`
	Order_Cart []ProductUser `json:"order_list" bson:"order_list"`
	Ordered_At time.Time `json:"ordered_at" bson:"ordered_at"`
	price int `json:"total_price" bson:"total_price"`
	Discount *int `json:"discount" bson:"discount"`
	Payment_Method *string `json:"payment_method" bson:"payment_method"`
}
type Payment struct{
	Digital bool `bson:"digital"`
	COD bool `bson:"cod"`
}
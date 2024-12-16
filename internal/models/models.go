package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	First_Name *string            `json:"first_name" validate:"required,min=2,max=30"`
	LastName   *string            `json:"last_name" validate:"required,min=2,max=30"`
	Password   *string            `json:"password" validate:"required,min=6"`
	Email      *string            `json:"email" validate:"email,required"`
	Phone      *string            `json:"phone" validate:"required"`
	// Role            string             `json:"role" bson:"role" validate:"require, oneof=admin seller customer"`
	Token           *string      `json:"token"`
	Refresh_Token   *string      `json:"refresh_token"`
	Create_At       time.Time    `json:"create_at"`
	Update_At       time.Time    `json:"update_at"`
	User_ID         string       `json:"user_id"`
	UserCart        []ProdutUser `json:"usercart" bson:"usercart"`
	Address_Details []Address    `json:"address" bson:"address"`
	Order_Status    []Order      `json:"orders" bson:"orders"`
}

type Product struct {
	Product_ID   primitive.ObjectID `bson:"_id"`
	Product_Name *string            `json:"product_name"`
	Price        *uint64            `json:"price"`
	Rating       *float64           `json:"rating"`
	Image        *string            `json:"image"`
}

type ProdutUser struct {
	Product_ID   primitive.ObjectID `bson:"_id"`
	Product_Name *string            `json:"product_name" bson:"product_name"`
	Price        int                `json:"price" bson:"price"`
	Rating       *float64           `json:"rating" bson:"rating"`
	Image        *string            `json:"image" bson:"image"`
}

type Address struct {
	Address_id primitive.ObjectID `bson:"_id"`
	House      *string            `json:"house_name" bson:"house_name"`
	Street     *string            `json:"street_name" bson:"street_name"`
	City       *string            `json:"city_name" bson:"city_name"`
	Pincode    *string            `json:"pin_code" bson:"pin_code"`
}

type Order struct {
	Order_ID       primitive.ObjectID `bson:"_id"`
	Order_Cart     []ProdutUser       `json:"order_list" bson:"order_list"`
	Ordered_At     time.Time          `json:"ordered_at" bson:"ordered_at"`
	Price          int                `json:"total_price" bson:"total_price"`
	Discount       *int               `json:"discount" bson:"discount"`
	Payment_method Payment            `json:"payment_method" bson:"payment_method"`
}

type Payment struct {
	Digital bool `json:"digital" bson:"digital"`
	COD     bool `json:"cod" bson:"cod"`
}

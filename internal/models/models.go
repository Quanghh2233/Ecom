package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	First_Name      *string            `json:"first_name" validate:"required,min=2,max=30"`
	LastName        *string            `json:"last_name" validate:"required,min=2,max=30"`
	Password        *string            `json:"password" validate:"required,min=6"`
	Email           *string            `json:"email" validate:"email,required"`
	Phone           *string            `json:"phone" validate:"required"`
	Role            *Role              `json:"role" bson:"role"`
	Token           *string            `json:"token"`
	Refresh_Token   *string            `json:"refresh_token"`
	Create_At       time.Time          `json:"create_at"`
	Update_At       time.Time          `json:"update_at"`
	User_ID         string             `json:"user_id"`
	UserCart        []ProdutUser       `json:"usercart" bson:"usercart"`
	Address_Details []Address          `json:"address" bson:"address"`
	Order_Status    []Order            `json:"orders" bson:"orders"`
}

type Product struct {
	Product_ID   primitive.ObjectID `json:"product_id" bson:"product_id"`
	Store_ID     primitive.ObjectID `json:"store_id" bson:"store_id"`
	Store_Name   string             `json:"store_name" bson:"store_name"`
	Product_Name string             `json:"product_name"`
	Description  string             `json:"description" bson:"description"`
	Quantity     int                `json:"quantity" bson:"quantity"`
	Price        *float64           `json:"price"`
	Rating       *float64           `json:"rating"`
	Options      []ProductOption    `json:"options" bson:"options"`
	Image        *string            `json:"image"`
}

type ProdutUser struct {
	Product_ID   primitive.ObjectID `bson:"_id"`
	Product_Name *string            `json:"product_name" bson:"product_name"`
	Quantity     int                `json:"quantity" bson:"quantity"`
	Price        *float64           `json:"price" bson:"price"`
	Rating       *float64           `json:"rating" bson:"rating"`
	Image        *string            `json:"image" bson:"image"`
}

type Address struct {
	Address_id primitive.ObjectID `bson:"_id"`
	House      *string            `json:"house_name" bson:"house_name"`
	Street     *string            `json:"street_name" bson:"street_name"`
	City       *string            `json:"city_name" bson:"city_name"`
	Pincode    *string            `json:"pin_code" bson:"pin_code"`
	Type       string             `bson:"type" json:"type"`
}

type Order struct {
	Order_ID       primitive.ObjectID `bson:"_id"`
	Order_Cart     []ProdutUser       `json:"order_list" bson:"order_list"`
	Ordered_At     time.Time          `json:"ordered_at" bson:"ordered_at"`
	Price          *float64           `json:"total_price" bson:"total_price"`
	Discount       *int               `json:"discount" bson:"discount"`
	Payment_method Payment            `json:"payment_method" bson:"payment_method"`
}

type Payment struct {
	Digital bool `json:"digital" bson:"digital"`
	COD     bool `json:"cod" bson:"cod"`
}

type Store struct {
	Store_Id    primitive.ObjectID `json:"store_id" bson:"store_id"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Owner       string             `json:"owner" bson:"owner"`
	Email       string             `json:"email" bson:"email"`
	Phone       string             `json:"phone" bson:"phone"`
	Status      string             `json:"status" bson:"status"`
	CreateAt    time.Time          `json:"create_at" bson:"create_at"`
}

type ProductOption struct {
	Name  string  `json:"name" bson:"name"`
	Value string  `json:"value" bson:"value"`
	Price float64 `json:"price" bson:"price"`
}

// Role struct để lưu trong database
type Role struct {
	Role_ID     primitive.ObjectID `json:"role_id" bson:"_id"`
	Name        string             `json:"role" bson:"name" validate:"required,oneof=ADMIN SELLER BUYER"`
	Description string             `json:"description" bson:"description"`
	Permissions []string           `json:"permissions" bson:"permissions"`
	CreateAt    time.Time          `json:"create_at" bson:"create_at"`
	UpdateAt    time.Time          `json:"update_at" bson:"update_at"`
}

func (r *Role) HasPermission(permission string) bool {
	if r == nil || r.Permissions == nil {
		return false
	}

	for _, p := range r.Permissions {
		if p == permission {
			return true
		}
	}
	return false
}

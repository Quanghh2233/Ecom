package cart

import (
	"context"
	"fmt"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddProductToCart(ctx context.Context, prodCollection, userCollection *mongo.Collection, productID primitive.ObjectID, userID string) error {
	// Check if product exists and is available
	var product models.Product
	err := prodCollection.FindOne(ctx, bson.M{"product_id": productID}).Decode(&product)
	if err != nil {
		return ErrCantFindProduct
	}

	// Check product availability
	if product.Quantity < 1 {
		return fmt.Errorf("product out of stock")
	}

	// Check if user exists
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return ErrUserIdIsNotValid
	}

	// Check if product already in cart
	var user models.User
	err = userCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return ErrCantFindUser
	}

	for _, item := range user.UserCart {
		if item.Product_ID == productID {
			return fmt.Errorf("product already in cart")
		}
	}

	// Add to cart
	update := bson.D{{Key: "$push", Value: bson.D{
		{Key: "usercart", Value: models.ProdutUser{
			Product_ID:   product.Product_ID,
			Product_Name: &product.Product_Name,
			Price:        product.Price,
			Quantity:     1,
		}},
	}}}

	_, err = userCollection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return ErrCantUpdateUser
	}

	return nil
}

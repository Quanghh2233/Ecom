package service

import (
	"context"
	"errors"
	"time"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UpdateProductData struct {
	ProductName *string  `json:"product_name"`
	Price       *uint64  `json:"price"`
	Rating      *float32 `json:"rating"`
	Image       *string  `json:"image"`
}

func (s *productService) UpdateProduct(ctx context.Context, userRole, userID, storeID, productID string, updateData *UpdateProductData) error {
	if userRole != models.ROLE_ADMIN && userRole != models.ROLE_SELLER {
		return errors.New("permission denied")
	}

	objStoreID, err := primitive.ObjectIDFromHex(storeID)
	if err != nil {
		return errors.New("invalid store ID format")
	}

	if userRole == models.ROLE_SELLER {
		objUserID, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			return errors.New("invalid user ID format")
		}

		_, err = s.repo.FindStoreByIDAndOwner(ctx, objStoreID, objUserID)
		if err != nil {
			return errors.New("you don't have permission to update products in this store")
		}
	}

	objProductID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return errors.New("invalid product ID format")
	}

	_, err = s.repo.FindProductByID(ctx, objProductID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("product not found")
		}
		return err
	}

	update := bson.M{"$set": bson.M{}}

	if updateData.ProductName != nil {
		update["$set"].(bson.M)["product_name"] = updateData.ProductName
	}
	if updateData.Price != nil {
		update["$set"].(bson.M)["price"] = updateData.Price
	}
	if updateData.Rating != nil {
		update["$set"].(bson.M)["rating"] = updateData.Rating
	}
	if updateData.Image != nil {
		update["$set"].(bson.M)["image"] = updateData.Image
	}

	update["$set"].(bson.M)["updated_at"] = time.Now()

	matchedCount, modifiedCount, err := s.repo.UpdateProduct(ctx, objProductID, update)
	if err != nil {
		return err
	}

	if matchedCount == 0 {
		return errors.New("product not found")
	}

	if modifiedCount == 0 {
		return errors.New("no changes made to the product")
	}

	return nil
}

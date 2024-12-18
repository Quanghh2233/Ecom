package cart

import (
	"context"
	"errors"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserOrders(ctx context.Context, userCollection *mongo.Collection, userID string) ([]models.Order, error) {
	uid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}
	var user models.User
	filter := bson.M{"_id": uid}
	err = userCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user.Order_Status, nil
}

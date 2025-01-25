package repository

import (
	"context"

	"github.com/Quanghh2233/Ecommerce/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *addressRepository) EditHomeAddress(ctx context.Context, userID primitive.ObjectID, address models.Address) (int64, error) {
	filter := bson.M{"_id": userID, "address.type": "home"}
	update := bson.M{"$set": bson.M{
		"address.$.house_name": address.House,
		"address.$.street":     address.Street,
		"address.$.city_name":  address.City,
		"address.$.pin_code":   address.Pincode,
	}}

	result, err := r.userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil
}

func (r *addressRepository) EditWorkAddress(ctx context.Context, userID primitive.ObjectID, address models.Address) (int64, error) {
	filter := bson.M{"_id": userID, "address.type": "work"}
	update := bson.M{"$set": bson.M{
		"address.$.house_name": address.House,
		"address.$.street":     address.Street,
		"address.$.city_name":  address.City,
		"address.$.pin_code":   address.Pincode,
	}}

	result, err := r.userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil
}

package token

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateAllToken(signedToken string, signedRefreshToken string, userID string) error {
	if userID == "" {
		return ErrInvalidUserID
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updateObj := bson.D{
		{Key: "token", Value: signedToken},
		{Key: "refresh_token", Value: signedRefreshToken},
		{Key: "updated_at", Value: time.Now()},
	}

	filter := bson.M{"user_id": userID}
	options := options.Update().SetUpsert(true)

	_, err := UserData.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: updateObj}}, options)
	if err != nil {
		return err
	}

	return nil
}

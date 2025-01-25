package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *userRepository) UpdateUserInfo(ctx context.Context, userID string, update bson.M) (int64, error) {
	result, err := r.userCollection.UpdateOne(
		ctx,
		bson.M{"user_id": userID},
		update,
	)
	if err != nil {
		return 0, err
	}
	return result.MatchedCount, nil
}

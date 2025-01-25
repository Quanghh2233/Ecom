package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *cartRepository) RemoveCartItem(ctx context.Context, productID primitive.ObjectID, userID string) error {
	// Thực hiện logic xóa sản phẩm khỏi giỏ hàng
	// Đây là triển khai mẫu, bạn nên thay thế bằng logic thực tế

	// Ví dụ: Kiểm tra xem sản phẩm có tồn tại trong giỏ hàng không
	user := r.userCollection.FindOne(ctx, bson.M{"user_id": userID, "cart": bson.M{"$elemMatch": bson.M{"_id": productID}}})
	if user.Err() != nil {
		return errors.New("sản phẩm không tồn tại trong giỏ hàng")
	}

	// Ví dụ: Xóa sản phẩm khỏi giỏ hàng của người dùng
	_, err := r.userCollection.UpdateOne(
		ctx,
		bson.M{"user_id": userID},
		bson.M{"$pull": bson.M{"cart": bson.M{"_id": productID}}},
	)
	if err != nil {
		return err
	}

	return nil
}

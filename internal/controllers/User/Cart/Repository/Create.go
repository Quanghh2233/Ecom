package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *cartRepository) AddProductToCart(ctx context.Context, productID primitive.ObjectID, userID string) error {
	// Thực hiện logic thêm sản phẩm vào giỏ hàng
	// Đây là triển khai mẫu, bạn nên thay thế bằng logic thực tế

	// Ví dụ: Kiểm tra xem sản phẩm có tồn tại không
	product := r.prodCollection.FindOne(ctx, bson.M{"_id": productID})
	if product.Err() != nil {
		return errors.New("sản phẩm không tồn tại")
	}

	// Ví dụ: Thêm sản phẩm vào giỏ hàng của người dùng
	_, err := r.userCollection.UpdateOne(
		ctx,
		bson.M{"user_id": userID},
		bson.M{"$push": bson.M{"cart": productID}},
	)
	if err != nil {
		return err
	}

	return nil
}

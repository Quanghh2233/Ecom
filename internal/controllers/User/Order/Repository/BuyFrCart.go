package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CartRepository interface {
	BuyItemsFromCart(ctx context.Context, userID string, selectedItems []string) ([]string, error)
}

type cartRepository struct {
	userCollection *mongo.Collection
}

func NewCartRepository(userCollection *mongo.Collection) CartRepository {
	return &cartRepository{userCollection: userCollection}
}

func (r *cartRepository) BuyItemsFromCart(ctx context.Context, userID string, selectedItems []string) ([]string, error) {
	// Kiểm tra xem người dùng có tồn tại không
	user := r.userCollection.FindOne(ctx, bson.M{"user_id": userID})
	if user.Err() != nil {
		return nil, errors.New("người dùng không tồn tại")
	}

	// Thực hiện logic mua hàng từ giỏ hàng
	// Đây là triển khai mẫu, bạn nên thay thế bằng logic thực tế

	// Giả sử orderList là danh sách các sản phẩm đã mua
	orderList := selectedItems

	return orderList, nil
}

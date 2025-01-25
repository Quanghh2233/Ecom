package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository interface {
	GetUserOrders(ctx context.Context, userID string) ([]bson.M, error)
}

type orderRepository struct {
	userCollection *mongo.Collection
}

func NewOrderRepository(userCollection *mongo.Collection) OrderRepository {
	return &orderRepository{userCollection: userCollection}
}

func (r *orderRepository) GetUserOrders(ctx context.Context, userID string) ([]bson.M, error) {
	// Thực hiện logic lấy danh sách đơn hàng của người dùng
	// Đây là triển khai mẫu, bạn nên thay thế bằng logic thực tế

	// Ví dụ: Lấy danh sách đơn hàng từ collection của người dùng
	cursor, err := r.userCollection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []bson.M
	if err = cursor.All(ctx, &orders); err != nil {
		return nil, err
	}

	return orders, nil
}

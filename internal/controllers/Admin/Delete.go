package Admin

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Quanghh2233/Ecommerce/internal/controllers/global"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// Lấy product_id từ URL
		productID := c.Param("product_id")
		if productID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
			return
		}

		// Chuyển đổi productID thành ObjectID
		objID, err := primitive.ObjectIDFromHex(productID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Product ID"})
			return
		}

		// Thực hiện xóa sản phẩm khỏi MongoDB
		result, err := global.ProductCollection.DeleteOne(ctx, bson.M{"product_id": objID})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
			return
		}

		if result.DeletedCount == 0 {
			// Nếu không tìm thấy, thử tìm theo chuỗi số 0
			result, err = global.ProductCollection.DeleteOne(ctx, bson.M{"product_id": "000000000000000000000000"})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
				return
			}

			// Kiểm tra nếu sản phẩm không tồn tại
			if result.DeletedCount == 0 {
				log.Println(err)
				c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
				return
			}
		}

		// Trả về kết quả thành công
		c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
	}
}

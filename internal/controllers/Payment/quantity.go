package payment

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Quanghh2233/Ecommerce/internal/controllers/global"
	"github.com/Quanghh2233/Ecommerce/internal/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuantityChange struct {
	ProductID   primitive.ObjectID `json:"product_id" bson:"product_id"`
	OldQuantity int                `json:"old_quantity" bson:"old_quantity"`
	NewQuantity int                `json:"new_quantity" bson:"new_quantity"`
	ChangedAt   time.Time          `json:"changed_at" bson:"changed_at"`
}

// Hàm ghi log thay đổi số lượng
func UpdateProductQuantityAfterPayment(userCart []models.ProdutUser) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		for _, item := range userCart {
			// Tìm sản phẩm trong database
			var product models.Product
			err := global.ProductCollection.FindOne(ctx, bson.M{"_id": item.Product_ID}).Decode(&product)
			if err != nil {
				log.Printf("Error finding product %s: %v", item.Product_ID, err)
				c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Product not found: %s", item.Product_ID)})
				return
			}

			// Kiểm tra số lượng còn đủ không
			if product.Quantity < item.Quantity {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": fmt.Sprintf(
						"Insufficient quantity for product %s. Available: %d, Requested: %d",
						product.Product_Name,
						product.Quantity,
						item.Quantity,
					),
				})
				return
			}

			// Cập nhật số lượng sản phẩm
			newQuantity := product.Quantity - item.Quantity
			update := bson.M{
				"$set": bson.M{"quantity": newQuantity},
			}

			result, err := global.ProductCollection.UpdateOne(
				ctx,
				bson.M{"_id": item.Product_ID},
				update,
			)

			if err != nil {
				log.Printf("Error updating product quantity: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product quantity"})
				return
			}

			if result.ModifiedCount == 0 {
				log.Printf("No product was updated for ID: %s", item.Product_ID)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product quantity"})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Product quantities updated successfully",
		})
	}
}

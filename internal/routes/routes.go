package route

import (
	controllers "github.com/Quanghh2233/Ecommerce/internal/controllers/Adm"
	Auth "github.com/Quanghh2233/Ecommerce/internal/controllers/Auth"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", Auth.Signup())
	incomingRoutes.POST("/users/login", Auth.Login())

	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewAdmin())
	incomingRoutes.PUT("/admin/updateproduct/:product_id", controllers.UpdateProduct())
	incomingRoutes.GET("/users/productview", controllers.SearchProduct())
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())
	incomingRoutes.DELETE("/admin/deleteproduct/:product_id", controllers.DeleteProduct())
}

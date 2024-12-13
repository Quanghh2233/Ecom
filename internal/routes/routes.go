package route

import (
	controller "github.com/Quanghh2233/Ecommerce/internal/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controller.Signup())
	incomingRoutes.POST("/users/login", controller.Login())
	incomingRoutes.POST("/admin/addproduct", controller.ProductViewAdmin())
	incomingRoutes.GET("/users/productview", controller.SearchProduct())
	incomingRoutes.GET("/users/search", controller.SearchProductByQuery())
}

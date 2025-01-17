package route

import (
	controllers "github.com/Quanghh2233/Ecommerce/internal/controllers/Adm"
	Auth "github.com/Quanghh2233/Ecommerce/internal/controllers/Auth"
	search "github.com/Quanghh2233/Ecommerce/internal/controllers/Search"
	store "github.com/Quanghh2233/Ecommerce/internal/controllers/Store"
	"github.com/Quanghh2233/Ecommerce/internal/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/signup", Auth.Signup())
	incomingRoutes.POST("/login", Auth.Login())

	incomingRoutes.GET("/view", controllers.SearchProduct())
	incomingRoutes.GET("/filter", controllers.FilterProd())
	// incomingRoutes.GET("/search", search.SearchHandler())

	//admin route
	// incomingRoutes.POST("/admin/addproduct", controllers.ProductViewAdmin())
	// incomingRoutes.PUT("/admin/updateproduct/:product_id", controllers.UpdateProduct())
	incomingRoutes.DELETE("/admin/deleteproduct/:product_id", controllers.DeleteProduct())
	incomingRoutes.POST("/admin/delete_multiple", controllers.DelMultiple())

	// search route
	incomingRoutes.GET("/search/product", search.SearchProductByQuery())
	incomingRoutes.GET("/search/store", search.SearchStore())

	//store route
	seller := incomingRoutes.Group("/store")
	seller.Use(middleware.Authentication())
	{
		seller.GET("/:store_id", store.GetStore())
		seller.POST("/:store_id/addproduct", store.CreateProduct())
		seller.DELETE("/delete/:product_id", store.DeleteProduct())
	}
}

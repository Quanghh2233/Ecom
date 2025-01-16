package route

import (
	controllers "github.com/Quanghh2233/Ecommerce/internal/controllers/Adm"
	Auth "github.com/Quanghh2233/Ecommerce/internal/controllers/Auth"
	search "github.com/Quanghh2233/Ecommerce/internal/controllers/Search"
	store "github.com/Quanghh2233/Ecommerce/internal/controllers/Store"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", Auth.Signup())
	incomingRoutes.POST("/users/login", Auth.Login())

	incomingRoutes.GET("/view", controllers.SearchProduct())
	incomingRoutes.GET("/filter", controllers.FilterProd())
	// incomingRoutes.GET("/search", search.SearchHandler())

	//admin route
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewAdmin())
	incomingRoutes.PUT("/admin/updateproduct/:product_id", controllers.UpdateProduct())
	incomingRoutes.DELETE("/admin/deleteproduct/:product_id", controllers.DeleteProduct())
	incomingRoutes.POST("/admin/delete_multiple", controllers.DelMultiple())

	// search route
	incomingRoutes.GET("/search/product", search.SearchProductByQuery())
	incomingRoutes.GET("/search/store", search.SearchStore())

	//store route
	incomingRoutes.GET("/stores/:store_id", store.GetStore())
	incomingRoutes.POST("/stores/:store_id/addproduct", store.CreateProduct())
}

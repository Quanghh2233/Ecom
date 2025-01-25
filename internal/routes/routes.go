package route

// import (
// 	controllers "github.com/Quanghh2233/Ecommerce/internal/controllers/Adm"
// 	Auth "github.com/Quanghh2233/Ecommerce/internal/controllers/Auth"
// 	search "github.com/Quanghh2233/Ecommerce/internal/controllers/Search"
// 	store "github.com/Quanghh2233/Ecommerce/internal/controllers/Store"
// 	"github.com/Quanghh2233/Ecommerce/internal/middleware"

// 	"github.com/gin-gonic/gin"
// )

// func UserRoutes(incomingRoutes *gin.Engine) {
// 	// Authentication routes
// 	incomingRoutes.POST("/signup", Auth.Signup())
// 	incomingRoutes.POST("/login", Auth.Login())

// 	// Public routes
// 	public := incomingRoutes.Group("")
// 	{
// 		public.GET("/view", controllers.SearchProduct())
// 		public.GET("/filter", controllers.FilterProd())
// 		public.GET("/search/product", search.SearchProductByQuery())
// 		public.GET("/search/store", search.SearchStore())
// 		public.GET("/store/:store_id", store.GetStore())

// 	}

// 	// Admin routes
// 	adminRoutes := incomingRoutes.Group("/admin")
// 	adminRoutes.Use(middleware.Authentication(), middleware.CheckPermission("admin"))
// 	{
// 		adminRoutes.DELETE("/deleteproduct/:product_id", controllers.DeleteProduct())
// 		adminRoutes.POST("/delete_multiple", controllers.DelMultiple())
// 		// adminRoutes.POST("/addproduct", controllers.ProductViewAdmin())
// 		// adminRoutes.PUT("/updateproduct/:product_id", controllers.UpdateProduct())
// 	}

// 	// Store routes
// 	stores := incomingRoutes.Group("/store")
// 	stores.Use(middleware.Authentication())
// 	{
// 		stores.POST("/:store_id/addproduct", middleware.CheckPermission("manage_products"), store.CreateProduct())
// 		stores.DELETE("/:store_id/delete/:product_id", middleware.CheckPermission("manage_products"), store.DeleteProduct())
// 		stores.PUT("/:store_id/product/:product_id", middleware.CheckPermission("manage_products"), store.UpdateProduct())
// 	}
// }

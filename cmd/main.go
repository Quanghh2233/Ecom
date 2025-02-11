package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/Quanghh2233/Ecommerce/internal/database"
	"github.com/Quanghh2233/Ecommerce/internal/middleware"
	route "github.com/Quanghh2233/Ecommerce/internal/routes"

	helper "github.com/Quanghh2233/Ecommerce/internal/Helper"
	addr "github.com/Quanghh2233/Ecommerce/internal/controllers/Address"
	cart "github.com/Quanghh2233/Ecommerce/internal/controllers/Cart"
	order "github.com/Quanghh2233/Ecommerce/internal/controllers/Order"
	store "github.com/Quanghh2233/Ecommerce/internal/controllers/Store"
	user "github.com/Quanghh2233/Ecommerce/internal/controllers/User"
)

// init() sẽ được chạy trước main(), đảm bảo biến môi trường được nạp
func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Lỗi khi tải file .env")
	}

}

func main() {
	// Lấy PORT từ biến môi trường (nếu không có, dùng mặc định là 8000)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// Khởi tạo dữ liệu ban đầu (Admin User)
	helper.SeedAdminUser()

	// database.InitRedis()

	// Khởi tạo các ứng dụng
	dbClient := database.Client
	prodCollection := database.ProductData(dbClient, "Products")
	userCollection := database.UserData(dbClient, "Users")
	storeCollection := database.StoreData(dbClient, "Store")
	orderCollection := database.OrderData(dbClient, "Orders")

	cartApp := cart.NewApplication(prodCollection, userCollection, orderCollection)
	orderApp := order.NewApplication(prodCollection, userCollection)
	storeApp := store.NewApplication(prodCollection, userCollection, storeCollection)
	userApp := user.NewApplication(prodCollection, userCollection, storeCollection, database.RedisClient)

	// Khởi tạo router Gin
	router := gin.New()
	router.Use(gin.Logger())

	// Định nghĩa các tuyến đường chung
	route.UserRoutes(router)
	router.Use(middleware.Authentication())

	// Buyer Routes (Yêu cầu Authentication)
	buyer := router.Group("/")
	buyer.Use(middleware.Authentication())

	registerBuyerRoutes(buyer, cartApp, orderApp)
	registerUserRoutes(router, userApp)
	registerStoreRoutes(router, storeApp)

	// Khởi động server
	log.Fatal(router.Run(":" + port))
}

// registerBuyerRoutes thiết lập các routes dành cho Buyer
func registerBuyerRoutes(r *gin.RouterGroup, cartApp *cart.Application, orderApp *order.Application) {
	r.POST("/addaddress", middleware.CheckPermission("manage_own_profile"), addr.AddAddress())
	r.PUT("/edithomeaddress", middleware.CheckPermission("manage_own_profile"), addr.EditHomeAddress())
	r.PUT("/editworkaddress", middleware.CheckPermission("manage_own_profile"), addr.EditWorkAddress())
	r.DELETE("/deleteaddress", middleware.CheckPermission("manage_own_profile"), addr.DeleteAddress())

	r.GET("/addtocart", middleware.CheckPermission("manage_own_cart"), cartApp.AddToCart())
	r.DELETE("/removeitem", middleware.CheckPermission("manage_own_cart"), cartApp.RemoveItem())
	r.GET("/listcart", middleware.CheckPermission("manage_own_cart"), cart.GetItemFromCart()) // listcart?user_id=

	r.GET("/cartcheckout", middleware.CheckPermission("place_orders"), orderApp.BuyFromCart())
	r.GET("/instantbuy", middleware.CheckPermission("place_orders"), orderApp.InstantBuy())
	r.DELETE("/cancelorder", middleware.CheckPermission("place_orders"), orderApp.CancelOrder())
	r.DELETE("/cancelall", middleware.CheckPermission("place_orders"), orderApp.CancelAll())
	r.GET("/order_list", middleware.CheckPermission("place_orders"), cartApp.GetOrders())
	r.GET("/cancel_list", middleware.CheckPermission("place_orders"), cartApp.CancelList())
}

// registerUserRoutes thiết lập các routes dành cho User
func registerUserRoutes(router *gin.Engine, userApp *user.Application) {
	router.GET("/user/info", middleware.Authentication(), userApp.GetUserInfo())
	router.POST("/user/info/change-role", middleware.Authentication(), userApp.ChangeRole())
}

// registerStoreRoutes thiết lập các routes dành cho Store
func registerStoreRoutes(router *gin.Engine, storeApp *store.Application) {
	router.POST("/store/register", middleware.CheckPermission("manage_own_profile"), storeApp.RegisterSeller())
}

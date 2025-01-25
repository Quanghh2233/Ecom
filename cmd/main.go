package main

// import (
// 	"log"
// 	"os"

// 	helper "github.com/Quanghh2233/Ecommerce/internal/Helper"
// 	addr "github.com/Quanghh2233/Ecommerce/internal/controllers/Addr"
// 	cart "github.com/Quanghh2233/Ecommerce/internal/controllers/Cart"
// 	store "github.com/Quanghh2233/Ecommerce/internal/controllers/Store"
// 	user "github.com/Quanghh2233/Ecommerce/internal/controllers/User"
// 	"github.com/joho/godotenv"

// 	order "github.com/Quanghh2233/Ecommerce/internal/controllers/Order"

// 	"github.com/Quanghh2233/Ecommerce/internal/database"
// 	"github.com/Quanghh2233/Ecommerce/internal/middleware"
// 	route "github.com/Quanghh2233/Ecommerce/internal/routes"
// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8000"
// 	}

// 	helper.SeedAdminUser()

// 	app := cart.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))
// 	orderApp := order.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))
// 	storeApp := store.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"), database.StoreData(database.Client, "Store"))
// 	userApp := user.NewApplication(database.UserData(database.Client, "Users"), database.StoreData(database.Client, "Store"))

// 	router := gin.New()
// 	router.Use(gin.Logger())

// 	route.UserRoutes(router)
// 	router.Use(middleware.Authentication())
// 	buyer := router.Group("/")
// 	buyer.Use(middleware.Authentication())

// 	{
// 		buyer.POST("/addaddress", middleware.CheckPermission("manage_own_profile"), addr.AddAddress())
// 		buyer.PUT("/edithomeaddress", middleware.CheckPermission("manage_own_profile"), addr.EditHomeAddress())
// 		buyer.PUT("/editworkaddress", middleware.CheckPermission("manage_own_profile"), addr.EditWorkAddress())
// 		buyer.DELETE("/deleteaddress", middleware.CheckPermission("manage_own_profile"), addr.DeleteAddress())

// 		buyer.GET("/addtocart", middleware.CheckPermission("manage_own_cart"), app.AddToCart())
// 		buyer.DELETE("/removeitem", middleware.CheckPermission("manage_own_cart"), app.RemoveItem())
// 		buyer.GET("/listcart", middleware.CheckPermission("manage_own_cart"), cart.GetItemFromCart()) //listcart?user_id=

// 		buyer.GET("/cartcheckout", middleware.CheckPermission("place_orders"), orderApp.BuyFromCart())
// 		buyer.GET("/instantbuy", middleware.CheckPermission("place_orders"), orderApp.InstantBuy())
// 		buyer.DELETE("/cancelorder", middleware.CheckPermission("place_orders"), orderApp.CancelOrder())
// 		buyer.DELETE("/cancelall", middleware.CheckPermission("place_orders"), orderApp.CancelAll())
// 		buyer.GET("/order_list", middleware.CheckPermission("place_orders"), app.GetOrders())
// 	}
// 	router.GET("/user/info", middleware.Authentication(), userApp.GetUserProfile())
// 	//store route
// 	router.POST("/admin/addstores", storeApp.AdmAddStore())
// 	router.POST("/store/register", middleware.CheckPermission("manage_own_profile"), storeApp.RegisterSeller())
// 	// router.POST("/stores/addproduct", storeApp.CreateProduct())

// 	log.Fatal(router.Run(":" + port))
// }

// func init() {
// 	if err := godotenv.Load(); err != nil {
// 		log.Fatal("Lỗi khi tải file .env")
// 	}
// }

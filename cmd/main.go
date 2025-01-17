package main

import (
	"log"
	"os"

	helper "github.com/Quanghh2233/Ecommerce/internal/Helper"
	addr "github.com/Quanghh2233/Ecommerce/internal/controllers/Addr"
	cart "github.com/Quanghh2233/Ecommerce/internal/controllers/Cart"
	store "github.com/Quanghh2233/Ecommerce/internal/controllers/Store"
	"github.com/joho/godotenv"

	order "github.com/Quanghh2233/Ecommerce/internal/controllers/Order"

	"github.com/Quanghh2233/Ecommerce/internal/database"
	"github.com/Quanghh2233/Ecommerce/internal/middleware"
	route "github.com/Quanghh2233/Ecommerce/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	helper.SeedAdminUser()

	app := cart.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))
	orderApp := order.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))
	storeApp := store.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"), database.StoreData(database.Client, "Store"))

	router := gin.New()
	router.Use(gin.Logger())

	route.UserRoutes(router)
	router.Use(middleware.Authentication())
	buyer := router.Group("/")
	buyer.Use(middleware.AuthRole("Buyer"))

	{
		buyer.POST("/addaddress", addr.AddAddress())
		buyer.PUT("/edithomeaddress", addr.EditHomeAddress())
		buyer.PUT("/editworkaddress", addr.EditWorkAddress())
		buyer.DELETE("/deleteaddress", addr.DeleteAddress())

		buyer.GET("/addtocart", app.AddToCart())
		buyer.DELETE("/removeitem", app.RemoveItem())
		buyer.GET("/listcart", cart.GetItemFromCart())

		buyer.GET("/cartcheckout", orderApp.BuyFromCart())
		buyer.GET("/instantbuy", orderApp.InstantBuy())
		buyer.DELETE("/cancelorder", orderApp.CancelOrder())
		buyer.DELETE("/cancelall", orderApp.CancelAll())
		buyer.GET("/order_list", app.GetOrders())
	}

	//store route
	router.POST("/admin/addstores", storeApp.AdmAddStore())
	router.POST("/store/register", storeApp.RegisterSeller())
	// router.POST("/stores/addproduct", storeApp.CreateProduct())

	log.Fatal(router.Run(":" + port))
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Lỗi khi tải file .env")
	}
}

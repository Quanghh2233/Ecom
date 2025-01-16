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

	// address route
	router.POST("/addaddress", addr.AddAddress())
	router.PUT("/edithomeaddress", addr.EditHomeAddress())
	router.PUT("/editworkaddress", addr.EditWorkAddress())
	router.DELETE("/deleteaddress", addr.DeleteAddress())

	//cart route
	router.GET("/addtocart", app.AddToCart())
	router.DELETE("/removeitem", app.RemoveItem())
	router.GET("/listcart", cart.GetItemFromCart())

	//order route
	router.GET("/cartcheckout", orderApp.BuyFromCart())
	router.GET("/instantbuy", orderApp.InstantBuy())
	router.DELETE("/cancelorder", orderApp.CancelOrder())
	router.DELETE("/cancelall", orderApp.CancelAll())
	router.GET("/order_list", app.GetOrders())

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

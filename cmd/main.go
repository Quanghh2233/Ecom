package main

import (
	"log"
	"os"

	addr "github.com/Quanghh2233/Ecommerce/internal/controllers/Addr"
	cart "github.com/Quanghh2233/Ecommerce/internal/controllers/Cart"

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

	app := cart.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	route.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/listcart", cart.GetItemFromCart())
	router.POST("/addaddress", addr.AddAddress())
	router.PUT("/edithomeaddress", addr.EditHomeAddress())
	router.PUT("/editworkaddress", addr.EditWorkAddress())
	router.GET("/deleteaddress", addr.DeleteAddress())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))

}

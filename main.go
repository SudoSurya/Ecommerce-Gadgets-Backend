package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/20pa5a1210/Ecommerce-Gadgets-Backend/database"
	"github.com/20pa5a1210/Ecommerce-Gadgets-Backend/handlers"
	"github.com/20pa5a1210/go-todo/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	_, err := database.NewDatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}
	// create a blank / routes whether its a GET  request

	router.GET("/", func(c *gin.Context) {
		readmeContent, err := ioutil.ReadFile("./README.md") // Replace "static/README.md" with the correct path to your README.md file
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read README file"})
			return
		}

		htmlContent := blackfriday.Run(readmeContent)

		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(http.StatusOK, string(htmlContent))
	})

	user := router.Group("/user")
	user.Use()
	{
		user.POST("/register", handlers.CreateUser)
		user.POST("/login", handlers.LoginUser)
		authGroup := user.Group("/auth")
		authGroup.Use(middleware.AuthMiddleware)
		{
			authGroup.GET("/profile", handlers.GetProfile)
		}
	}
	products := router.Group("/products")
	products.Use()
	{
		products.GET("", handlers.GetProducts)
		products.GET("/:id", handlers.GetProductById)
		products.GET("/page", handlers.GetProductsByPage)
	}
	cart := router.Group("/cart")
	cart.Use(middleware.AuthMiddleware)
	{
		cart.GET("/:username", handlers.GetCartItems)
		cart.POST("/add/:username", handlers.AddProductToCart)
		cart.DELETE("/remove/:username/:productID", handlers.DeleteProductFromCart)
		cart.DELETE("/clear/:username", handlers.ClearItemsFromCart)
		cart.PUT("/increment/:username/:productID", handlers.IncrementProductQuantity)
		cart.PUT("/decrement/:username/:productID", handlers.DecrementProductQuantity)
	}
	router.Run(":8080")
}

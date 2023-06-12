package main

import (
	"log"
	"time"

	"github.com/20pa5a1210/Ecommerce-Gadgets-Backend/database"
	"github.com/20pa5a1210/Ecommerce-Gadgets-Backend/handlers"
	"github.com/20pa5a1210/go-todo/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
	router.Run(":8080")
}

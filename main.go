package main

import (
	"log"
	"net/http"

	"github.com/20pa5a1210/Ecommerce-Gadgets-Backend/database"
	"github.com/20pa5a1210/go-todo/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	_, err := database.NewDatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}
	router.GET("/", func(c *gin.Context) {
		utils.RespondWithError(c, http.StatusOK, "Working")
	})
	router.Run(":8080")
}

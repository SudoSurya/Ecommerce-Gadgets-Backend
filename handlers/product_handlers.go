package handlers

import (
	"log"
	"net/http"

	"github.com/20pa5a1210/Ecommerce-Gadgets-Backend/collections"
	"github.com/20pa5a1210/Ecommerce-Gadgets-Backend/database"
	"github.com/20pa5a1210/go-todo/utils"
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []collections.Product
	database, err := database.NewDatabaseConnection()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to connect to database")
		return
	}
	productCollection := collections.ProductCollectionInit(database.Database)
	products, err = productCollection.GetAllProducts()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to get products")
		log.Println(err)
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, "products", products)
}

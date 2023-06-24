package handlers

import (
	"fmt"
	"net/http"

	"github.com/20pa5a1210/Ecommerce-Gadgets-Backend/collections"
	"github.com/20pa5a1210/Ecommerce-Gadgets-Backend/database"
	"github.com/20pa5a1210/go-todo/utils"
	"github.com/gin-gonic/gin"
)

func connectToDatabase() (*database.DatabaseConnection, error) {
	db, err := database.NewDatabaseConnection()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db, nil
}

func GetCartItems(c *gin.Context) {
	db, err := connectToDatabase()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to connect to database")
	}
	CartCollection := collections.CartCollectionInit(db.Database)
	cartItems, err := CartCollection.GetCartItems()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to get cart items")
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, "cartItems", cartItems)
}

package handlers

import (
	"fmt"
	"net/http"

	"github.com/20pa5a1210/Ecommerce-Gadgets-Backend/collections"
	"github.com/20pa5a1210/Ecommerce-Gadgets-Backend/database"
	"github.com/20pa5a1210/Ecommerce-Gadgets-Backend/models"
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
	username := c.Param("username")
	db, err := connectToDatabase()
	if err != nil {
		utils.RespondWithError(c, http.StatusBadGateway, "Failed to connect to database")
	}
	CartCollection := collections.CartCollectionInit(db.Database)
	cartItems, err := CartCollection.GetCartItems(username)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Failed to get cart items")
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, "cartItems", cartItems)
}

func AddProductToCart(c *gin.Context) {
	username := c.Param("username")
	var product models.Cart
	err := c.BindJSON(&product)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	db, err := connectToDatabase()
	if err != nil {
		utils.RespondWithError(c, http.StatusBadGateway, "Failed to connect to database")
		return
	}
	CartCollection := collections.CartCollectionInit(db.Database)

	existingProduct, err := CartCollection.GetProductByID(username, product.ID)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, err.Error())
		return
	}
	if existingProduct.ID != "" {
		utils.RespondWithError(c, http.StatusConflict, "Product already exists")
		return
	}

	res, err := CartCollection.AddProductToCart(username, product)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, err.Error())
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, "cartItem", res)

}

func DeleteProductFromCart(c *gin.Context) {
	username := c.Param("username")
	productID := c.Param("productID")
	db, err := connectToDatabase()
	if err != nil {
		utils.RespondWithError(c, http.StatusBadGateway, "Failed to connect to database")
		return
	}
	CartCollection := collections.CartCollectionInit(db.Database)
	err = CartCollection.DeleteProductFromCart(username, productID)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, err.Error())
		return
	}

	remainingCartItems, err := CartCollection.GetCartItems(username)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Failed to get cart items")
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, "cartItems", remainingCartItems)
}

func ClearItemsFromCart(c *gin.Context) {
	username := c.Param("username")
	db, err := connectToDatabase()
	if err != nil {
		utils.RespondWithError(c, http.StatusBadGateway, "Failed to connect to database")
		return
	}
	CartCollection := collections.CartCollectionInit(db.Database)
	err = CartCollection.ClearCart(username)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, "cartItems", []models.Cart{})
}

func IncrementProductQuantity(c *gin.Context) {
	username := c.Param("username")
	productID := c.Param("productID")
	db, err := connectToDatabase()
	if err != nil {
		utils.RespondWithError(c, http.StatusBadGateway, "Failed to connect to database")
		return
	}
	CartCollection := collections.CartCollectionInit(db.Database)
	err = CartCollection.IncrementProductQuantity(username, productID)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	cartItems, err := CartCollection.GetCartItems(username)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Failed to get cart items")
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, "cartItems", cartItems)
}

func DecrementProductQuantity(c *gin.Context) {
	username := c.Param("username")
	productID := c.Param("productID")
	db, err := connectToDatabase()
	if err != nil {
		utils.RespondWithError(c, http.StatusBadGateway, "Failed to connect to database")
		return
	}
	CartCollection := collections.CartCollectionInit(db.Database)
	err = CartCollection.DecrementProductQuantity(username, productID)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadGateway, err.Error())
		return
	}
	cartItems, err := CartCollection.GetCartItems(username)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Failed to get cart items")
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, "cartItems", cartItems)
}

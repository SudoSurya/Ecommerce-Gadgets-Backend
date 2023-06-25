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
		return
	}
	cartCollection := collections.CartCollectionInit(db.Database)
	cartItems, err := cartCollection.GetCartItems(username)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to get cart items")
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
	cartCollection := collections.CartCollectionInit(db.Database)

	existingProduct, err := cartCollection.GetProductByID(username, product.ID)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	if existingProduct.ID != "" {
		utils.RespondWithError(c, http.StatusConflict, "Product already exists")
		return
	}

	res, err := cartCollection.AddProductToCart(username, product)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
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
	cartCollection := collections.CartCollectionInit(db.Database)
	err = cartCollection.DeleteProductFromCart(username, productID)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	remainingCartItems, err := cartCollection.GetCartItems(username)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to get cart items")
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
	cartCollection := collections.CartCollectionInit(db.Database)
	err = cartCollection.ClearCart(username)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
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
	cartCollection := collections.CartCollectionInit(db.Database)
	err = cartCollection.IncrementProductQuantity(username, productID)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	cartItems, err := cartCollection.GetCartItems(username)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to get cart items")
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
	cartCollection := collections.CartCollectionInit(db.Database)
	err = cartCollection.DecrementProductQuantity(username, productID)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	cartItems, err := cartCollection.GetCartItems(username)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to get cart items")
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, "cartItems", cartItems)
}

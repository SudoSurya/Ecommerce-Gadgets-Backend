package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/20pa5a1210/Ecommerce-Gadgets-Backend/collections"
	"github.com/20pa5a1210/Ecommerce-Gadgets-Backend/database"
	"github.com/20pa5a1210/Ecommerce-Gadgets-Backend/models"
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

func GetProductById(c *gin.Context) {
	var product models.Product
	database, err := database.NewDatabaseConnection()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to connect to database")
		return
	}
	productCollection := collections.ProductCollectionInit(database.Database)
	product, err = productCollection.GetProductById(c.Param("id"))
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		log.Println(err)
		return
	}
	utils.RespondWithJSON(c, http.StatusOK, "product", product)
}

func GetProductsByPage(c *gin.Context) {
	page := c.DefaultQuery("page", "1")          // Default to page 1 if not provided
	pageSize := c.DefaultQuery("pageSize", "10") // Default to a page size of 10 if not provided

	// Convert page and pageSize to appropriate data types
	pageNumber, _ := strconv.Atoi(page)
	pageSizeNumber, _ := strconv.Atoi(pageSize)

	// Retrieve products from the database based on the page and page size
	database, err := database.NewDatabaseConnection()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to connect to database")
		return
	}
	productCollection := collections.ProductCollectionInit(database.Database)
	products, err := productCollection.GetProductByPage(pageNumber, pageSizeNumber)
	wholeProducts, err := productCollection.GetAllProducts()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to get products")
		log.Println(err)
		return
	}
	totalProducts := len(wholeProducts)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		log.Println(err)
		return
	}

	if len(products) == 0 {
		utils.RespondWithError(c, http.StatusNotFound, "No products found")
		return
	}

	// Return the products as the API response
	c.JSON(http.StatusOK, gin.H{
		"page":          pageNumber,
		"pageSize":      pageSizeNumber,
		"products":      products,
		"totalProducts": totalProducts,
	})
}

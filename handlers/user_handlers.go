package handlers

import (
	"net/http"

	"github.com/20pa5a1210/Ecommerce-Gadgets-Backend/collections"
	"github.com/20pa5a1210/Ecommerce-Gadgets-Backend/database"
	"github.com/20pa5a1210/Ecommerce-Gadgets-Backend/models"
	"github.com/20pa5a1210/go-todo/middleware"
	"github.com/20pa5a1210/go-todo/utils"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if user.Password != user.ConfirmPassword {
		utils.RespondWithError(c, http.StatusForbidden, "Password and Confirm Password should be same")
		return
	}
	database, err := database.NewDatabaseConnection()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Database connection failed")
		return
	}

	userCollection := collections.UserCollectionInit(database.Database)

	existingUser, err := userCollection.GetUserByEmail(user.Email)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	if existingUser.Email != "" {
		utils.RespondWithError(c, http.StatusForbidden, "Email already exists")
		return
	}

	checkinguser, err := userCollection.GetUserByUsername(user.Username)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	if checkinguser.Username != "" {
		utils.RespondWithError(c, http.StatusForbidden, "Username already exists")
		return
	}

	CreateUser, err := userCollection.CreateUser(user)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create user")
		return
	}
	cartRepo := collections.CartCollectionInit(database.Database)
	cartInstance := models.UserCart{
		UserName: CreateUser.Username,
		Cart:     []models.Cart{},
	}
	_, err = cartRepo.CreateCart(cartInstance)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create cart")
		return
	}
	utils.RespondWithJSON(c, http.StatusCreated, "User:", CreateUser)

}

func LoginUser(c *gin.Context) {
	var LoginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&LoginData); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	database, err := database.NewDatabaseConnection()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	userCollection := collections.UserCollectionInit(database.Database)
	user, err := userCollection.GetUserByEmail(LoginData.Email)
	if err != nil {
		utils.RespondWithError(c, http.StatusUnauthorized, err.Error())
		return
	}
	if user.Password != LoginData.Password {
		utils.RespondWithError(c, http.StatusUnauthorized, "wrong Password(mismatch)")
		return
	}
	token, err := middleware.GenerateJwt(user.Id.Hex())
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed To generate Token")
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func GetProfile(c *gin.Context) {
	userID, _ := c.Get("user_id")
	database, err := database.NewDatabaseConnection()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	userCollection := collections.UserCollectionInit(database.Database)

	user, err := userCollection.GetUserByID(userID.(string))
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to Fetch user")
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

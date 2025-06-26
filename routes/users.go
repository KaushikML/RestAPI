package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

// signup godoc
// @Summary      Sign-up a new user
// @Description  Creates a user and returns a JWT
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        payload  body      models.User     true  "user credentials"
// @Success      201      {object}  models.Message
// @Failure 404 {object} models.ErrorResponse
// @Router       /signup [post]
func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// login godoc
// @Summary      Log-in and receive JWT
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        payload  body      models.UserLogin  true  "email & password"
// @Success      200      {object}  models.Token
// @Failure 404 {object} models.ErrorResponse
// @Router       /login [post]
func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})
}

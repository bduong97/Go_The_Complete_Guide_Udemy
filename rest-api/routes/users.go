package routes

import (
	"fmt"
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user) //storing body of request into variable
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request value"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not save user"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created succesfully"})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user) //storing body of request into variable
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request value"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user"})
		return
	}
	// After logging in, need to set the userID
	token, err := utils.GenerateToken(user.Email, user.ID) //utilize user ID and email to generate a token after authenticating
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token"})
		return
	}
	context.JSON(http.StatusOK, 
		gin.H{
			"message": "Login successful", 
			"token": token })
}

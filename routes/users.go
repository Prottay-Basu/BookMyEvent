package routes

import (
	"net/http"
	"prottaybasu/book-my-event/models"
	"prottaybasu/book-my-event/utils"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {

	var user models.User

	err := context.ShouldBind(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request body"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Sorry, could not save the user."})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"message": "User created successfully"})
}

func login(context *gin.Context) {

	var user models.User
	err := context.ShouldBind(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request body"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid credentials sent for the user"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Sorry, could not authenticate the user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login Successful", "token": token})

}

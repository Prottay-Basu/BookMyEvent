package routes

import (
	"net/http"
	"prottaybasu/book-my-event/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the event id"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event"})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register the user for event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User registered for the event"})

}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the event id"})
		return
	}

	var event models.Event

	event.ID = eventId
	err = event.CancelRegistrations(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel the user for event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User cancelled for the event"})
}

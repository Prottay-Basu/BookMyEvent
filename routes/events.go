package routes

import (
	"net/http"
	"prottaybasu/book-my-event/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Sorry, could not fetch data from events table"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {

	var event models.Event

	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request body"})
		return
	}

	userId := context.GetInt64("userId")
	event.UserID = userId
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Sorry, could not create event. Please try again later"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event is created", "event": event})
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the event id"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Sorry, could not fetch the event by the id"})
		return
	}

	context.JSON(http.StatusOK, event)

}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the event id"})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)

	if userId != event.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User not authorized to update the event"})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Sorry, could not fetch the event by the id"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request body"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Sorry, could not update the event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event Updated!"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the event id"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Sorry, could not fetch the event by the id"})
		return
	}

	userId := context.GetInt64("userId")

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User not authorized to delete the event"})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Sorry, could not delete the event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event Deleted!"})

}

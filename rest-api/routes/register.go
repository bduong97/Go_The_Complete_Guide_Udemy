package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	claims := context.GetStringMap("claims")
	UserID := int64(claims["userId"].(float64))
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID"})
		fmt.Println(err)
		return
	}
	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event by ID"})
		return
	}

	err = event.Register(UserID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user to event"})
	}
	context.JSON(http.StatusAccepted, gin.H{"message": "User registered for event"})
}

func cancelRegistration(context *gin.Context) {
	claims := context.GetStringMap("claims")
	UserID := int64(claims["userId"].(float64))
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	var event models.Event
	event.ID = eventID
	event.CancelRegistration(UserID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Cancelled"})
}

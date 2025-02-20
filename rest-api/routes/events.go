package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID"})
		fmt.Println(err)
		return
	}
	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get event by ID from DB"})
		fmt.Println(err)
		return
	}
	context.JSON(http.StatusOK, *event)
}
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events, Try again later"})
		fmt.Println(err)
		return
	}
	context.JSON(http.StatusOK, events) //list, gin package automatically transcribes to json format

}

func createEvent(context *gin.Context) {
	claims := context.GetStringMap("claims")

	var event models.Event                //empty struct
	err := context.ShouldBindJSON(&event) //storing body of request into variable
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request value"})
		return
	}
	event.UserID = int64(claims["userId"].(float64))
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events, Try again later"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func updateEvent(context *gin.Context) {
	claims := context.GetStringMap("claims")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID"})
		fmt.Println(err)
		return
	}

	//logic for comparing userID to eventID
	fetchedEvent, err := models.GetEventByID(eventID)
	UserID := int64(claims["userId"].(float64))
	if UserID != fetchedEvent.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User is not authorized to update event"})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event, Try again later"})
		fmt.Println(err)
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request value"})
		return
	}

	updatedEvent.ID = eventID
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not update event value"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated"})
}

func deleteEvent(context *gin.Context) {
	claims := context.GetStringMap("claims")

	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID"})
		fmt.Println(err)
		return
	}

	fetchedEvent, err := models.GetEventByID(eventID)
	UserID := int64(claims["userId"].(float64))
	if UserID != fetchedEvent.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User is not authorized to delete event"})
		return
	}
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event, Try again later"})
		fmt.Println(err)
		return
	}

	err = fetchedEvent.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event, Try again later"})
		fmt.Println(err)
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
}

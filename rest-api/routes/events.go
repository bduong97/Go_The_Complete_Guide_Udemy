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
	var event models.Event                //empty struct
	err := context.ShouldBindJSON(&event) //storing body of request into variable
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request value"})
		return
	}

	event.ID = 1
	event.UserID = 1
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events, Try again later"})
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

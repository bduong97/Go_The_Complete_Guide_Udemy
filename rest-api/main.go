package main

import (
	"net/http"
	"example.com/rest-api/models"
	"example.com/rest-api/db"
	"github.com/gin-gonic/gin"
)
func main() {
	db.InitDB()
	server :=	gin.Default()

	server.GET("/events", getEvents) 
	server.POST("/events", createEvent)
	server.Run(":8080") //localhost:8080
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events) //list, gin package automatically transcribes to json format

}

func createEvent(context *gin.Context) {
	var event models.Event //empty struct
	err := context.ShouldBindJSON(&event) //storing body of request into variable
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse request value"})
	}

	event.ID = 1
	event.UserID = 1
	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
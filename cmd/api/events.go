package main

import (
	"fmt"
	"go-event-crud/internal/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (app *application) createEvent (c *gin.Context) {
	var event database.Event

	// Return a 400 Bad Request if the JSON is invalid
	if err:= c.BindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err:= app.models.Events.Insert(&event)

	// Return a 500 Internal Server Error if the event could not be created
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create event : %s", err.Error())})
		return
	}

	// Return a 201 Created status code with the created event
	c.JSON(http.StatusCreated, event)
}


func (app *application) getEventById (c *gin.Context) {
	id,err := strconv.Atoi(c.Param("id"))

	// Return a 400 Bad Request if the ID is not a valid integer
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	event, err := app.models.Events.GetById(id)

	// Return a 404 Not Found if the event does not exist
	if event == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	// Return a 500 Internal Server Error if there was an error retrieving the event
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to retrieve event: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, event)
}

func (app *application) getAllEvents (c *gin.Context) {
	allEvents,err := app.models.Events.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to retrieve events: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, allEvents)
}

func (app *application) updateEvent (c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	existingEvent, err := app.models.Events.GetById(id)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to retrieve event: %s", err.Error())})
		return
	}

	if existingEvent == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	var updatedEvent database.Event

	if err := c.ShouldBindJSON(&updatedEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	updatedEvent.Id = id

	err = app.models.Events.Update(&updatedEvent)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to update event: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, updatedEvent)
}

func (app *application) deleteEvent (c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	err = app.models.Events.Delete(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to delete event: %s", err.Error())})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
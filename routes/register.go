package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

// registerForEvent godoc
// @Summary      Register current user for an event
// @Tags         events
// @Security     BearerAuth
// @Produce      json
// @Param        id  path  int  true  "Event ID"
// @Success      200  {object}  models.Message
// @Failure 404 {object} models.ErrorResponse
// @Router       /events/{id}/register [post]

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for event."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered!"})
}

// cancelRegistration godoc
// @Summary      Cancel event registration
// @Tags         events
// @Security     BearerAuth
// @Produce      json
// @Param        id  path  int  true  "Event ID"
// @Success      200  {object}  models.Message
// @Failure 404 {object} models.ErrorResponse   
// @Router       /events/{id}/register [delete]
func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Cancelled!"})
}

package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/models"
)
func ListAllEvent(c *gin.Context){
	listEvent := models.FindAllEvent()
		c.JSON(http.StatusOK, lib.Message{
			Success: true,
			Message: "success",
			Results: listEvent,
		})		
	}
func DetailEvent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := models.FindOneEvent(id)
	fmt.Println(id)

	if data.Id == id {
		c.JSON(http.StatusOK, lib.Message{
			Success: true,
			Message: "events Found",
			Results: data,
		})
		return
	} else {
		c.JSON(http.StatusNotFound, lib.Message{
			Success: false,
			Message: "events Not Found",
		})
	}
}
func CreateEvent(c *gin.Context) {
    newEvents := models.Event{}
    id, _ := c.Get("userId")
    err := models.CreateEvent(newEvents, id.(int))
    
    if err := c.ShouldBind(&newEvents); err != nil {
        c.JSON(http.StatusBadRequest, lib.Message{
            Success: false,
            Message: "Invalid input data",
        })
        return
    }
    
    if err != nil {
        c.JSON(http.StatusInternalServerError, lib.Message{
            Success: false,
            Message: "Failed to create Profile",
        })
        return
    }
    
    newEvents.CreateBy = id.(int)
    // createBy := c.Keys["userId"].(int)
    // newEvents.CreateBy = &createBy
    c.JSON(http.StatusOK, lib.Message{
        Success: true,
        Message: "Events created successfully",
        Results: newEvents,
    })

}
func DeleteEvent(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	dataEvent := models.FindOneEvent(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, lib.Message{
			Success:  false,
			Message: "Invalid Event ID",
		})
		return
	}

	err = models.RemoveEvent(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, lib.Message{
			Success:  false,
			Message: "Id Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, lib.Message{
		Success:  true,
		Message: "Event deleted successfully",
		Results: dataEvent,
	})

}

func UpdateEvent(c *gin.Context) {
    param := c.Param("id")
    id,_  := strconv.Atoi(param)
    data := models.FindAllEvent()
	

    event := models.Event{}
    err := c.Bind(&event)
    if err != nil {
        fmt.Println(err)
        return
    }

    result := models.Event{}
    for _, v := range data {
        if v.Id == id {
            result = v
        }
    }

    if result.Id == 0 {
        c.JSON(http.StatusNotFound, lib.Message{
            Success: false,
            Message: "event with id " + param + " not found",
        })
        return
    }

    idEvent := 0
    for _, v := range data {
        idEvent = v.Id
    }
    event.Id = idEvent

    models.EditEvent(*event.Image,*event.Title,*event.Date,*event.Descriptions, *event.LocationId, event.CreateBy, param)

    c.JSON(http.StatusOK, lib.Message{
        Success: true,
        Message: "Event  id " + param + " edit Success",
        Results:event,
    })
}
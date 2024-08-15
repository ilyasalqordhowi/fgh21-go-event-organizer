package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/models"
)

// type FormTransaction struct {
// 	EventId         int `form:"eventId"`
// 	SectionId       []int `form:"sectionId"`
// 	TicketQyt       []int `form:"ticketQyt"`
// 	PaymentMethodId int `form:"paymentMethodId"`

// }

func CreateTransaction(c *gin.Context) {
	newEvents := models.Transaction{}
    id := c.GetInt("userId")
    data , err := models.CreateTransaction(newEvents,id)
    
    if err := c.ShouldBind(&newEvents);
	 err != nil {
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
    c.JSON(http.StatusOK, lib.Message{
        Success: true,
        Message: "Events created successfully",
        Results: data,
    })
}

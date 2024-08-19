package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/models"
)

type FormTransactions struct{
	EventId         	int `json:"eventId" form:"eventId" db:"event_id"`
	PaymentMethodId 	int `json:"paymentMethodId" form:"paymentMethodId" db:"payment_method_id"`
	SectionId       	[]int `json:"sectionId" form:"sectionId" db:"section_id"`
	TicketQty       	[]int `json:"ticketQty" form:"ticketQty" db:"ticket_qty"`
}

func CreateTransaction(ctx *gin.Context) {
	form := FormTransactions{}

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Message{
			Success: false,
			Message: "invalid input data",
		})
		return
	}
	fmt.Println(form,"testttt")

  trx := models.CreateTransaction(models.Transaction{
		UserId: ctx.GetInt("userId"),
		PaymentMethodId: form.PaymentMethodId,
		EventId: form.EventId,
	})
    for i := range form.SectionId{
        models.CreateTransactionDetail(models.TransactionDetail{
            SectionId: form.SectionId[i],
			TicketQty: form.TicketQty[i],
			TransactionId: trx.Id,
        })
    }

	data,err := models.CreateDetailTransactions()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "Transaction created successfully",
		Results: data,
	})
}

func FindTransactionByUserId(ctx *gin.Context){
	UserId := ctx.GetInt("userId")

	detailTransactionbyId := models.FindOneTransactionById(UserId)

	ctx.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "Transaction User Id",
		Results: detailTransactionbyId,
	})
}
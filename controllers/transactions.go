package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/models"
)

func CreateTransaction(ctx *gin.Context) {
    form := models.Transaction{}
    err := ctx.ShouldBind(&form);
    if  err != nil {
        ctx.JSON(http.StatusBadRequest,
            lib.Message{
                Success: false,
                Message: "Created Transaction Failed",
               
            })
        return
    }

    trx := models.CreateNewTransactions(models.Transaction{
        UserId:    ctx.GetInt("userId"),
        PaymentMethodId: form.PaymentMethodId,
        EventId:   form.EventId,
    })
    for i := range form.SectionId {
		 models.CreateTransactionDetail(models.TransactionDetail{
			SectionId:      form.SectionId[i],
            TicketQuantity: form.TicketQty[i],
            TransactionId:  trx.Id,
        })
    }
    ctx.JSON(http.StatusOK,
        lib.Message{
            Success: true,
            Message: "Transaction success",
            Results: trx,
        })
}

func FindTransactionByUserId(ctx *gin.Context){
    id := ctx.GetInt("userId")

	result, err := models.DetailsTransaction(id)
	fmt.Print(err)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			lib.Message{
				Success: false,
				Message: "Transaction Not Found",
                Results: result,
			})
		return
	}
	ctx.JSON(http.StatusOK,
		lib.Message{
			Success: true,
			Message: "Transaction Found",
			Results: result,
		})
}
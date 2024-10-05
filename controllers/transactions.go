package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/dtos"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/repository"
	"github.com/jackc/pgx/v5"
)

func CreateTransaction(ctx *gin.Context) {
    form := dtos.Transaction{}
    err := ctx.ShouldBind(&form);
    if  err != nil {
        ctx.JSON(http.StatusBadRequest,
            lib.Message{
                Success: false,
                Message: "Created Transaction Failed",
               
            })
        return
    }
    tx, err := lib.DB().BeginTx(context.Background(), pgx.TxOptions{})
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, lib.Message{
            Success: false,
            Message: "Failed to start transaction",
        })
        return
    }
    trx, err := repository.CreateNewTransactions(tx, dtos.Transaction{
        UserId:    ctx.GetInt("userId"),
        PaymentMethodId: form.PaymentMethodId,
        EventId:   form.EventId,
    })
    if err != nil {
        tx.Rollback(context.Background())
        ctx.JSON(http.StatusInternalServerError, lib.Message{
            Success: false,
            Message: "Failed to create transaction: " + err.Error(),
        })
        return
    }
    for i := range form.SectionId {
		_,err :=  repository.CreateTransactionDetail(tx,dtos.TransactionDetail{
			SectionId:      form.SectionId[i],
            TicketQuantity: form.TicketQty[i],
            TransactionId:  trx.Id,
        })
        if err != nil {
            tx.Rollback(context.Background())
            ctx.JSON(http.StatusInternalServerError, lib.Message{
                Success: false,
                Message: "Failed to create transaction detail: " + err.Error(),
            })
            return
        }
    }
    if err := tx.Commit(context.Background()); err != nil {
        ctx.JSON(http.StatusInternalServerError, lib.Message{
            Success: false,
            Message: "Failed to commit transaction: " + err.Error(),
        })
        return
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

	result, err := repository.DetailsTransaction(id)
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
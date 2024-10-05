package controllers

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/dtos"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/lib"
	"github.com/ilyasalqordhowi/fgh21-go-event-organizer/repository"
	"github.com/jackc/pgx/v5"
)

func CreateTransaction(ctx *gin.Context) {
    form := dtos.Transaction{}
	err := ctx.ShouldBind(&form)
	if err != nil {
		lib.HandlerBadRequest(ctx, "Created Transaction Failed")
		return
	}

	tx, err := lib.DB().BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		lib.HandlerBadRequest(ctx, "Failed to start transaction")
		return
	}

	trx, err := repository.CreateNewTransactions(tx, dtos.Transaction{
		UserId:          ctx.GetInt("userId"),
		PaymentMethodId: form.PaymentMethodId,
		EventId:         form.EventId,
	})
	if err != nil {
		tx.Rollback(context.Background())
		lib.HandlerBadRequest(ctx, "Failed to create transaction: "+err.Error())
		return
	}

	for i := range form.SectionId {
		_, err := repository.CreateTransactionDetail(tx, dtos.TransactionDetail{
			SectionId:      form.SectionId[i],
			TicketQuantity: form.TicketQty[i],
			TransactionId:  trx.Id,
		})
		if err != nil {
			tx.Rollback(context.Background())
			lib.HandlerBadRequest(ctx, "Failed to create transaction detail: "+err.Error())
			return
		}
	}

	if err := tx.Commit(context.Background()); err != nil {
		lib.HandlerBadRequest(ctx, "Failed to commit transaction: "+err.Error())
		return
	}

	lib.HandlerOk(ctx, "Transaction success", nil, trx)
}

func FindTransactionByUserId(ctx *gin.Context){
    id := ctx.GetInt("userId")

	result, err := repository.DetailsTransaction(id)
	fmt.Print(err)
	if err != nil {
		lib.HandlerBadRequest(ctx, "Transaction Not Found")
		return
	}

	lib.HandlerOk(ctx, "Transaction Found", nil, result)
}
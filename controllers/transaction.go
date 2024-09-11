package controllers

import (
	"fmt"
	"net/http"

	"github.com/fariqmaasir/fgh21-go-event-organizer/lib"
	"github.com/fariqmaasir/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func CreateTransaction(ctx *gin.Context) {
	form := models.Transactions{}
	// id := ctx.GetInt("userId")
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	trx := models.CreateNewTransactions(models.Transactions{
		UserId:    ctx.GetInt("userId"),
		PaymentId: form.PaymentId,
		EventId:   form.EventId,
	})
	fmt.Println("test")
	for i := range form.SectionId {
		models.CreateTransactionDetail(models.TransactionDetail{
			SectionId:      form.SectionId[i],
			TicketQuantity: form.TicketQuantity[i],
			TransactionId:  trx.Id,
		})
	}
	fmt.Println("test")
	details, err := models.AddDetailsTransaction()
	// fmt.Println(result)
	fmt.Println("test")
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			lib.Message{
				Success: false,
				Message: "Transaction Failed To Created",
				// Results: result,
			})
		return
	}
	ctx.JSON(http.StatusOK,
		lib.Message{
			Success: true,
			Message: "Create Transaction success",
			Results: details,
		})

}

// func GetOneTransaction(ctx *gin.Context) {
// 	id := ctx.GetInt("userId")
// 	// data, err := models.ListOneTransaction(id)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	ctx.JSON(http.StatusOK,
// 		lib.Message{
// 			Success: true,
// 			Message: "OK",
// 			Results: data,
// 		})
// }

func FindTransactionByUserId(ctx *gin.Context) {
	id := ctx.GetInt("userId")

	result, err := models.ListAllTransactionById(id)
	fmt.Print(err)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			lib.Message{
				Success: false,
				Message: "Transaction Not Found",
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

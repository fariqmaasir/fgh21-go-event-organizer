package controllers

import (
	"net/http"

	"github.com/fariqmaasir/fgh21-go-event-organizer/lib"
	"github.com/fariqmaasir/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func FindAllNationalities(ctx *gin.Context) {
	result, err := models.ListAllNationalities()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, lib.Message{
			Success: false,
			Message: "Payment Not Found",
			Results: result,
		})
	}
	ctx.JSON(http.StatusOK, lib.Message{
		Success: true,
		Message: "List All Payment",
		Results: result,
	})
}

package controllers

import (
	"fmt"
	"net/http"

	"github.com/fariqmaasir/fgh21-go-event-organizer/lib"
	"github.com/fariqmaasir/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

func FindAllPartners(ctx *gin.Context) {
	result, err := models.ListAllPartners()
	fmt.Print(err)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			lib.Message{
				Success: false,
				Message: "Partners Not Found",
			})
		return
	}
	ctx.JSON(http.StatusOK,
		lib.Message{
			Success: true,
			Message: "Partners Found",
			Results: result,
		})
}

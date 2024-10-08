package controllers

import (
	"net/http"

	"github.com/fariqmaasir/fgh21-go-event-organizer/lib"
	"github.com/fariqmaasir/fgh21-go-event-organizer/models"
	"github.com/gin-gonic/gin"
)

type Token struct {
	JWToken string `json:"token"`
}

func AuthLogin(ctx *gin.Context) {
	var user models.User
	ctx.Bind(&user)

	found := models.FindOneUserByEmail(user.Email)

	if found == (models.User{}) {
		ctx.JSON(http.StatusUnauthorized,
			lib.Message{
				Success: false,
				Message: "Wrong Email or Password",
			})
		return
	}

	isVerified := lib.Verify(user.Password, found.Password)

	if isVerified {
		JWToken := lib.GenerateUserIdToken(found.Id)
		ctx.JSON(http.StatusOK,
			lib.Message{
				Success: true,
				Message: "OK",
				Results: Token{JWToken},
			})
	} else {
		ctx.JSON(http.StatusUnauthorized,
			lib.Message{
				Success: false,
				Message: "Wrong Email or Password",
				// Results: "Wrong Email or Password",
			})
	}
}

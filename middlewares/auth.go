package middlewares

import (
	"net/http"

	"github.com/fariqmaasir/fgh21-go-event-organizer/lib"
	"github.com/gin-gonic/gin"
)

func tokenFailed(ctx *gin.Context) {
	if e := recover(); e != nil {
		ctx.JSON(http.StatusUnauthorized, lib.Message{
			Success: false,
			Message: "Unauthorized",
		})
		ctx.Abort()
	}
}

func AuthMiddlewares() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer tokenFailed(ctx)
		token := ctx.GetHeader("Authorization")[7:]
		isValidated, userId := lib.ValidateToken(token)
		if isValidated {
			ctx.Set("userId", userId)
			ctx.Next()
		} else {
			panic("Error: Token invalid")
		}
	}
}

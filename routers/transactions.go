package routers

import (
	"github.com/fariqmaasir/fgh21-go-event-organizer/controllers"
	"github.com/fariqmaasir/fgh21-go-event-organizer/middlewares"
	"github.com/gin-gonic/gin"
)

func TransactionRouter(routerGroup *gin.RouterGroup) {
	routerGroup.Use(middlewares.AuthMiddlewares())
	routerGroup.POST("/payment", controllers.CreateTransaction)
	routerGroup.GET("/:id", controllers.GetOneTransaction)
	routerGroup.GET("/users", controllers.FindTransactionByUserId)
}

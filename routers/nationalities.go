package routers

import (
	"github.com/fariqmaasir/fgh21-go-event-organizer/controllers"
	"github.com/fariqmaasir/fgh21-go-event-organizer/middlewares"
	"github.com/gin-gonic/gin"
)

func NationalitiesRouter(routerGroup *gin.RouterGroup) {
	routerGroup.Use(middlewares.AuthMiddlewares())
	routerGroup.GET("", controllers.FindAllNationalities)
}

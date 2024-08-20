package routers

import (
	"github.com/fariqmaasir/fgh21-go-event-organizer/controllers"
	"github.com/gin-gonic/gin"
)

func PartnersRouter(routerGroup *gin.RouterGroup) {
	routerGroup.GET("", controllers.FindAllPartners)
}

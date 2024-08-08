package routers

import (
	"github.com/fariqmaasir/fgh21-go-event-organizer/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRouter(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/login", controllers.AuthLogin)
	routerGroup.POST("/register", controllers.CreateProfileUser)
}

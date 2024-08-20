package routers

import (
	"github.com/fariqmaasir/fgh21-go-event-organizer/controllers"
	"github.com/fariqmaasir/fgh21-go-event-organizer/middlewares"
	"github.com/gin-gonic/gin"
)

func AuthRouter(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/profile", middlewares.AuthMiddlewares(), controllers.DetailProfile)
	routerGroup.PATCH("/edit", middlewares.AuthMiddlewares(), controllers.EditProfileUser)
	routerGroup.PATCH("/password", middlewares.AuthMiddlewares(), controllers.ChangeUserPassword)
	routerGroup.POST("/login", controllers.AuthLogin)
	routerGroup.POST("/register", controllers.CreateProfileUser)
}

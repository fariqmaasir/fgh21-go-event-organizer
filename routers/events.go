package routers

import (
	"github.com/fariqmaasir/fgh21-go-event-organizer/controllers"
	"github.com/gin-gonic/gin"
)

func EventRouters(r *gin.RouterGroup) {
	r.POST("create", controllers.CreateEvents)
}

package routers

import (
	"github.com/fariqmaasir/fgh21-go-event-organizer/controllers"
	"github.com/fariqmaasir/fgh21-go-event-organizer/middlewares"
	"github.com/gin-gonic/gin"
)

func EventRouters(r *gin.RouterGroup) {
	r.GET("/list", controllers.ListAllEvents)
	r.GET("/list/:id", controllers.DetailEvent)
	r.GET("/section/:id", controllers.SectionEvent)
	r.GET("/users", middlewares.AuthMiddlewares(), controllers.UserCreatedEvent)
	r.GET("/payment", middlewares.AuthMiddlewares(), controllers.FindAllPaymentMethod)
	r.POST("/create", middlewares.AuthMiddlewares(), controllers.CreateEvents)
	r.PATCH("/edit/:id", middlewares.AuthMiddlewares(), controllers.UpdateEvent)
	r.DELETE("/:id", middlewares.AuthMiddlewares(), controllers.DeleteEvent)
	//wishlist
	// r.GET("/users", middlewares.AuthMiddlewares(), controllers.UserWishlist)
	r.POST("/wishlist/:id", middlewares.AuthMiddlewares(), controllers.CreateWishlist)
	r.DELETE("/wishlist/:id", middlewares.AuthMiddlewares(), controllers.DeleteWishlist)
	r.GET("/wishlist/user", middlewares.AuthMiddlewares(), controllers.FindWishlist)
}

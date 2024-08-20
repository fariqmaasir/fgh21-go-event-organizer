package routers

import "github.com/gin-gonic/gin"

func RouterCombine(r *gin.Engine) {
	UserRouter(r.Group("/users"))
	AuthRouter(r.Group("/auth"))
	EventRouters(r.Group("/events"))
	TransactionRouter(r.Group("/transactions"))
	NationalitiesRouter(r.Group("/nationality"))
	PartnersRouter(r.Group("/partners"))
}

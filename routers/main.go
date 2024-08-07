package routers

import "github.com/gin-gonic/gin"

func RouterCombine(r *gin.Engine) {
	UserRouter(r.Group("/users"))
	AuthRouter(r.Group("/auth"))
	EventRouters(r.Group("/events"))
}

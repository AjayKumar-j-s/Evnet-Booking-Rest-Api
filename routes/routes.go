package routes

import (
	"github.com/AjayKumar-j-s/EventPlanner/middleware"
	"github.com/gin-gonic/gin"
)

func Registerroutes(server *gin.Engine){

	server.GET("/events",GetEvents)

	server.GET("/event/:id",GetEvent)
	
	authenticate := server.Group("/")
	authenticate.Use(middleware.Authenticate)
	authenticate.POST("/events",CreateEvents)
	authenticate.PUT("/event/:id",Updateevent)
	authenticate.DELETE("/event/:id",DeleteEvent)
	authenticate.POST("/event/:id/register",RegisterEvent)
	authenticate.DELETE("/event/:id/register",CancelRegistration)


	server.POST("/signup",Signup)

	server.POST("/login",Login)

}



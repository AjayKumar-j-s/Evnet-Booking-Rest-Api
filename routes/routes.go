package routes

import (
	"github.com/gin-gonic/gin"
)

func Registerroutes(server *gin.Engine){

	server.GET("/events",GetEvents)

	server.GET("/event/:id",GetEvent)
	
	server.POST("/events",CreateEvents)

	server.PUT("/event/:id",Updateevent)

	server.DELETE("/event/:id",DeleteEvent)

	server.POST("/signup",Signup)

	server.POST("/login",Login)

}



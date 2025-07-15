package main

import (

	"github.com/AjayKumar-j-s/EventPlanner/db"
	"github.com/AjayKumar-j-s/EventPlanner/routes"
	"github.com/gin-gonic/gin"
)

func main(){

	db.InitDB()

	server := gin.Default()
	routes.Registerroutes(server)
	server.Run(":8080")
}


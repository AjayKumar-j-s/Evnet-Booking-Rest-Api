package routes

import (
	"net/http"
	"strconv"

	"github.com/AjayKumar-j-s/EventPlanner/models"
	"github.com/gin-gonic/gin"
)

func RegisterEvent(context *gin.Context){

	userId := context.GetInt64("id")
	event_id,err := strconv.ParseInt(context.Param("id"),10,64)

	if(err != nil){
		context.JSON(http.StatusBadRequest,gin.H{"message":"couid not fetch..Try Again"})
		return
	}

	// event,err := models.GetEvent(event_id)

	// if(err!=nil){
	// 	context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not execute"})
	// 	return 
	// }

	var r models.Reg
	r.Eventid = event_id
	r.Userid = userId

	err = r.Register()

	if(err!=nil){
	context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not execute"})
	 	return 
	}

	context.JSON(http.StatusOK,gin.H{"message":"Successfull registered"})



}	

func CancelRegistration(context *gin.Context){

}
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/AjayKumar-j-s/EventPlanner/models"
	"net/http"
	"fmt"
	"strconv"
)

func GetEvents(context *gin.Context){

	events,err := models.GetEvents()
	if(err!=nil){
	tpe := fmt.Sprint(err)
	context.JSON(http.StatusInternalServerError,gin.H{"message":tpe})
	return
	}
	context.JSON(http.StatusOK,events)
}	


func CreateEvents(context *gin.Context){
	var event models.Event

	//this will collectt the data from the post request and add the respected data to the struct variable

	err := context.ShouldBindJSON(&event)

	if(err != nil){
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not parse provide the essential information"})
		return
	}

	err = event.Save()


	if(err != nil){
		res := fmt.Sprint(err)

		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not parse provide the essential informatio","err":res})
		return

	}

	context.JSON(http.StatusCreated,gin.H{"message": "Event crreated!","event":event})


}


func GetEvent(context *gin.Context){

	id,err := strconv.ParseInt(context.Param("id"),10,64)

	if(err != nil){
		context.JSON(http.StatusBadRequest,gin.H{"message":"couid not fetch..Try Again"})
		return
	}

	event,err := models.GetEvent(id)

	if(err != nil){
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not fetch event"})
		return
	}
	context.JSON(http.StatusOK,event)
}

func Updateevent(context *gin.Context){

	id,err := strconv.ParseInt(context.Param("id"),10,64)

	if(err != nil){
		context.JSON(http.StatusBadRequest,gin.H{"message":"couid not fetch..Try Again"})
		return
	}

	_,err = models.GetEvent(id)

	if(err != nil){
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not fetch event"})
		return
	}
	


	var e models.Event

	err = context.ShouldBind(&e)

	if(err != nil){
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not fetch data from req"})
		return
	}

	e.ID = id

	err = e.UpdateEve()

	if(err != nil){
		res := fmt.Sprint(err)
		context.JSON(http.StatusInternalServerError,gin.H{"message":res})
		return
	}
	

	context.JSON(http.StatusOK,gin.H{"message":"Event updates Successfully"})
}



func DeleteEvent(context *gin.Context){
	id,err := strconv.ParseInt(context.Param("id"),10,64)

	if(err != nil){
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not delete Event"})
		return
	}

	err = models.DeleteEvent(id)

	if(err != nil){
		context.JSON(http.StatusInternalServerError,gin.H{"message":"Failed To Delete Event"})
		return
	}

	context.JSON(http.StatusOK,gin.H{"message":"Deleted Successful"})
}
package routes

import (
	"net/http"
	"fmt"
	"github.com/AjayKumar-j-s/EventPlanner/models"
	"github.com/gin-gonic/gin"
)

func Signup(context *gin.Context){

	var u models.User

	err := context.ShouldBindJSON(&u)

	if(err != nil){
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not parse provide the essential information of Users"})

		return 
	}

	err = u.Save()

	if(err != nil){
		res := fmt.Sprint(err)

		context.JSON(http.StatusInternalServerError,gin.H{"message":"Could not signin","err":res})
		return

	}


	context.JSON(http.StatusOK,gin.H{"message":"successfulled signed Up"})


}


func Login(context *gin.Context){

	var u models.User

	err := context.ShouldBindJSON(&u)

	if(err != nil){
		context.JSON(http.StatusBadRequest,gin.H{"message":"Could not login"})
		return
	}

	err = u.ValidatePassword()

	if(err != nil){
		context.JSON(http.StatusUnauthorized,gin.H{"message":"invalid Credentials"})
		return
	}

	context.JSON(200,gin.H{"message":"Login Successful"})




}
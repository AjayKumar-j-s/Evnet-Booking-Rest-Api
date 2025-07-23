package middleware

import (
	"net/http"

	"github.com/AjayKumar-j-s/EventPlanner/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context){
	token := context.Request.Header.Get("Authorization")//extracting token from http request header

	if(token == ""){

		context.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"message":"Not Authorized empty","err":token})
		return
	
	}
	id,err := utils.VerifyToken(token)
	if(err != nil){
		context.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"message":"Not Authorized verify","err":err})
		return
	}

	context.Set("uid",id)

	context.Next()
}
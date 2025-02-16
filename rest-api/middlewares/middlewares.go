package middlewares

import (
	"fmt"
	"net/http"

	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		//AbortWithStatusJSON aborts the request if it doesn't pass the middleware
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"}) 
		return
	}

	claims, err := utils.VerifyToken(token)
	if err != nil {
		fmt.Println(err)
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	//converting Map.Claims to map[string]interface{}
	//refer to GenerateToken in utils.jwt for list of claims
	genericMap := make(map[string]interface{})
	for key, value := range claims {
		genericMap[key] = value
	}
	context.Set("claims", genericMap)
	context.Next()
}

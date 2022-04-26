package middleware

import (
	"errors"
	"fmt"
	"sondr-backend/utils/response"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// tracingHeaders - Headers used in tracing
var tracingHeaders = [9]string{
	"x-request-id",
	"x-b3-traceid",
	"x-b3-spanid",
	"x-b3-parentspanid",
	"x-b3-sampled",
	"x-b3-flags",
	"x-ot-span-context",
	"x-cloud-trace-context",
	"traceparent",
}

// TracingMiddleware - middleware
func TracingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		value := c.GetHeader("Authorization")
		fmt.Println("header", value)
		if value == "" {
			c.AbortWithStatusJSON(400, response.ErrorMessage(400, errors.New("Token Not Found")))
			return
		}
		tokendata := strings.Split(value, " ")
		fmt.Println("token ", tokendata[1])
		token, err := jwt.Parse(tokendata[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing")
			}
			return []byte(viper.GetString("secret.Key")), nil
		})
		if err != nil {
			//var err Error
			//err = SetError(err, "Your Token has been expired")
			c.AbortWithStatusJSON(400, response.ErrorMessage(400, errors.New("Invalid Token")))
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("role", claims["role"])
		}
		c.Next()

	}
}

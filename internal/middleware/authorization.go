// package middleware

// import (
// 	"net/http"
// 	"strings"

// 	"github.com/gin-gonic/gin"
// )

// func JWTGate(allowedRole, ...string) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		bearerToken := ctx.GetHeader("Authorization")
// 		if bearerToken == "" {
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 				"message": "please login first",
// 			})
// 		}

// 		if !strings.Contains(bearerToken, "Bearer") {
// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 				"message": "please login again",
// 			})
// 		}

// 		token := strings.Replace(bearerToken, "Bearer")
// 	}

	
// }
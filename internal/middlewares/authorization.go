package middlewares

import (
	"coffee-shop-golang/internal/models"
	"coffee-shop-golang/internal/repositories"
	"coffee-shop-golang/pkg"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTGate(allowedRole []string, r *repositories.RepositoryAuth) gin.HandlerFunc {
	return func (ctx *gin.Context){
		bearerToken := ctx.GetHeader("Authorization")
		if bearerToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "please login first",
			})
			return
		}

		if !strings.Contains(bearerToken, "Bearer") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "please login again",
			})
			return
		}

		token := strings.Replace(bearerToken, "Bearer ", "", -1)

		//cek bearer token from db
		result := []models.JwtUsers{}
		query := `SELECT users_id, token_jwt FROM users_tokenjwt WHERE token_jwt = $1`
		if err := r.Select(&result, query, token); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		if len(result) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token, please login again",
			})
			return
		}

		payload, err := pkg.VerifyToken(token)
		if err != nil {
			if strings.Contains(err.Error(), "expired") {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": "token expired please login again",
				})
				return
			}
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}

		var allowed = false
		for _, role := range allowedRole {
			if payload.Roles_id == role {
				allowed = true
				break
			}
		}

		if !allowed {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "you dont have permission",
			})
			return
		}
		ctx.Set("Payload", payload)
		ctx.Next()
	}
}
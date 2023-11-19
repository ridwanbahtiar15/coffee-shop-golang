package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func New(db *sqlx.DB) *gin.Engine {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})

	RouterCategories(router, db)

	return router
}
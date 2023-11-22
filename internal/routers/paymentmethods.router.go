package routers

import (
	"coffee-shop-golang/internal/handlers"
	"coffee-shop-golang/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RouterPaymentmethods(g *gin.Engine, db *sqlx.DB) {
	route := g.Group("/paymentmethods")
	repository := repositories.InitializeRepoPaymentmethods(db)
	handler := handlers.InitializeHandlerPaymentmethods(repository)

	route.GET("/", handler.GetAllPaymentmethods)
}
package routers

import (
	"coffee-shop-golang/internal/handlers"
	"coffee-shop-golang/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RouterOrders(g *gin.Engine, db *sqlx.DB) {
	route := g.Group("/orders")
	repository := repositories.InitializeRepoOrders(db)
	handler := handlers.InitializeHandlerOrders(repository)

	route.GET("/", handler.GetAllOrders)
	route.POST("/", handler.CreateOrders)
}
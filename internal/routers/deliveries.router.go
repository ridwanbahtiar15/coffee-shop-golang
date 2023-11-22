package routers

import (
	"coffee-shop-golang/internal/handlers"
	"coffee-shop-golang/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RouterDeliveries(g *gin.Engine, db *sqlx.DB) {
	route := g.Group("/deliveries")
	repository := repositories.InitializeRepoDeliveries(db)
	handler := handlers.InitializeHandlerDeliveries(repository)

	route.GET("/", handler.GetAllDeliveries)
}
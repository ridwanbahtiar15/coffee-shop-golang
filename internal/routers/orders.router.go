package routers

import (
	"coffee-shop-golang/internal/handlers"
	"coffee-shop-golang/internal/middlewares"
	"coffee-shop-golang/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RouterOrders(g *gin.Engine, db *sqlx.DB) {
	route := g.Group("/orders")
	repository := repositories.InitializeRepoOrders(db)
	handler := handlers.InitializeHandlerOrders(repository)
	repositoryAuth := repositories.InitializeRepoAuth(db)

	route.GET("/", middlewares.JWTGate([]string{"1"}, repositoryAuth), handler.GetAllOrders)
	route.POST("/", middlewares.JWTGate([]string{"1", "2"}, repositoryAuth), handler.CreateOrders)
	route.PATCH("/:id", middlewares.JWTGate([]string{"1"}, repositoryAuth), handler.UpdateOrders)
}
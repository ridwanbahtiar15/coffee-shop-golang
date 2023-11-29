package routers

import (
	"coffee-shop-golang/internal/handlers"
	"coffee-shop-golang/internal/middlewares"
	"coffee-shop-golang/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RouterPromos(g *gin.Engine, db *sqlx.DB) {
	route := g.Group("/promos")
	repository := repositories.InitializeRepoPromos(db)
	handler := handlers.InitializeHandlerPromos(repository)
	repositoryAuth := repositories.InitializeRepoAuth(db)

	route.GET("/", middlewares.JWTGate([]string{"1", "2"}, repositoryAuth), handler.GetAllPromos)
	route.POST("/", middlewares.JWTGate([]string{"1"}, repositoryAuth), handler.CreateProomos)
	route.PATCH("/:id", middlewares.JWTGate([]string{"1"}, repositoryAuth), handler.UpdatePromos)
	route.DELETE("/:id", middlewares.JWTGate([]string{"1"}, repositoryAuth), handler.DeletePromos)
}
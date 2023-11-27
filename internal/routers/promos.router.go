package routers

import (
	"coffee-shop-golang/internal/handlers"
	"coffee-shop-golang/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RouterPromos(g *gin.Engine, db *sqlx.DB) {
	route := g.Group("/promos")
	repository := repositories.InitializeRepoPromos(db)
	handler := handlers.InitializeHandlerPromos(repository)

	route.GET("/", handler.GetAllPromos)
	route.POST("/", handler.CreateProomos)
	route.PATCH("/:id", handler.UpdatePromos)
	route.DELETE("/:id", handler.DeletePromos)
}
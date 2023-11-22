package routers

import (
	"coffee-shop-golang/internal/handlers"
	"coffee-shop-golang/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RouterRoles(g *gin.Engine, db *sqlx.DB) {
	route := g.Group("/roles")
	repository := repositories.InitializeRepoRoles(db)
	handler := handlers.InitializeHandlerRoles(repository)

	route.GET("/", handler.GetAllRoles)
}
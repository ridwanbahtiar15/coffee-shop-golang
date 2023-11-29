package routers

import (
	"coffee-shop-golang/internal/handlers"
	"coffee-shop-golang/internal/middlewares"
	"coffee-shop-golang/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RouterCategories(g *gin.Engine, db *sqlx.DB) {
	route := g.Group("/categories")
	repository := repositories.InitializeRepoCategories(db)
	handler := handlers.InitializeHandlerCategories(repository)
	repositoryAuth := repositories.InitializeRepoAuth(db)

	route.GET("/", middlewares.JWTGate([]string{"1", "2"}, repositoryAuth), handler.GetAllCategories)
}
package routers

import (
	"coffee-shop-golang/internal/handlers"
	"coffee-shop-golang/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RouterCategories(g *gin.Engine, db *sqlx.DB) {
	route := g.Group("/categories")
	repository := repositories.InitializeRepository(db)
	handler := handlers.InitializeHandler(repository)

	route.GET("/", handler.GetAllCategories)
	route.POST("/", handler.CreateCategories)
	route.PATCH("/:id", handler.UpdateCategories)
	route.DELETE("/:id", handler.DeleteCategories)
}
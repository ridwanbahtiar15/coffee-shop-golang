package routers

import (
	"coffee-shop-golang/internal/handlers"
	"coffee-shop-golang/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RouterProducts(g *gin.Engine, db *sqlx.DB) {
	route := g.Group("/products")
	repository := repositories.InitializeRepoProducts(db)
	handler := handlers.InitializeHandlerProducts(repository)

	route.GET("/", handler.GetAllProducts)
	route.GET("/:id", handler.GetProductsById)
	route.POST("/", handler.CreateProducts)
	route.PATCH("/:id", handler.UpdateProducts)
	route.DELETE("/:id", handler.DeleteProducts)
}
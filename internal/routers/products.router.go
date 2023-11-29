package routers

import (
	"coffee-shop-golang/internal/handlers"
	"coffee-shop-golang/internal/middlewares"
	"coffee-shop-golang/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RouterProducts(g *gin.Engine, db *sqlx.DB) {
	route := g.Group("/products")
	repository := repositories.InitializeRepoProducts(db)
	handler := handlers.InitializeHandlerProducts(repository)
	repositoryAuth := repositories.InitializeRepoAuth(db)

	route.GET("/", middlewares.JWTGate([]string{"1", "2"}, repositoryAuth), handler.GetAllProducts)
	route.GET("/:id", middlewares.JWTGate([]string{"1", "2"}, repositoryAuth), handler.GetProductsById)
	route.POST("/", middlewares.JWTGate([]string{"1"}, repositoryAuth), handler.CreateProducts)
	route.PATCH("/:id", middlewares.JWTGate([]string{"1"}, repositoryAuth), handler.UpdateProducts)
	route.DELETE("/:id", middlewares.JWTGate([]string{"1"}, repositoryAuth), handler.DeleteProducts)
}
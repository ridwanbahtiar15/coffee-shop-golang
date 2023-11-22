package routers

import (
	"coffee-shop-golang/internal/handlers"
	"coffee-shop-golang/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RouterUsers(g *gin.Engine, db *sqlx.DB) {
	route := g.Group("/users")
	repository := repositories.InitializeRepoUsers(db)
	handler := handlers.InitializeHandlerUsers(repository)

	route.GET("/", handler.GetAllUsers)
	route.GET("/:id", handler.GetUsersById)
	route.POST("/", handler.CreateUsers)
	route.PATCH("/:id", handler.UpdateUsers)
	route.DELETE("/:id", handler.DeleteUsers)
}
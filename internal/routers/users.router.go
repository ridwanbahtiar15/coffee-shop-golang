package routers

import (
	"coffee-shop-golang/internal/handlers"
	"coffee-shop-golang/internal/middlewares"
	"coffee-shop-golang/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RouterUsers(g *gin.Engine, db *sqlx.DB) {
	route := g.Group("/users")
	repository := repositories.InitializeRepoUsers(db)
	handler := handlers.InitializeHandlerUsers(repository)
	repositoryAuth := repositories.InitializeRepoAuth(db)

	route.GET("/", middlewares.JWTGate([]string{"1"}, repositoryAuth), handler.GetAllUsers)
	route.GET("/:id", middlewares.JWTGate([]string{"1"}, repositoryAuth), handler.GetUsersById)
	route.GET("/profile", middlewares.JWTGate([]string{"1", "2"}, repositoryAuth), handler.GetUserProfile)
	route.POST("/", middlewares.JWTGate([]string{"1"}, repositoryAuth), handler.CreateUsers)
	route.PATCH("/:id", middlewares.JWTGate([]string{"1"}, repositoryAuth), handler.UpdateUsers)
	route.DELETE("/:id", middlewares.JWTGate([]string{"1"}, repositoryAuth), handler.DeleteUsers)
	route.PATCH("/profile", middlewares.JWTGate([]string{"1", "2"}, repositoryAuth), handler.UserProfile)
}
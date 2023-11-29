package routers

import (
	"coffee-shop-golang/internal/handlers"
	"coffee-shop-golang/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RouterAuth(g *gin.Engine, db *sqlx.DB) {
	route := g.Group("/auth")
	repository := repositories.InitializeRepoAuth(db)
	handler := handlers.InitializeHandlerAuth(repository)

	route.POST("/register", handler.RegisterUsers)
	route.POST("/login", handler.LoginUsers)
	route.DELETE("/logout", handler.LogoutUsers)
}
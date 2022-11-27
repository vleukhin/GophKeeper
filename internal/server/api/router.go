package api

import (
	"github.com/gin-gonic/gin"

	"github.com/vleukhin/GophKeeper/internal/pkg/logger"
	"github.com/vleukhin/GophKeeper/internal/server/core"
)

type Router struct {
	uc core.GophKeeperServer
	l  logger.Interface
}

func NewRouter(e *gin.Engine, g core.GophKeeperServer, l logger.Interface) {
	r := &Router{g, l}
	h := e.Group("/api")
	e.GET("/health", r.HealthCheck)

	needAuth := h.Group("", r.AuthMiddleware())

	needAuth.GET("me", r.UserInfo)

	needAuth.GET("creds", r.GetCreds)
	needAuth.POST("creds", r.AddCred)
	needAuth.DELETE("creds/:id", r.DelCred)
	needAuth.PATCH("creds/:id", r.UpdateCred)

	needAuth.GET("cards", r.GetCards)
	needAuth.POST("cards", r.AddCard)
	needAuth.DELETE("cards/:id", r.DelCard)
	needAuth.PATCH("cards/:id", r.UpdateCard)

	needAuth.GET("notes", r.GetNotes)
	needAuth.POST("notes", r.AddNote)
	needAuth.DELETE("notes/:id", r.DelNote)
	needAuth.PATCH("notes/:id", r.UpdateNote)

	auth := h.Group("/auth")
	{
		auth.POST("/register", r.Register)
		auth.POST("/login", r.LogIn)
		auth.GET("/refresh", r.RefreshAccessToken)
		auth.GET("/logout", r.LogoutUser)
	}
}

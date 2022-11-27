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
	h := e.Group("/api", r.AuthMiddleware())

	e.GET("/health", r.HealthCheck)

	h.GET("me", r.UserInfo)

	h.GET("creds", r.GetCreds)
	h.POST("creds", r.AddCred)
	h.DELETE("creds/:id", r.DelCred)
	h.PATCH("creds/:id", r.UpdateCred)

	h.GET("cards", r.GetCards)
	h.POST("cards", r.AddCard)
	h.DELETE("cards/:id", r.DelCard)
	h.PATCH("cards/:id", r.UpdateCard)

	h.GET("notes", r.GetNotes)
	h.POST("notes", r.AddNote)
	h.DELETE("notes/:id", r.DelNote)
	h.PATCH("notes/:id", r.UpdateNote)

	authAPI := h.Group("/auth")
	{
		authAPI.POST("/register", r.Register)
		authAPI.POST("/login", r.LogIn)
		authAPI.GET("/refresh", r.RefreshAccessToken)
		authAPI.GET("/logout", r.LogoutUser)
	}
}

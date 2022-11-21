package v1

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
	h := e.Group("/api/v1")

	r := &Router{g, l}

	h.GET("/health", r.HealthCheck)

	h.GET("me", r.ProtectedByAccessToken(), r.UserInfo)

	h.GET("logins", r.ProtectedByAccessToken(), r.GetLogins)
	h.POST("logins", r.ProtectedByAccessToken(), r.AddLogin)
	h.DELETE("logins/:id", r.ProtectedByAccessToken(), r.DelLogin)
	h.PATCH("logins/:id", r.ProtectedByAccessToken(), r.UpdateLogin)

	h.GET("cards", r.ProtectedByAccessToken(), r.GetCards)
	h.POST("cards", r.ProtectedByAccessToken(), r.AddCard)
	h.DELETE("cards/:id", r.ProtectedByAccessToken(), r.DelCard)
	h.PATCH("cards/:id", r.ProtectedByAccessToken(), r.UpdateCard)

	h.GET("notes", r.ProtectedByAccessToken(), r.GetNotes)
	h.POST("notes", r.ProtectedByAccessToken(), r.AddNote)
	h.DELETE("notes/:id", r.ProtectedByAccessToken(), r.DelNote)
	h.PATCH("notes/:id", r.ProtectedByAccessToken(), r.UpdateNote)

	authAPI := h.Group("/auth")
	{
		authAPI.POST("/register", r.SignUpUser)
		authAPI.POST("/login", r.SignInUser)
		authAPI.GET("/refresh", r.RefreshAccessToken)
		authAPI.GET("/logout", r.LogoutUser)
	}
}

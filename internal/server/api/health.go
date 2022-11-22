package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) HealthCheck(ctx *gin.Context) {
	err := r.uc.HealthCheck(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())

		return
	}
}

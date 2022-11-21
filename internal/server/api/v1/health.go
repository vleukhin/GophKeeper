package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) HealthCheck(ctx *gin.Context) {
	err := r.uc.HealthCheck()
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "connected"})
}

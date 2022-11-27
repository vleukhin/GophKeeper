package api

import (
	"errors"
	"net/http"

	"github.com/vleukhin/GophKeeper/internal/models"

	"github.com/gin-gonic/gin"

	"github.com/vleukhin/GophKeeper/internal/helpers/errs"
)

func (r *Router) Register(ctx *gin.Context) {
	var payload *models.LoginPayload

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	user, err := r.uc.SignUpUser(ctx, payload.Name, payload.Password)
	if err == nil {
		ctx.JSON(http.StatusOK, user)

		return
	}

	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

func (r *Router) LogIn(ctx *gin.Context) {
	var payload *models.LoginPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	jwt, err := r.uc.SignInUser(ctx, payload.Name, payload.Password)

	if err == nil {
		ctx.JSON(http.StatusOK, jwt)
		return
	}

	if errors.Is(err, errs.ErrWrongCredentials) {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

func (r *Router) RefreshAccessToken(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "refresh token has not been found")

		return
	}

	jwt, err := r.uc.RefreshAccessToken(ctx, refreshToken)

	if err == nil {
		ctx.JSON(http.StatusOK, jwt)
		return
	}

	errorResponse(ctx, http.StatusBadRequest, err.Error())
}

func (r *Router) LogoutUser(ctx *gin.Context) {
	domainName := r.uc.GetDomainName()
	ctx.SetCookie("access_token", "", -1, "/", domainName, false, true)
	ctx.SetCookie("refresh_token", "", -1, "/", domainName, false, true)
	ctx.SetCookie("logged_in", "", -1, "/", domainName, false, false)

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

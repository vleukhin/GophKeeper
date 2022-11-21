package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vleukhin/GophKeeper/internal/helpers/errs"
)

type loginPayload struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (r *Router) Register(ctx *gin.Context) {
	var payload *loginPayload

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	user, err := r.uc.SignUpUser(ctx, payload.Name, payload.Password)
	if err == nil {
		ctx.JSON(http.StatusCreated, user)

		return
	}

	if errors.Is(err, errs.ErrWrongEmail) || errors.Is(err, errs.ErrEmailAlreadyExists) {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

func (r *Router) LogIn(ctx *gin.Context) {
	var payload *loginPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	jwt, err := r.uc.SignInUser(ctx, payload.Name, payload.Password)

	if err == nil {
		ctx.SetCookie("access_token", jwt.AccessToken, jwt.AccessTokenMaxAge, "/", jwt.Domain, false, true)
		ctx.SetCookie("refresh_token", jwt.RefreshToken, jwt.RefreshTokenMaxAge, "/", jwt.Domain, false, true)
		ctx.SetCookie("logged_in", "true", jwt.AccessTokenMaxAge, "/", jwt.Domain, false, false)

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
		ctx.SetCookie("access_token", jwt.AccessToken, jwt.AccessTokenMaxAge, "/", jwt.Domain, false, true)
		ctx.SetCookie("logged_in", "true", jwt.AccessTokenMaxAge, "/", jwt.Domain, false, false)

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

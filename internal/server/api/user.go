package api

import (
	"net/http"

	"github.com/vleukhin/GophKeeper/internal/models"

	"github.com/gin-gonic/gin"

	"github.com/vleukhin/GophKeeper/internal/helpers/errs"
)

func (r *Router) getUserFromCtx(ctx *gin.Context) (user models.User, err error) {
	currentUser, ok := ctx.Get("currentUser")
	if !ok {
		err = errs.ErrUnexpectedError

		return
	}

	return currentUser.(models.User), nil
}

func (r *Router) UserInfo(ctx *gin.Context) {
	currentUser, err := r.getUserFromCtx(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, errs.ErrUnexpectedError.Error())

		return
	}

	ctx.JSON(http.StatusOK, currentUser)
}

package api

import (
	"net/http"

	"github.com/vleukhin/GophKeeper/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/vleukhin/GophKeeper/internal/helpers/errs"
)

func (r *Router) GetCreds(ctx *gin.Context) {
	currentUser, err := r.getUserFromCtx(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, errs.ErrUnexpectedError.Error())
	}

	userLogins, err := r.uc.GetCred(ctx, currentUser)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	if len(userLogins) == 0 {
		ctx.Status(http.StatusNoContent)

		return
	}

	ctx.JSON(http.StatusOK, userLogins)
}

func (r *Router) AddCred(ctx *gin.Context) {
	currentUser, err := r.getUserFromCtx(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, errs.ErrUnexpectedError.Error())
	}

	var payloadLogin *models.Cred

	if err := ctx.ShouldBindJSON(&payloadLogin); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	if err := r.uc.AddCred(ctx, payloadLogin, currentUser.ID); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	ctx.JSON(http.StatusAccepted, payloadLogin)
}

func (r *Router) DelCred(ctx *gin.Context) {
	loginUUID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	currentUser, err := r.getUserFromCtx(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, errs.ErrUnexpectedError.Error())
	}

	if err := r.uc.DelCred(ctx, loginUUID, currentUser.ID); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())

		return
	}

	ctx.Status(http.StatusAccepted)
}

func (r *Router) UpdateCred(ctx *gin.Context) {
	credID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	currentUser, err := r.getUserFromCtx(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, errs.ErrUnexpectedError.Error())

		return
	}

	var cred *models.Cred

	if err := ctx.ShouldBindJSON(&cred); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	cred.ID = credID

	if err := r.uc.UpdateCred(ctx, cred, currentUser.ID); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}

	ctx.Status(http.StatusAccepted)
}

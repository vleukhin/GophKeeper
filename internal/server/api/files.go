package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/vleukhin/GophKeeper/internal/helpers/errs"
	"github.com/vleukhin/GophKeeper/internal/models"
)

var errFileNameNotGiven = errors.New("binary name has not given")

func (r *Router) GetFiles(ctx *gin.Context) {
	currentUser, err := r.getUserFromCtx(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, errs.ErrUnexpectedError.Error())
	}

	userBinaries, err := r.uc.GetFiles(ctx, currentUser)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	if len(userBinaries) == 0 {
		ctx.Status(http.StatusNoContent)

		return
	}

	ctx.JSON(http.StatusOK, userBinaries)
}

func (r *Router) AddFile(ctx *gin.Context) {
	currentUser, err := r.getUserFromCtx(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, errs.ErrUnexpectedError.Error())

		return
	}
	var file models.File
	if ctx.Query("name") == "" {
		errorResponse(ctx, http.StatusBadRequest, errFileNameNotGiven.Error())

		return
	}
	file.Name = ctx.Query("name")

	uploadedFile, err := ctx.FormFile("uploadedFile")
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())

		return
	}
	file.FileName = uploadedFile.Filename
	if err = r.uc.AddFile(ctx, file, uploadedFile, currentUser.ID); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())

		return
	}
	ctx.JSON(http.StatusCreated, file)
}

func (r *Router) DownloadFile(ctx *gin.Context) {
	binaryUUID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}
	currentUser, err := r.getUserFromCtx(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, errs.ErrUnexpectedError.Error())

		return
	}

	filePath, err := r.uc.GetFile(ctx, currentUser, binaryUUID)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())

		return
	}

	ctx.Status(http.StatusOK)

	ctx.File(filePath)
}

func (r *Router) DelFile(ctx *gin.Context) {
	binaryUUID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())

		return
	}
	currentUser, err := r.getUserFromCtx(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, errs.ErrUnexpectedError.Error())

		return
	}

	if err := r.uc.DelFile(ctx, currentUser, binaryUUID); err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())

		return
	}

	ctx.Status(http.StatusAccepted)
}

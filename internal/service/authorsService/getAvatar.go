package authorsService

import (
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (as *authorsService) GetAvatar(ctx *gin.Context) {
	// Получаем параметр fileName из запроса
	fileName, existFileName := ctx.Params.Get("fileName")

	if !existFileName {
		logger.Log.Error(fmt.Sprintf("GetAvatar: Parameter fileName exist - %t", existFileName))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"Invalid author fileName parameter.",
			nil,
		)
		return
	}

	val, errGetAvatar := as.repository.GetAvatar(fileName)

	if errGetAvatar != nil {
		logger.Log.Error(fmt.Sprintf("GetAvatar: %v", errGetAvatar))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while getting the author avatar.",
			nil,
		)
		return
	}

	stat, errStat := val.Stat()
	if errStat != nil {
		logger.Log.Error(fmt.Sprintf("GetAvatar: %v", errStat))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while getting the author avatar.",
			nil,
		)
		return
	}

	defer val.Close()

	ctx.Status(http.StatusOK)
	ctx.Header("Content-Disposition", "attachment; filename="+stat.Key)
	ctx.Stream(func(w io.Writer) bool {
		_, errCopy := io.Copy(w, val)
		return errCopy == nil && false
	})
}

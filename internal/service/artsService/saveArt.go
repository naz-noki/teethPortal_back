package artsService

import (
	"MySotre/internal/service"
	"MySotre/pkg/logger"
	"MySotre/pkg/sendResponse"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create new art
// @Tags arts
// @Accept json
// @Produce json
// @Param file formData []file true "New files"
// @Param data formData service.SaveArtBody true "Art information"
// @Success 201 {object} sendResponse.Response
// @Failure 400 {object} sendResponse.Response
// @Failure 500 {object} sendResponse.Response
// @Router /api/arts [post]
func (as *artsService) SaveArt(ctx *gin.Context) {
	// Получаем файлы из запроса
	form, errForm := ctx.MultipartForm()
	if errForm != nil {
		logger.Log.Error(fmt.Sprintf("SaveArt: %v", errForm))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"An error occurred while retrieving data from the request body.",
			nil,
		)
		return
	}
	// Извлекаем массив файлов по ключу "file"
	files := form.File["file"]
	// Проверяем, что файлы найдены
	if len(files) == 0 {
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"No files were uploaded.",
			nil,
		)
		return
	}
	// Парсим тело запроса
	body := new(service.SaveArtBody)
	errUnmarshal := json.Unmarshal(
		[]byte(ctx.PostForm("data")),
		body,
	)

	if errUnmarshal != nil {
		logger.Log.Error(fmt.Sprintf("SaveArt: %v", errUnmarshal))
		sendResponse.Send(
			ctx,
			http.StatusBadRequest,
			"error",
			"An error occurred while retrieving data from the request body.",
			nil,
		)
		return
	}
	// Сохраняем запись
	artId, errSaveArt := as.repository.SaveArt(body.Title, body.Description, body.Content, body.Type, body.AuthorId)

	if errSaveArt != nil {
		logger.Log.Error(fmt.Sprintf("SaveArt: %v", errSaveArt))
		sendResponse.Send(
			ctx,
			http.StatusInternalServerError,
			"error",
			"An error occurred while saving data for art.",
			nil,
		)
		return
	}
	// Сохраняем файлы
	for i := 0; i < len(files); i++ {
		errSaveFile := as.repository.SaveFile(artId, files[i])

		if errSaveFile != nil {
			logger.Log.Error(fmt.Sprintf("SaveArt: %v", errSaveFile))
			sendResponse.Send(
				ctx,
				http.StatusInternalServerError,
				"error",
				"An error occurred while saving files for art.",
				nil,
			)
			return
		}
	}

	sendResponse.Send(
		ctx,
		http.StatusCreated,
		"success",
		"OK.",
		nil,
	)
}

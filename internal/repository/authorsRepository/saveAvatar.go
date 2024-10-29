package authorsRepository

import (
	"MySotre/pkg/minioDB"
	"mime/multipart"
)

func (ar *authorsRepository) SaveAvatar(
	fileHeader *multipart.FileHeader,
) (string, error) {
	file, errOpen := fileHeader.Open()
	if errOpen != nil {
		return "", errOpen
	}
	defer file.Close()

	avatarId, errUpload := minioDB.Client.Upload(ar.bucketName, fileHeader.Filename, fileHeader.Size, file)
	if errUpload != nil {
		return "", errUpload
	}

	return avatarId, nil
}

package categoriesRepository

import (
	"MySotre/pkg/minioDB"
	"mime/multipart"
)

func (t *categoriesRepository) SavePreviewFile(
	fileHeader *multipart.FileHeader,
) (string, error) {
	file, errOpen := fileHeader.Open()
	if errOpen != nil {
		return "", errOpen
	}
	defer file.Close()

	id, errUpload := minioDB.Client.Upload(t.bucketName, fileHeader.Filename, fileHeader.Size, file)
	if errUpload != nil {
		return "", errUpload
	}

	return id, nil
}

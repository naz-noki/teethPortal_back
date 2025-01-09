package categoriesRepository

import (
	"MySotre/pkg/minioDB"
	"mime/multipart"
)

func (t *categoriesRepository) UpdatePreviewFile(
	oldFileName string,
	fileHeader *multipart.FileHeader,
) (string, error) {
	errRemove := minioDB.Client.Remove(oldFileName, t.bucketName)

	if errRemove != nil {
		return "", errRemove
	}

	file, errOpen := fileHeader.Open()
	if errOpen != nil {
		return "", errOpen
	}
	defer file.Close()

	avatarId, errUpload := minioDB.Client.Upload(t.bucketName, fileHeader.Filename, fileHeader.Size, file)
	if errUpload != nil {
		return "", errUpload
	}

	return avatarId, nil
}

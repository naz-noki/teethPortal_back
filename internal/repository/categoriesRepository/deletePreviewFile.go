package categoriesRepository

import (
	"MySotre/pkg/minioDB"
)

func (t *categoriesRepository) DeletePreviewFile(fileName string) error {
	return minioDB.Client.Remove(fileName, t.bucketName)
}

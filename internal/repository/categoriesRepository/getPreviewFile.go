package categoriesRepository

import (
	"MySotre/pkg/minioDB"

	"github.com/minio/minio-go/v7"
)

func (t *categoriesRepository) GetPreviewFile(fileName string) (*minio.Object, error) {
	return minioDB.Client.Get(t.bucketName, fileName)
}

package artsRepository

import (
	"MySotre/pkg/minioDB"

	"github.com/minio/minio-go/v7"
)

func (ar *artsRepository) GetFile(fileName string) (*minio.Object, error) {
	return minioDB.Client.Get(ar.bucketName, fileName)
}

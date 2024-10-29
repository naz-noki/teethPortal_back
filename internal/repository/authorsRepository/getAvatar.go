package authorsRepository

import (
	"MySotre/pkg/minioDB"

	"github.com/minio/minio-go/v7"
)

func (ar *authorsRepository) GetAvatar(fileName string) (*minio.Object, error) {
	return minioDB.Client.Get(ar.bucketName, fileName)
}

package minioDB

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type client struct {
	minioClient *minio.Client
}

var (
	Client *client
)

func New(
	endpoint, accessKeyID, secretAccessKey string,
	useSSL bool,
) error {
	minioClient, errNew := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if errNew != nil {
		return errNew
	}

	Client = new(client)
	Client.minioClient = minioClient
	return nil
}

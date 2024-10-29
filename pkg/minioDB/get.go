package minioDB

import (
	"context"

	"github.com/minio/minio-go/v7"
)

func (c *client) Get(
	bucketName, fileName string,
) (*minio.Object, error) {
	return c.minioClient.GetObject(context.Background(), bucketName, fileName, minio.GetObjectOptions{})
}

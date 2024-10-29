package minioDB

import (
	"context"

	"github.com/minio/minio-go/v7"
)

func (c *client) Remove(
	fileName, bucketName string,
) error {
	return c.minioClient.RemoveObject(context.Background(), bucketName, fileName, minio.RemoveObjectOptions{})
}

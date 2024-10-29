package minioDB

import (
	"context"

	"github.com/minio/minio-go/v7"
)

func (c *client) createBucket(name string) error {
	return c.minioClient.MakeBucket(context.Background(), name, minio.MakeBucketOptions{})
}

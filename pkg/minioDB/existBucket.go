package minioDB

import "context"

func (c *client) existBucket(name string) (bool, error) {
	exist, errBucketExists := c.minioClient.BucketExists(context.Background(), name)

	if errBucketExists != nil {
		return false, errBucketExists
	}
	return exist, nil
}

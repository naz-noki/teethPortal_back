package minioDB

import (
	"context"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
)

func (c *client) Upload(
	bucketName, fileName string,
	fileSize int64, file multipart.File,
) (string, error) {
	exist, errExistBucket := c.existBucket(bucketName)

	if errExistBucket != nil {
		return "", errExistBucket
	}
	if !exist {
		if err := c.createBucket(bucketName); err != nil {
			return "", err
		}
	}

	name, errGenerateFileName := c.generateFileName(fileName)
	if errGenerateFileName != nil {
		return "", errGenerateFileName
	}

	_, errPutObject := c.minioClient.PutObject(
		context.Background(),
		bucketName,
		name,
		file,
		fileSize,
		minio.PutObjectOptions{},
	)

	if errPutObject != nil {
		return "", errPutObject
	}

	return name, nil
}

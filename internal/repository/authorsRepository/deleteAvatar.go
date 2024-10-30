package authorsRepository

import "MySotre/pkg/minioDB"

func (ar *authorsRepository) DeleteAvatar(fileName string) error {
	return minioDB.Client.Remove(fileName, ar.bucketName)
}

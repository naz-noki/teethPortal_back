package imagesRepository

import "os"

func (i *imagesRepository) DeleteImage(imagePath string) error {
	return os.Remove(imagePath)
}

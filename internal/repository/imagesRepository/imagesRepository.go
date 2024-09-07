package imagesRepository

type imagesRepository struct{}

func NewImagesRepository() *imagesRepository {
	ir := imagesRepository{}

	return &ir
}

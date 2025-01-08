package categoriesRepository

type categoriesRepository struct {
	bucketName string
}

func New() *categoriesRepository {
	return &categoriesRepository{
		bucketName: "categories-preview",
	}
}

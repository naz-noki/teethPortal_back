package artsRepository

type artsRepository struct {
	bucketName string
}

func New() *artsRepository {
	return &artsRepository{
		bucketName: "authors-arts",
	}
}

package authorsRepository

type authorsRepository struct {
	bucketName string
}

func New() *authorsRepository {
	ar := new(authorsRepository)

	ar.bucketName = "authors-avatars"

	return ar
}

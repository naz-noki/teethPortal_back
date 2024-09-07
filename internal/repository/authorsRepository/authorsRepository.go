package authorsRepository

type authorsRepository struct {
}

func NewAuthorsRepository() *authorsRepository {
	ar := authorsRepository{}

	return &ar
}

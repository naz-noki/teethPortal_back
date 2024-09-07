package repository

const (
	TableAuthors = "authors"
)

// MODELS

type User struct {
	Id       int
	Login    string
	Password string
	IsAdmin  bool
}

type Image struct {
	Id          int
	Path        string
	Title       string
	Description string
	UserId      int
	AuthorId    int
	CreatedAt   string
}

type Author struct {
	Id          int
	Name        string
	Description string
	UserId      int
	Avatar      string
}

package repository

// MODELS

type User struct {
	Id       int
	Login    string
	Password string
	IsAdmin  bool
}

type Author struct {
	Id          int
	Name        string
	Description string
	UserId      int
	AvatarId    string
}

type Art struct {
	Id          int
	Title       string
	Description string
	Content     string
	AuthorId    int
	Type        string // enum (painting or photo or product or text)
}

type ArtFile struct {
	Id     int
	ArtId  int
	FileId string
}

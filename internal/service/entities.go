package service

import (
	"MySotre/internal/repository"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
	"github.com/naz-noki/teethPortal_proto/gen/go/sso/authApi"
	"github.com/naz-noki/teethPortal_proto/gen/go/sso/tokensApi"
)

// ----------------------------------------
// SERVICES
// ----------------------------------------

type SsoService interface {
	authApi.AuthServer
	tokensApi.TokensServer
}

// ----------------------------------------
// REPOSITORIES
// ----------------------------------------

type SsoRepository interface {
	GetUserByLogin(login string) (*repository.User, error)
	SetUser(login, password string, isAdmin bool) error
	SetRefreshToken(userId int, token string) error
	GetRefreshTokenByUserId(userId int) (string, error)
	GetUserById(id int) (*repository.User, error)
}

type AuthRepository interface {
	GetUserIdByLogin(login string) (int, error)
}

type AuthorsRepository interface {
	SaveAuthor(
		name, description string,
		userId int,
		avatarId string,
	) (int, error)
	SaveAvatar(
		fileHeader *multipart.FileHeader,
	) (string, error)
	GetAuthorById(id int) (*repository.Author, error)
	GetAvatar(fileName string) (*minio.Object, error)
	GetAllAuthors() ([]*repository.Author, error)
	UpdateAuthor(
		name, description string,
		authorId, userId int,
	) error
	UpdateAvatar(
		oldFileName string,
		fileHeader *multipart.FileHeader,
	) (string, error)
	UpdateAvatarId(
		authorId int,
		avatarId string,
	) error
	DeleteAuthor(authorId int) error
	DeleteAvatar(fileName string) error
	GetAvatarId(authorId int) (string, error)
}

type ArtsRepository interface {
	SaveArt(
		title, description,
		content, artType string,
		authorId int,
	) (int, error)
	SaveFile(
		artId int,
		fileHeader *multipart.FileHeader,
	) error
	GetFileIds(artId int) ([]string, error)
	GetArts() ([]*repository.Art, error)
	GetArtById(id int) (*repository.Art, error)
	GetAuthorArts(authorId int) ([]*repository.Art, error)
	GetFile(fileName string) (*minio.Object, error)
	UpdateArt(
		title, description,
		content, artType string,
		artId, authorId int,
	) error
	UpdateFile(
		artId int,
		oldFileName string,
		fileHeader *multipart.FileHeader,
	) error
	DeleteFile(fileName string) error
	DeleteArt(artId int) error
}

// ----------------------------------------
// PAYLOAD FOR TOKENS
// ----------------------------------------

type UserPayload struct {
	UserId      int  `json:"userId"`
	UserIsAdmin bool `json:"userIsAdmin"`
}

// ----------------------------------------
// REQUEST BODIES
// ----------------------------------------

type RegistrationBody struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"isAdmin"`
}

type AuthorizationBody struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UpdateTokensBody struct {
	Login string `json:"login"`
}

type SaveAuthorBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Login       string `json:"login"`
}

type SaveArtBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	AuthorId    int    `json:"authorId"`
	Type        string `json:"type"`
}

// ----------------------------------------
// RESPONSE BODIES
// ----------------------------------------

type AuthorizationResponse struct {
	AccessToken string `json:"accessToken"`
}

type UpdateTokensResponse struct {
	AccessToken string `json:"accessToken"`
}

type GetAuthorByIdResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AvatarId    string `json:"avatarId"`
}

type GetArtResponse struct {
	Id          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Content     string   `json:"content"`
	AuthorId    int      `json:"author_id"`
	Type        string   `json:"type"`
	Files       []string `json:"files"`
}

// ----------------------------------------
// MODELS
// ----------------------------------------

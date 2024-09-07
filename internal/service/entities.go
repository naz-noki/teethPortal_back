package service

import (
	"MySotre/internal/repository"

	"github.com/gin-gonic/gin"
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

type UsersService interface {
	Registartion(ctx *gin.Context)
	Authorization(ctx *gin.Context)
}

type TokensService interface {
	UpdateTokens(ctx *gin.Context)
}

type ImagesService interface {
	AddNewImage(ctx *gin.Context)
	GetImage(ctx *gin.Context)
	GetAllImages(ctx *gin.Context)
	GetAuthorImages(ctx *gin.Context)
	PutDataForImage(ctx *gin.Context)
	DeleteImage(ctx *gin.Context)
}

type AuthorsService interface {
	AddNewAuthor(ctx *gin.Context)
	GetAllAuthors(ctx *gin.Context)
	GetAuthor(ctx *gin.Context)
	GetAvatarForAuthor(ctx *gin.Context)
	PutAuthor(ctx *gin.Context)
	PutAuthorAvatar(ctx *gin.Context)
	DeleteAuthor(ctx *gin.Context)
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

type TokensRpository interface {
	GetUserIdByLogin(login string) (int, error)
}

type ImagesRepository interface {
	GetUserIdByLogin(login string) (int, error)
	SaveImage(
		dirPath, fileName string,
		userId int,
		data *[]byte,
	) (string, error)
	SaveImageData(
		userId, authorId int,
		path, title,
		description, createdAt string,
	) error
	GetImagesFromDB() (*[]repository.Image, error)
	GetImagesFromDBbyAuthorId(authorId int) (*[]repository.Image, error)
	GetImagePathById(id int) (string, error)
	UpdateImageData(
		imageId int,
		newTitle, newDescription string,
	) error
	DeleteImage(imagePath string) error
	DeleteImageData(imageId int, imagePath string) error
	GetAuthorIdByName(name string) (int, error)
}

type AuthorsRepository interface {
	GetUserIdByLogin(login string) (int, error)
	SaveAuthor(
		name, description, avatarPath string,
		userId int,
	) error
	GetAuthors() (*[]repository.Author, error)
	GetAuthorById(id int) (*repository.Author, error)
	UpdateAuthor(
		name, description string,
		id int,
	) error
	DeleteAuthor(id int) error
	SaveAuthorAvatar(
		authorName, dirPath,
		avatarName string, avatarData []byte,
	) (string, error)
	DeleteAuthorAvatar(path string) error
	GetAvatarPathByAvtorName(name string) (string, error)
	UpdatePathToAuthorAvatar(
		avatarPath, authorName string,
	) error
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

type RegistartionBody struct {
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

type AddNewImageBody struct {
	Login       string `json:"login"`
	AuthorName  string `json:"authorName"`
	FileName    string `json:"fileName"`
	File        []byte `json:"file"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
}

type PutDataForImageBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type AddNewAuthorBody struct {
	Login       string `json:"login"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AvatarName  string `json:"avatarName"`
	AvatarData  []byte `json:"avatarData"`
}

type PutAuthorBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type PutAuthorAvatarBody struct {
	Name       string `json:"name"`
	AvatarName string `json:"avatarName"`
	AvatarData []byte `json:"avatarData"`
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

type GetAllImagesResponse struct {
	Images *[]ImageForUser `json:"images"`
}

type GetAllAuthorsResponse struct {
	Authors *[]AuthorForUser `json:"authors"`
}

type GetAuthorResponse struct {
	Author *AuthorForUser `json:"author"`
}

// ----------------------------------------
// MODELS
// ----------------------------------------

type ImageForUser struct {
	Id          int    `json:"imageId"`
	Title       string `json:"imageTitle"`
	Description string `json:"imageDescription"`
	CreatedAt   string `json:"createdAt"`
}

type AuthorForUser struct {
	Id          int    `json:"authorId"`
	Name        string `json:"authorName"`
	Description string `json:"authorDescription"`
}

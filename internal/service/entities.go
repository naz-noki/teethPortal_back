package service

import (
	"MySotre/internal/repository"
	"mime/multipart"
	"time"

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

type GlobalSearchRepository interface {
	GetArts(searchText string) ([]*repository.Art, error)
	GetAuthors(searchText string) ([]*repository.Author, error)
	GetCategories(searchText string) ([]*repository.Category, error)
}

type CategoriesRepository interface {
	CreateCategory(
		name, description,
		previewFileId string,
	) error
	SavePreviewFile(
		fileHeader *multipart.FileHeader,
	) (string, error)
	CheckExistCategory(id int) (bool, error)
	AddAuthor(authorId, categoryId int) error
	DeleteAuthor(authorId, categoryId int) error
	DeleteCategory(id int) error
	GetCategoryById(id int) (*repository.Category, error)
	UpdateCategory(
		id int,
		name, description string,
	) error
	GetAllCategoriesIds() ([]int, error)
	GetPreviewFile(fileName string) (*minio.Object, error)
	DeletePreviewFile(fileName string) error
	UpdatePreviewFile(
		oldFileName string,
		fileHeader *multipart.FileHeader,
	) (string, error)
	UpdatePreviewFileId(
		id int,
		newPreviewFileId string,
	) error
}

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
	GetAllAuthors(limit, offset int) ([]*repository.Author, error)
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
	CheckExistAuthor(id int) (bool, error)
	CountAllRecords() (int, error)
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
	GetArts(limit, offset int, artType string) ([]*repository.Art, error)
	GetAuthorArts(authorId, limit, offset int, artType string) ([]*repository.Art, error)
	GetArtById(id int) (*repository.Art, error)
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
	CheckExistArt(id int) (bool, error)
	CountAllRecords() (int, error)
	CountAuthorArts(authorId int) (int, error)
}

type AuctionRepository interface {
	AddLot(
		startTime, endTime time.Duration,
		price, artId, authorId int,
		paid, sent bool,
	) error
	UpdateLot(
		auctionId int,
		startTime, endTime time.Duration,
		price, artId, authorId, userId int,
		paid, sent bool,
	) error
	PlaceBet(
		auctionId, price, userId int,
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

type CreateCategoryBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type AddAuthorsBody struct {
	AuthorsIds []int `json:"authorsIds"`
}

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

type SearchResponse struct {
	Arts       []*repository.Art      `json:"arts"`
	Authors    []*repository.Author   `json:"authors"`
	Categories []*repository.Category `json:"categories"`
}

// ----------------------------------------
// MODELS
// ----------------------------------------

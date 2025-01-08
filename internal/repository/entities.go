package repository

import "time"

const (
	MaxLimit = 9223372036854775807
)

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

type Auction struct {
	Id        int
	StartTime time.Duration
	EndTime   time.Duration
	Price     int
	Lot       int // art_id
	Seller    int // author_id
	Buyer     int // user_id
	Paid      bool
	Sent      bool
}

type Category struct {
	Id          int
	Name        string
	Description string
	AuthorsIds  []int
}

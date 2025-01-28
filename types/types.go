package types

import (
	"mime/multipart"
	"time"
)

type AuthStore interface {
	CreateUser(user *User) error
}

type LikedSongsStore interface {
	AddToLikedSongs(userID int, songID int) error
	GetUserLikedSongs(userID int) ([]LikedSongs, error)
	RemoveFromLikedSongs(likedSongID int) error
}

type SongStore interface {
	GetSongs() ([]Song, error)
	GetSongById(songID int) ([]Song, error)
	UploadSong(file multipart.File, header *multipart.FileHeader, artist, name string, userID int) error
	DeleteSong(songID int) error
}

type User struct {
	RawData           map[string]interface{}
	Provider          string
	Email             string
	Name              string
	FirstName         string
	LastName          string
	NickName          string
	Description       string
	UserID            string
	AvatarURL         string
	Location          string
	AccessToken       string
	AccessTokenSecret string
	RefreshToken      string
	ExpiresAt         time.Time
	IDToken           string
}

type Song struct {
	ID        int
	Name      string
	Artist    string
	FileName  string
	UserID    int
	CreatedAt time.Time
}

type LikedSongs struct {
	ID      int       `json:"id"`
	UserID  int       `json:"user_id"`
	SongID  int       `json:"song_id"`
	LikedAt time.Time `json:"liked_at"`
}

type AddLikeRequestPayload struct {
	UserID int `json:"user_id"`
	SongID int `json:"song_id"`
}

type RemoveLikeRequestPayload struct {
	LikedID int `json:"id"`
}

type GetUserLikesRequestPayload struct {
	UserID int `json:"user_id"`
}

type DeleteSongRequestPayload struct {
	SongID int `json:"id"`
}

package liked_songs

import (
	"database/sql"
	"fmt"
	"playr-server/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) AddToLikedSongs(userID int, songID int) error {
	insertQuery := `
		INSERT INTO liked_songs (user_id, song_id)
		VALUES ($1, $2)
	`
	_, err := s.db.Exec(insertQuery, userID, songID)
	if err != nil {
		return fmt.Errorf("failed to add song to Liked Songs: %w", err)
	}
	return nil
}

func (s *Store) GetUserLikedSongs(userID int) ([]types.LikedSongs, error) {
	query := `
        SELECT id, user_id, song_id, liked_at FROM liked_songs
        WHERE user_id = $1
    `
	rows, err := s.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get the user's liked songs: %w", err)
	}
	defer rows.Close()

	var likedSongs []types.LikedSongs
	for rows.Next() {
		song, err := ScanRowIntoLikedSongs(rows)
		if err != nil {
			return nil, err
		}
		likedSongs = append(likedSongs, *song)
	}

	if len(likedSongs) == 0 {
		return nil, fmt.Errorf("no liked songs found for user %d", userID)
	}

	return likedSongs, nil
}

func (s *Store) RemoveFromLikedSongs(likedSongID int) error {
	query := `
        DELETE FROM liked_songs
        WHERE id = $1;
    `
	_, err := s.db.Exec(query, likedSongID)
	if err != nil {
		return fmt.Errorf("failed to remove liked song: %w", err)
	}
	return nil
}

func ScanRowIntoLikedSongs(rows *sql.Rows) (*types.LikedSongs, error) {
	likedSong := new(types.LikedSongs)

	err := rows.Scan(
		&likedSong.ID,
		&likedSong.UserID,
		&likedSong.SongID,
		&likedSong.LikedAt,
	)

	if err != nil {
		return nil, err
	}

	return likedSong, nil
}

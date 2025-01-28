package songs

import (
	"database/sql"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"playr-server/types"
	"playr-server/utils"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetSongs() ([]types.Song, error) {
	query := `
		SELECT * FROM songs
	`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all songs: %w", err)
	}
	defer rows.Close()

	var songs []types.Song
	for rows.Next() {
		song, err := ScanRowIntoSongs(rows)
		if err != nil {
			return nil, err
		}
		songs = append(songs, *song)
	}

	if len(songs) == 0 {
		return nil, fmt.Errorf("no songs found :(")
	}

	return songs, nil
}

func (s *Store) GetSongById(songID int) ([]types.Song, error) {
	query := `
		SELECT * FROM songs
		WHERE id = $1;
	`
	rows, err := s.db.Query(query, songID)
	if err != nil {
		return nil, fmt.Errorf("failed to get song by ID: %w", err)
	}
	defer rows.Close()

	var songs []types.Song
	for rows.Next() {
		song, err := ScanRowIntoSongs(rows)
		if err != nil {
			return nil, err
		}
		songs = append(songs, *song)
	}

	if len(songs) == 0 {
		return nil, fmt.Errorf("no song found with ID %d", songID)
	}

	return songs, nil
}

func (s *Store) UploadSong(file multipart.File, header *multipart.FileHeader, artist, name string, userID int) error {
	randomID := utils.GenerateRandomID()

	fileName := fmt.Sprintf("%s_%s_%d%s", artist, name, randomID, filepath.Ext(header.Filename))
	filePath := filepath.Join("uploads", fileName)

	query := `
		INSERT INTO songs (name, artist, file_name, user_id)
		VALUES ($1, $2, $3, $4)
	`
	_, err := s.db.Exec(query, name, artist, fileName, userID)
	if err != nil {
		return fmt.Errorf("failed to insert song into database: %w", err)
	}

	outFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	if err != nil {
		return fmt.Errorf("failed to save file: %w", err)
	}

	return nil
}

func (s *Store) DeleteSong(songID int) error {
	var fileName string
	query := `
        SELECT file_name FROM songs
        WHERE id = $1;
    `
	err := s.db.QueryRow(query, songID).Scan(&fileName)
	if err != nil {
		return fmt.Errorf("failed to retrieve song file name: %w", err)
	}

	filePath := filepath.Join("uploads", fileName)

	query = `
        DELETE FROM songs
        WHERE id = $1;
    `
	_, err = s.db.Exec(query, songID)
	if err != nil {
		return fmt.Errorf("failed to delete song from database: %w", err)
	}

	err = os.Remove(filePath)
	if err != nil {
		return fmt.Errorf("failed to delete song file: %w", err)
	}

	return nil
}

func ScanRowIntoSongs(rows *sql.Rows) (*types.Song, error) {
	song := new(types.Song)

	err := rows.Scan(
		&song.ID,
		&song.Name,
		&song.Artist,
		&song.FileName,
		&song.CreatedAt,
		&song.UserID,
	)

	if err != nil {
		return nil, err
	}

	return song, nil
}

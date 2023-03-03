package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"gocloudcamp/internal/playlist"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Database struct {
	config   *Config
	db       *sql.DB
	playlist *playlist.Playlist
}

// Creating new database
func NewDB(pl *playlist.Playlist) *Database {
	return &Database{
		config:   NewConfig(),
		playlist: pl,
	}
}

// Open database and creating table if table not exists
func (d *Database) Open() error {
	db, err := sql.Open("pgx", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		d.config.Host, d.config.Port, d.config.Username, d.config.Name, d.config.Password))
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	d.db = db

	query := `CREATE TABLE IF NOT EXISTS playlist
				(
				id integer NOT NULL GENERATED ALWAYS AS IDENTITY,
				name text NOT NULL,
				duration integer NOT NULL,
				PRIMARY KEY (id));`

	_, err = d.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

// Closing database
func (d *Database) Close() {
	d.db.Close()
}

// Getting all songs from database
func (d *Database) GetAllSongs() error {
	rows, err := d.db.Query("SELECT name, duration FROM playlist")

	defer rows.Close()
	if err != nil {
		log.Println(err)
		return err
	}

	for rows.Next() {
		var name string
		var duration int
		err = rows.Scan(&name, &duration)
		if err != nil {
			log.Println(err)
			return err
		}
		d.playlist.AddSong(name, int(duration))
	}

	err = rows.Err()
	if err != nil {
		return err
	}

	return nil
}

// Saving new playlist in database after playlist is stopped
func (d *Database) SetPlaylist() error {
	_, err := d.db.Exec("TRUNCATE TABLE playlist")
	if err != nil {
		return err
	}

	playlist := d.playlist.SetMemoryPlaylist()
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	stmt, err := tx.PrepareContext(context.Background(),
		`INSERT INTO playlist (name, duration) VALUES ($1, $2)`)
	if err != nil {
		return err
	}

	defer stmt.Close()
	for _, song := range playlist {
		_, err = stmt.ExecContext(context.Background(), song.Name, song.Duration)
		if err != nil {
			return err
		}
	}
	return tx.Commit()

}

type Song struct {
	Name     string
	Duration int
}

// Getting song by name from database
func (d *Database) GetSongByName(name string) (Song, bool) {
	var song Song
	err := d.db.QueryRow("SELECT name, duration FROM playlist WHERE name = $1", name).Scan(&song.Name, &song.Duration)
	if err != nil {
		return Song{}, false
	}
	return song, true

}

// Updating song in database
func (d *Database) UpdateSong(oldName, newName string, duration int) error {
	var name string
	err := d.db.QueryRow("SELECT name FROM playlist WHERE name = $1", oldName).Scan(&name)
	if err != nil {
		return errors.New("Song not found")
	}
	_, err = d.db.Exec("UPDATE playlist SET name = $1, duration = $2 WHERE name = $3", newName, duration, oldName)
	if err != nil {
		return err
	}
	return nil
}

// Deleting song from database
func (d *Database) DeleteSong(name string) error {
	var testName string
	err := d.db.QueryRow("SELECT name FROM playlist WHERE name = $1", name).Scan(&testName)
	if err != nil {
		return errors.New("Song not found")
	}
	_, err = d.db.Exec("DELETE FROM playlist WHERE name = $1", name)
	if err != nil {
		return err
	}
	return nil
}

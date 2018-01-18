package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Datastore interface {
	AllShows() ([]*Show, error)
	AddShow(Name string, Season int32, Episode int32) error
	UpdateShow(ID int32, Name string, Season int32, Episode int32) error
	LoadShow(ID int32) (*Show, error)
	DeleteShow(ID int32) error
  NextEpisode(ID int32) error
  PreviousEpisode(ID int32) error
}

type DB struct {
	*sql.DB
}

func NewDB(dataSourceName string) (*DB, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

package storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Storage struct {
	*sqlx.DB
	Users UsersStorage
}

func Open(url string) (*Storage, error) {
	db, err := sqlx.Open("postgres", url)

	if err != nil {
		return nil, err
	}

	return &Storage{
		DB:    db,
		Users: &Users{},
	}, nil

}

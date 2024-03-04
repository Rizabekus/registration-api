package storage

import "database/sql"

type PersonDB struct {
	DB *sql.DB
}

func CreateUserStorage(db *sql.DB) *PersonDB {
	return &PersonDB{DB: db}
}

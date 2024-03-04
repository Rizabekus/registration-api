package storage

import (
	"database/sql"

	"github.com/Rizabekus/registration-api/internal/models"
)

type Storage struct {
	UserStorage models.UserStorage
}

func StorageInstance(db *sql.DB) *Storage {
	return &Storage{UserStorage: CreateUserStorage(db)}
}

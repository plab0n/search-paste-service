package storage

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Initializes the postgres driver
	"github.com/plab0n/search-paste/internal/model"
)

type StorageInterface interface {
	AddPaste(ctx context.Context, book model.AddPasteRequest) (int, error)
	GetPaste(ctx context.Context, id int) (model.Book, error)
	UpdatePaste(ctx context.Context, book model.UpdateBookRequest) (int, error)
	DeletePaste(ctx context.Context, id int) error
	VerifyPasteExists(ctx context.Context, id int) (bool, error)
}

// Storage contains an SQL db. Storage implements the StorageInterface.
type Storage struct {
	db *sqlx.DB
}

func (s *Storage) Close() error {
	if err := s.db.Close(); err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetDB() *sqlx.DB {
	return s.db
}

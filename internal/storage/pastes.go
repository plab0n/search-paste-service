package storage

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/plab0n/search-paste/internal/model"
)

func (s *Storage) AddPaste(ctx context.Context, pasteRequest model.AddPasteRequest) (int, error) {
	var id int
	err := s.db.Get(&id, `INSERT INTO paste(title, text,created_at, updated_at)
			VALUES($1,$2,$3,$4) RETURNING id`, pasteRequest.Title, pasteRequest.Text, time.Now().UTC(), time.Now().UTC())

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Storage) GetPaste(ctx context.Context, id int) (model.Book, error) {
	var book model.Book

	err := s.db.Get(&book, `Select * from books where id=$1`, id)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (s *Storage) UpdatePaste(ctx context.Context, book model.UpdateBookRequest) (int, error) {
	var columns []string
	var argCount = 1
	var args []interface{}

	if book.Title != "" {
		columns = append(columns, fmt.Sprintf("title=$%d", argCount))
		args = append(args, book.Title)
		argCount++
	}

	if book.Author != "" {
		columns = append(columns, fmt.Sprintf("author=$%d", argCount))
		args = append(args, book.Author)
		argCount++
	}

	if book.CoverURL != "" {
		columns = append(columns, fmt.Sprintf("cover_url=$%d", argCount))
		args = append(args, book.CoverURL)
		argCount++
	}

	if book.PostURL != "" {
		columns = append(columns, fmt.Sprintf("post_url=$%d", argCount))
		args = append(args, book.PostURL)
		argCount++
	}

	columns = append(columns, fmt.Sprintf("updated_at=$%d", argCount))
	args = append(args, time.Now().UTC())
	argCount++

	if len(columns) == 0 {
		return 0, errors.New("No fields to update")
	}

	args = append(args, book.ID)

	query := fmt.Sprintf(`UPDATE books SET %s WHERE id=$%d RETURNING id`, strings.Join(columns, ", "), argCount)

	var id int
	err := s.db.Get(&id, query, args...)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Storage) DeletePaste(ctx context.Context, id int) error {
	_, err := s.db.Exec(`DELETE FROM books WHERE id=$1`, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) VerifyPasteExists(ctx context.Context, id int) (bool, error) {
	var exists bool
	err := s.db.Get(&exists, `SELECT EXISTS(SELECT 1 from books where id=$1)`, id)
	if err != nil {
		return false, err
	}

	return exists, nil
}

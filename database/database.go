package database

import (
	"context"
	"docker-golang/book"
)

type Database interface {
	Create(ctx context.Context, b book.Book) (book.Book, error)
}

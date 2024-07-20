package book

import "time"

type Book struct {
	ID         int
	Name       string
	AuthorName string
	CreateTime time.Time
}

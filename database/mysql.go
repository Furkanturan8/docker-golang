package database

import (
	"context"
	"database/sql"
	"docker-golang/book"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlDB struct {
	db *sql.DB
}

type MysqlConfig struct {
	Username string
	Password string
	DbName   string
	Port     uint
	Host     string
}

func NewMysqlDB(conf MysqlConfig) (MysqlDB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.Username, conf.Password, conf.Host, conf.Port, conf.DbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return MysqlDB{}, fmt.Errorf("impossible to open SQL connection: %w", err)
	}
	err = db.Ping()
	if err != nil {
		return MysqlDB{}, fmt.Errorf("impossible to ping db: %w", err)
	}
	return MysqlDB{db: db}, nil
}

func (s MysqlDB) Create(ctx context.Context, b book.Book) (book.Book, error) {
	query := "INSERT INTO `book` (`create_time`, `name`, `author_name`) VALUES (?, ?, ?)"
	insertResult, err := s.db.ExecContext(ctx, query, b.CreateTime, b.Name, b.AuthorName)
	if err != nil {
		return b, fmt.Errorf("error while insert: %w", err)
	}
	id, err := insertResult.LastInsertId()
	if err != nil {
		return b, fmt.Errorf("error while get last insert id: %w", err)
	}
	b.ID = int(id)

	return b, nil
}

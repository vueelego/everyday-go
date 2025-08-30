package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	db *sqlx.DB
}

func NewStore(dsn string) Store {
	db, err := sqlx.Connect("sqlite3", dsn)
	if err != nil {
		panic(err)
	}
	return Store{db: db}
}

type Backup struct {
	ID         int    `json:"id" db:"id"`
	Api        string `json:"api" db:"api"`
	Query      string `json:"query" db:"query"`
	Param      string `json:"param" db:"param"`
	JSONString string `json:"jsonString" db:"jsonString"`
	CreatedAt  string `json:"createdAt" db:"createdAt"`
	UpdatedAt  string `json:"updatedAt" db:"updatedAt"`
}

func (store *Store) CreateTable() {
	var schema = `
		CREATE TABLE IF NOT EXISTS backup (
			id integer not null primary key,
			api text not null,
			query text not null,
			param text not null,
			jsonString text not null,
			created_at text not null default (datetime('now', 'localtime')),
			updated_at text not null default (datetime('now', 'localtime')),
		);
	`
	// 执行错误会panic
	_ = store.db.MustExec(schema)
}

// fuzzyGet 匹配 api 查找
func (store *Store) fuzzyGetByApi() {}

// fuzzyGetByQuery 匹配 api query 查找
func (store *Store) fuzzyGetByQuery() {}

// fuzzyGetByParam 匹配 api param 查找
func (store *Store) fuzzyGetByParam() {}

// PreciseGet 匹配 api query param 查找
func (store *Store) PreciseGet() {}

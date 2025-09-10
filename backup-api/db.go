package main

import (
	"fmt"
	"strings"

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
	ID        int    `json:"id" db:"id"`
	Api       string `json:"api" db:"api"`
	Query     string `json:"query" db:"query"`
	Body      string `json:"body" db:"body"`
	Response  string `json:"response" db:"response"`
	CreatedAt string `json:"createdAt,omitempty" db:"createdAt"`
	UpdatedAt string `json:"updatedAt,omitempty" db:"updatedAt"`
}

func (store *Store) CreateTable() {
	var schema = `
		CREATE TABLE IF NOT EXISTS backup (
			id integer not null primary key autoincrement,
			api text not null,
			query text not null,
			body text not null,
			response text not null,
			created_at text not null default (datetime('now', 'localtime')),
			updated_at text not null default (datetime('now', 'localtime'))
		);
	`
	// 执行错误会panic
	_ = store.db.MustExec(schema)
}

// QueryBodys 查询参数结构体,如果字段值是nil则不参与where查询条件
type QueryBodys struct {
	Api   *string `db:"api"`
	Query *string `db:"query"`
	Body  *string `db:"body"`
}

// Validate 做一些必要的验证
func (q *QueryBodys) Validate() error {
	if q.Api == nil || (*q.Api) == "" {
		return fmt.Errorf("Api 查询字段必填")
	}
	return nil
}

// BuildDyWhere 构建动态where查询条件
func (q *QueryBodys) BuildDyWhere() (string, map[string]interface{}) {
	var conditions []string
	whereBodys := make(map[string]interface{})

	if q.Api != nil && *q.Api != "" {
		conditions = append(conditions, "api = :api") // 命名参数绑定
		whereBodys["api"] = *q.Api                    // 取值
	}

	if q.Query != nil && *q.Query != "" {
		conditions = append(conditions, "query = :query") // 命名参数绑定
		whereBodys["query"] = *q.Query                    // 取值
	}

	if q.Body != nil && *q.Body != "" {
		conditions = append(conditions, "body = :body") // 命名参数绑定
		whereBodys["body"] = *q.Body                    // 取值
	}

	return strings.Join(conditions, " AND "), whereBodys
}

// Save 保存数据
func (store *Store) Save(backup *Backup) error {
	sql := `insert into backup(api, query, body, response) values(?, ?, ?, ?)`
	result, err := store.db.Exec(sql, backup.Api, backup.Query, backup.Body, backup.Response)
	if err != nil {
		return err
	}
	newID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	if newID <= 0 {
		return fmt.Errorf("添加失败，lastInsertId: %d", newID)
	}
	return nil
}

// Get 动态条件查询
func (store *Store) Get(queries *QueryBodys) (backup Backup, err error) {
	err = queries.Validate()
	if err != nil {
		return backup, err
	}
	whereSql, whereArgs := queries.BuildDyWhere()
	sql := `select id, api, query, body, response from backup `
	if whereSql != "" {
		sql += ` where ` + whereSql
	}

	sql += " limit 1 "

	fmt.Println("query sql: ", sql)
	stmt, err := store.db.PrepareNamed(sql)
	if err != nil {
		return backup, err
	}
	defer stmt.Close()

	err = stmt.Get(&backup, whereArgs)
	if err != nil {
		return backup, err
	}

	return backup, nil
}

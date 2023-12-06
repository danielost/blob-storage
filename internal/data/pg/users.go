package pg

import (
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/dl7850949/blob-storage/internal/data"
)

const usersTableName = "users"

type usersQ struct {
	db  *pgdb.DB
	sql sq.SelectBuilder
}

func NewUsersQ(db *pgdb.DB) data.UsersQ {
	return &usersQ{
		db:  db.Clone(),
		sql: sq.Select("n.*").From(fmt.Sprintf("%s as n", usersTableName)),
	}
}

func (q *usersQ) New() data.UsersQ {
	return NewUsersQ(q.db)
}

func (q *usersQ) Get() (*data.User, error) {
	var result data.User
	err := q.db.Get(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q *usersQ) Select() ([]data.User, error) {
	var result []data.User
	err := q.db.Select(&result, q.sql)
	return result, err
}

func (q *usersQ) Transaction(fn func(q data.UsersQ) error) error {
	return q.db.Transaction(func() error {
		return fn(q)
	})
}

func (q *usersQ) Insert(value data.User) (*data.User, error) {
	value.CreatedAt = time.Now().UTC()
	clauses := structs.Map(value)
	clauses["login"] = value.Login
	clauses["password"] = value.Password

	var result data.User
	stmt := sq.Insert(usersTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)

	return &result, err
}

func (q *usersQ) FilterByID(ids ...int64) data.UsersQ {
	q.sql = q.sql.Where(sq.Eq{"n.id": ids})
	return q
}

func (q *usersQ) FilterByLogin(logins ...string) data.UsersQ {
	q.sql = q.sql.Where(sq.Eq{"n.login": logins})
	return q
}

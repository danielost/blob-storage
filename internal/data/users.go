package data

import (
	"time"
)

type UsersQ interface {
	New() UsersQ

	Get() (*User, error)

	Select() ([]User, error)

	Transaction(fn func(q UsersQ) error) error

	Insert(data User) (*User, error)

	FilterByID(id ...int64) UsersQ

	FilterByLogin(login ...string) UsersQ
}

type User struct {
	ID        int64     `db:"id" structs:"-"`
	CreatedAt time.Time `db:"created_at" structs:"created_at"`
	Login     string    `db:"login" structs:"-"`
	Password  string    `db:"password" structs:"-"`
}

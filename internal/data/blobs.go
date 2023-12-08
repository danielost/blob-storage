package data

import (
	"encoding/json"
	"time"

	"gitlab.com/distributed_lab/kit/pgdb"
)

type BlobsQ interface {
	New() BlobsQ

	Get() (*Blob, error)

	Select() ([]Blob, error)

	Delete(id int64) error

	Transaction(fn func(q BlobsQ) error) error

	Insert(data Blob) (*Blob, error)

	Page(pageParams pgdb.OffsetPageParams) BlobsQ

	FilterByID(id ...int64) BlobsQ

	FilterByOwnerId(id ...int64) BlobsQ
}

type Blob struct {
	ID        int64           `db:"id" structs:"-"`
	CreatedAt time.Time       `db:"created_at" structs:"created_at"`
	Value     json.RawMessage `db:"value" structs:"-"`
	OwnerId   int64           `db:"owner_id" structs:"-"`
}

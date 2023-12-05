package pg

import (
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/dl7850949/blob-storage/internal/data"
)

const blobsTableName = "blobs"

type blobsQ struct {
	db  *pgdb.DB
	sql sq.SelectBuilder
}

func NewBlobsQ(db *pgdb.DB) data.BlobsQ {
	return &blobsQ{
		db:  db.Clone(),
		sql: sq.Select("n.*").From(fmt.Sprintf("%s as n", blobsTableName)),
	}
}

func (q *blobsQ) New() data.BlobsQ {
	return NewBlobsQ(q.db)
}

func (q *blobsQ) Get() (*data.Blob, error) {
	return nil, nil
}

func (q *blobsQ) Delete() (*data.Blob, error) {
	return nil, nil
}

func (q *blobsQ) Select() ([]data.Blob, error) {
	return nil, nil
}

func (q *blobsQ) Transaction(fn func(q data.BlobsQ) error) error {
	return q.db.Transaction(func() error {
		return fn(q)
	})
}

func (q *blobsQ) Insert(value data.Blob) (*data.Blob, error) {
	value.CreatedAt = time.Now().UTC()
	clauses := structs.Map(value)
	fmt.Println(clauses)
	clauses["value"] = value.Value

	var result data.Blob
	stmt := sq.Insert(blobsTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)

	return &result, err
}

func (q *blobsQ) Page(pageParams pgdb.OffsetPageParams) data.BlobsQ {
	return nil
}

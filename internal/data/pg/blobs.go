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
	var result data.Blob
	err := q.db.Get(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q *blobsQ) Delete(id int64) error {
	stmt := sq.Delete(blobsTableName).Where(sq.Eq{"id": id})
	err := q.db.Exec(stmt)
	return err
}

func (q *blobsQ) Select() ([]data.Blob, error) {
	var result []data.Blob
	err := q.db.Select(&result, q.sql)
	return result, err
}

func (q *blobsQ) Transaction(fn func(q data.BlobsQ) error) error {
	return q.db.Transaction(func() error {
		return fn(q)
	})
}

func (q *blobsQ) Insert(value data.Blob) (*data.Blob, error) {
	value.CreatedAt = time.Now().UTC()
	clauses := structs.Map(value)
	clauses["value"] = value.Value
	clauses["owner_id"] = value.OwnerId

	var result data.Blob
	stmt := sq.Insert(blobsTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)

	return &result, err
}

func (q *blobsQ) Page(pageParams pgdb.OffsetPageParams) data.BlobsQ {
	return nil
}

func (q *blobsQ) FilterByID(ids ...int64) data.BlobsQ {
	q.sql = q.sql.Where(sq.Eq{"n.id": ids})
	return q
}

func (q *blobsQ) FilterByOwnerId(ids ...int64) data.BlobsQ {
	q.sql = q.sql.Where(sq.Eq{"n.owner_id": ids})
	return q
}

package horizon

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/go/xdrbuild"
	"gitlab.com/tokend/horizon-connector"
	regources "gitlab.com/tokend/regources/generated"
)

// iota is used in case if we need more types in the future
const (
	BLOB = iota
)

// endpoints
const (
	TRANSACTIONS = "/v3/transactions"
	DATA         = "/v3/data"
)

type Blob struct {
	Value json.RawMessage
}

func (blob Blob) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value json.RawMessage `json:"value"`
	}{
		Value: blob.Value,
	})
}

type BlobsOps struct {
	Tx      *xdrbuild.Transaction
	Horizon *horizon.Connector
}

func New(tx *xdrbuild.Transaction, horizon *horizon.Connector) *BlobsOps {
	return &BlobsOps{
		Tx:      tx,
		Horizon: horizon,
	}
}

func (ops BlobsOps) Create(blob Blob, waitForIngest *bool) (int, []byte, error) {
	tx := ops.Tx.Op(&xdrbuild.CreateData{
		Type:  BLOB,
		Value: blob,
	})

	envelope, err := tx.Marshal()
	if err != nil {
		return http.StatusInternalServerError, nil, errors.Wrap(err, "failed to build tx envelope")
	}

	return ops.Horizon.Client().PostJSON(TRANSACTIONS, &regources.SubmitTransactionBody{
		Tx:            envelope,
		WaitForIngest: waitForIngest,
	})
}

func (ops BlobsOps) Get(pageParams pgdb.OffsetPageParams) ([]byte, error) {
	endpoint := fmt.Sprintf(DATA+"?page[limit]=%d&page[number]=%d&page[order]=%s",
		pageParams.Limit, pageParams.PageNumber, pageParams.Order)
	return ops.Horizon.Client().Get(endpoint)
}

func (ops BlobsOps) GetById(id int64) ([]byte, error) {
	endpoint := fmt.Sprintf(DATA+"/%d", id)
	return ops.Horizon.Client().Get(endpoint)
}

func (ops BlobsOps) Delete(id int64, waitForIngest *bool) (int, []byte, error) {
	tx := ops.Tx.Op(&xdrbuild.RemoveData{
		ID: uint64(id),
	})

	envelope, err := tx.Marshal()
	if err != nil {
		return http.StatusInternalServerError, nil, errors.Wrap(err, "failed to build tx envelope")
	}

	return ops.Horizon.Client().PostJSON(TRANSACTIONS, &regources.SubmitTransactionBody{
		Tx:            envelope,
		WaitForIngest: waitForIngest,
	})
}

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

type Blob struct {
	Value   json.RawMessage
	OwnerId int64
}

func (blob Blob) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Value   json.RawMessage `json:"value"`
		OwnerId int64           `json:"owner_id"`
	}{
		Value:   blob.Value,
		OwnerId: blob.OwnerId,
	})
}

type BlobsOps struct {
	Tx      *xdrbuild.Transaction
	Horizon *horizon.Connector
}

func New(tx *xdrbuild.Transaction, horizon *horizon.Connector) BlobsOps {
	return BlobsOps{
		Tx:      tx,
		Horizon: horizon,
	}
}

func (ops BlobsOps) Create(blob Blob, waitForIngest *bool) (int, []byte, error) {
	tx := ops.Tx
	tx = tx.Op(&xdrbuild.CreateData{
		Type:  BLOB,
		Value: blob,
	})

	envelope, err := tx.Marshal()
	if err != nil {
		return http.StatusInternalServerError, nil, errors.Wrap(err, "failed to build tx envelope")
	}

	status, response, err := ops.Horizon.Client().PostJSON("/v3/transactions", &regources.SubmitTransactionBody{
		Tx:            envelope,
		WaitForIngest: waitForIngest,
	})
	if err != nil {
		return http.StatusInternalServerError, nil, errors.Wrap(err, "failed to insert blob")
	}

	return status, response, nil
}

func (ops BlobsOps) Get(pageParams pgdb.OffsetPageParams) ([]byte, error) {
	endpoint := fmt.Sprintf("/v3/data?page[limit]=%d&page[number]=%d&page[order]=%s", pageParams.Limit, pageParams.PageNumber, pageParams.Order)
	response, err := ops.Horizon.Client().Get(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get blob list")
	}

	return response, nil
}

func (ops BlobsOps) GetById(id int64) ([]byte, error) {
	endpoint := fmt.Sprintf("/v3/data/%d", id)
	response, err := ops.Horizon.Client().Get(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get blob")
	}

	return response, nil
}

func (ops BlobsOps) Delete(id int64, waitForIngest *bool) (int, []byte, error) {
	tx := ops.Tx
	tx = tx.Op(&xdrbuild.RemoveData{
		ID: uint64(id),
	})

	envelope, err := tx.Marshal()
	if err != nil {
		return http.StatusInternalServerError, nil, errors.Wrap(err, "failed to build tx envelope")
	}

	status, response, err := ops.Horizon.Client().PostJSON("/v3/transactions", &regources.SubmitTransactionBody{
		Tx:            envelope,
		WaitForIngest: waitForIngest,
	})
	if err != nil {
		return http.StatusInternalServerError, nil, errors.Wrap(err, "failed to delete blob")
	}

	return status, response, nil
}

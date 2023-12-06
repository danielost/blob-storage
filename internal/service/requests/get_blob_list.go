package requests

import (
	"net/http"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/urlval"
)

type GetNotificationsListRequest struct {
	pgdb.OffsetPageParams
}

func NewGetBlobsListRequest(r *http.Request) (GetNotificationsListRequest, error) {
	var request GetNotificationsListRequest

	err := urlval.Decode(r.URL.Query(), &request)
	if err != nil {
		return request, err
	}

	return request, nil
}

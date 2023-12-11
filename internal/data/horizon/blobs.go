package horizon

import (
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/tokend/horizon-connector"
)

type BlobsQ struct {
	log     *logan.Entry
	horizon *horizon.Connector
}

type voidParams struct {
}

func (p *voidParams) Encode() string {
	return ""
}

func NewBlobsQ(log *logan.Entry, horizon *horizon.Connector) *BlobsQ {
	return &BlobsQ{
		horizon: horizon,
		log:     log,
	}
}

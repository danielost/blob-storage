package service

import (
	"context"
	"time"

	"gitlab.com/dl7850949/blob-storage/internal/data"
	"gitlab.com/tokend/regources"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/running"
)

type LazyInfo struct {
	ctx    context.Context
	log    *logan.Entry
	infoer data.Info
	info   *regources.Info
}

func NewLazyInfo(ctx context.Context, log *logan.Entry, infoer data.Info) *LazyInfo {
	return &LazyInfo{
		ctx:    ctx,
		log:    log,
		infoer: infoer,
	}
}

func (i *LazyInfo) Info() (*regources.Info, error) {
	if i.info == nil {
		i.obtainInfo()
	}
	return i.info, nil
}

func (i *LazyInfo) obtainInfo() {
	running.UntilSuccess(i.ctx, i.log, "info-getter", func(ctx context.Context) (bool, error) {
		info, err := i.infoer.Info()
		if err != nil {
			return false, err
		}
		i.info = info
		return true, nil
	}, 1*time.Second, 1*time.Minute)
}

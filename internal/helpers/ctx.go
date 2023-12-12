package helpers

import (
	"context"
	"net/http"

	"gitlab.com/dl7850949/blob-storage/internal/data"
	horizon2 "gitlab.com/dl7850949/blob-storage/internal/data/horizon"
	"gitlab.com/tokend/go/xdrbuild"
	"gitlab.com/tokend/horizon-connector"
	"gitlab.com/tokend/keypair"
	"gitlab.com/tokend/regources"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	blobsQCtxKey
	horizonBlobsQCtxKey
	usersQCtxKey
	horizonConnectorCtxKey
	txBuilderCtxKey
	coreInfoCtxKey
)

func BlobsOps(r *http.Request) *horizon2.BlobsOps {
	txbuilderbuilder := r.Context().Value(txBuilderCtxKey).(data.Infobuilder)
	info := r.Context().Value(coreInfoCtxKey).(data.Info)
	coreInfo, err := info.Info()
	master, err := keypair.ParseAddress(coreInfo.MasterAccountID)

	if err != nil {
		Log(r).WithError(err).Error("error parsing address")
		return nil
	}

	tx := txbuilderbuilder(info, master)

	return horizon2.New(
		tx,
		Horizon(r),
	)
}

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

// Postgres
func CtxBlobsQ(entry data.BlobsQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, blobsQCtxKey, entry)
	}
}

// Postgres
func BlobsQ(r *http.Request) data.BlobsQ {
	return r.Context().Value(blobsQCtxKey).(data.BlobsQ).New()
}

func CtxUsersQ(entry data.UsersQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, usersQCtxKey, entry)
	}
}

func UsersQ(r *http.Request) data.UsersQ {
	return r.Context().Value(usersQCtxKey).(data.UsersQ).New()
}

func CtxHorizon(q *horizon.Connector) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, horizonConnectorCtxKey, q)
	}
}

func Horizon(r *http.Request) *horizon.Connector {
	return r.Context().Value(horizonConnectorCtxKey).(*horizon.Connector).Clone()
}

func CtxCoreInfo(s data.Info) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, coreInfoCtxKey, s)
	}
}

func CoreInfo(r *http.Request) *regources.Info {
	info, err := r.Context().Value(coreInfoCtxKey).(data.Info).Info()
	if err != nil {
		//TODO handle error
		panic(err)
	}
	return info
}

func CtxTransaction(txbuilder data.Infobuilder) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		ctx = context.WithValue(ctx, txBuilderCtxKey, txbuilder)
		return ctx
	}
}

func Transaction(r *http.Request) *xdrbuild.Transaction {
	txbuilderbuilder := r.Context().Value(txBuilderCtxKey).(data.Infobuilder)
	info := r.Context().Value(coreInfoCtxKey).(data.Info)
	master := keypair.MustParseAddress(CoreInfo(r).MasterAccountID)
	return txbuilderbuilder(info, master)
}

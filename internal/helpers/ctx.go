package helpers

import (
	"context"
	"net/http"

	"gitlab.com/dl7850949/blob-storage/internal/data"
	horizon2 "gitlab.com/dl7850949/blob-storage/internal/data/horizon"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	blobsQCtxKey
	horizonBlobsQCtxKey
	usersQCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxBlobsQ(entry data.BlobsQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, blobsQCtxKey, entry)
	}
}

func BlobsQ(r *http.Request) data.BlobsQ {
	return r.Context().Value(blobsQCtxKey).(data.BlobsQ).New()
}

func CtxHBlobsQ(q *horizon2.BlobsQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, horizonBlobsQCtxKey, q)
	}
}

func HBlobsQ(r *http.Request) *horizon2.BlobsQ {
	return r.Context().Value(horizonBlobsQCtxKey).(*horizon2.BlobsQ)
}

func CtxUsersQ(entry data.UsersQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, usersQCtxKey, entry)
	}
}

func UsersQ(r *http.Request) data.UsersQ {
	return r.Context().Value(usersQCtxKey).(data.UsersQ).New()
}

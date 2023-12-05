package service

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/dl7850949/blob-storage/internal/data/pg"
	"gitlab.com/dl7850949/blob-storage/internal/service/handlers"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxBlobsQ(pg.NewBlobsQ(s.db)),
		),
	)
	r.Route("/integrations/blobs", func(r chi.Router) {
		r.Post("/", handlers.CreateBlob)
		// r.Get("/", func(w http.ResponseWriter, r *http.Request) {})
		// r.Route("/{id}", func(r chi.Router) {
		// 	r.Get("/", func(w http.ResponseWriter, r *http.Request) {})
		// 	r.Delete("/", func(w http.ResponseWriter, r *http.Request) {})
		// })
	})

	return r
}

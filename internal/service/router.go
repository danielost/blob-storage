package service

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/dl7850949/blob-storage/internal/data/pg"
	"gitlab.com/dl7850949/blob-storage/internal/middleware"
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
			handlers.CtxUsersQ(pg.NewUsersQ(s.db)),
		),
	)

	r.Route("/integrations/blobs", func(r chi.Router) {
		// Protected routes
		r.Route("/", func(r chi.Router) {
			// Custom JWT middleware
			r.Use(middleware.ValidateJWT)

			r.Post("/", handlers.CreateBlob)
			r.Get("/", handlers.GetBlobsList)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", handlers.GetBlob)
				r.Delete("/", handlers.DeleteBlob)
			})
		})

		// Unprotected routes
		r.Post("/signup", handlers.SignUp)
		r.Post("/signin", handlers.SignIn)
	})

	return r
}

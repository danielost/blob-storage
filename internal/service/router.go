package service

import (
	"os"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	horizon2 "gitlab.com/dl7850949/blob-storage/internal/data/horizon"
	"gitlab.com/dl7850949/blob-storage/internal/data/pg"
	"gitlab.com/dl7850949/blob-storage/internal/helpers"
	"gitlab.com/dl7850949/blob-storage/internal/middleware"
	"gitlab.com/dl7850949/blob-storage/internal/service/handlers"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()
	jwtSecret := os.Getenv("JWT_SECRET")
	horizonBlobsQ := horizon2.NewBlobsQ(s.log, s.horizon)

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			helpers.CtxLog(s.log),
			helpers.CtxBlobsQ(pg.NewBlobsQ(s.db)),
			helpers.CtxHBlobsQ(horizonBlobsQ),
			helpers.CtxUsersQ(pg.NewUsersQ(s.db)),
		),
	)

	r.Route("/integrations/blobs", func(r chi.Router) {
		// Protected routes
		r.Route("/", func(r chi.Router) {
			// Custom JWT middleware
			r.Use(middleware.ValidateJWT(jwtSecret))

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

package service

import (
	"os"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/dl7850949/blob-storage/internal/data"
	"gitlab.com/dl7850949/blob-storage/internal/data/pg"
	"gitlab.com/dl7850949/blob-storage/internal/helpers"
	"gitlab.com/dl7850949/blob-storage/internal/middleware"
	"gitlab.com/dl7850949/blob-storage/internal/service/handlers"
	"gitlab.com/tokend/go/xdrbuild"
	"gitlab.com/tokend/keypair"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()
	jwtSecret := os.Getenv("JWT_SECRET")
	builder := func(info data.Info, source keypair.Address) *xdrbuild.Transaction {
		inf, err := info.Info()
		if err != nil {
			s.log.WithError(err).Error("failed to fetch info")
			return nil
		}

		master := os.Getenv("MASTER_SEED")
		kp, err := keypair.ParseSeed(master)
		if err != nil {
			s.log.WithError(err).Error("failed to parse master seed")
			return nil
		}

		return xdrbuild.
			NewBuilder(inf.GetPassphrase(), inf.GetTXExpire()).
			Transaction(source).Sign(kp)
	}

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			helpers.CtxLog(s.log),
			helpers.CtxBlobsQ(pg.NewBlobsQ(s.db)),
			helpers.CtxUsersQ(pg.NewUsersQ(s.db)),
			helpers.CtxHorizon(s.horizon),
			helpers.CtxCoreInfo(s.infoer),
			helpers.CtxTransaction(builder),
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

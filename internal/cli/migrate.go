package cli

import (
	migrate "github.com/rubenv/sql-migrate"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/dl7850949/blob-storage/internal/assets"
	"gitlab.com/dl7850949/blob-storage/internal/config"
)

var migrations = &migrate.PackrMigrationSource{
	Box: assets.Migrations,
}

func MigrateUp(cfg config.Config) error {
	applied, err := migrate.Exec(cfg.DB().RawDB(), "postgres", migrations, migrate.Up)
	if err != nil {
		return errors.Wrap(err, "failed to apply migrations")
	}
	cfg.Log().WithField("applied", applied).Info("migrations applied")
	return nil
}

func MigrateDown(cfg config.Config) error {
	applied, err := migrate.Exec(cfg.DB().RawDB(), "postgres", migrations, migrate.Down)
	if err != nil {
		return errors.Wrap(err, "failed to apply migrations")
	}
	cfg.Log().WithField("applied", applied).Info("migrations applied")
	return nil
}

package repo

import (
	"context"

	"github.com/cockroachdb/errors"
	"gorm.io/gorm"
)

type DbInitializer struct {
	Db *gorm.DB
}

func (p *DbInitializer) Name() string {
	return "db_initializer"
}

func (p *DbInitializer) IsNeedInit(ctx context.Context) (bool, error) {
	return true, nil
}

// Initialize AutoMigrate自动建表。
func (p *DbInitializer) Initialize(ctx context.Context) error {
	if err := p.migrate(ctx); err != nil {
		return err
	}

	return nil
}

func (p *DbInitializer) migrate(ctx context.Context) error {
	err := p.Db.WithContext(ctx).AutoMigrate()

	return errors.WithStack(err)
}

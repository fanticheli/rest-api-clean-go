package productrepository

import (
	"github.com/fanticheli/rest-api-clean-go/adapter/postgres"
	"github.com/fanticheli/rest-api-clean-go/core/domain"
)

type repository struct {
	db postgres.PoolInterface
}

// New returns contract implementation of ProductRepository
func New(db postgres.PoolInterface) domain.ProductRepository {
	return &repository{
		db: db,
	}
}

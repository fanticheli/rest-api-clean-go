package di

import (
	"github.com/fanticheli/rest-api-clean-go/adapter/http/rest/productservice"
	"github.com/fanticheli/rest-api-clean-go/adapter/postgres"
	"github.com/fanticheli/rest-api-clean-go/adapter/postgres/productrepository"
	"github.com/fanticheli/rest-api-clean-go/core/domain"
	"github.com/fanticheli/rest-api-clean-go/core/usecase/productusecase"
)

// ConfigProductDI return a ProductService abstraction with dependency injection configuration
func ConfigProductDI(conn postgres.PoolInterface) domain.ProductService {
	productRepository := productrepository.New(conn)
	productUseCase := productusecase.New(productRepository)
	productService := productservice.New(productUseCase)

	return productService
}

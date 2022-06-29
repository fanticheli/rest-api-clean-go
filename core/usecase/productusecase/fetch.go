package productusecase

import (
	"github.com/fanticheli/rest-api-clean-go/core/domain"
	"github.com/fanticheli/rest-api-clean-go/core/dto"
)

func (usecase usecase) Fetch(paginationRequest *dto.PaginationRequestParms) (*domain.Pagination, error) {
	products, err := usecase.repository.Fetch(paginationRequest)

	if err != nil {
		return nil, err
	}

	return products, nil
}

package products

type Service interface {
	GetAllProducts() ([]Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllProducts() ([]Product, error) {
	products, err := s.repository.GetProducts()
	if err != nil {
		return products, err
	}

	return products, nil
}

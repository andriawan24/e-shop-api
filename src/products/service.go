package products

type Service interface {
	GetAllProducts() ([]Product, error)
	GetCategories() ([]Category, error)
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

func (s *service) GetCategories() ([]Category, error) {
	categories, err := s.repository.GetCategories()
	if err != nil {
		return categories, err
	}

	return categories, nil
}

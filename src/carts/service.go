package carts

type Service interface {
	GetUserCart(userID int) (Cart, error)
	SaveCart(input SaveCartInput, UserID int) (Cart, error)
	RemoveCart(userID int) (bool, error)
	RemoveProduct(userID int, input RemoveProductInput) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetUserCart(userID int) (Cart, error) {

	cart, err := s.repository.GetUserCart(userID)
	if err != nil {
		return cart, err
	}

	return cart, nil
}

func (s *service) SaveCart(input SaveCartInput, userID int) (Cart, error) {

	cart, err := s.repository.CreateOrGetCartByUserID(userID)
	if err != nil {
		return cart, err
	}

	detail := CartDetail{
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
		CartID:    cart.ID,
	}

	cart, err = s.repository.SaveCart(detail, cart)
	if err != nil {
		return cart, err
	}

	return cart, nil
}

func (s *service) RemoveCart(userID int) (bool, error) {
	cart, err := s.repository.GetUserCart(userID)
	if err != nil {
		return false, err
	}

	success, err := s.repository.RemoveCart(cart)
	if err != nil {
		return false, err
	}

	return success, err
}

func (s *service) RemoveProduct(userID int, input RemoveProductInput) (bool, error) {
	cart, err := s.repository.GetUserCart(userID)
	if err != nil {
		return false, err
	}

	success, err := s.repository.RemoveProduct(cart.ID, input.ProductID)
	if err != nil {
		return false, err
	}

	return success, nil
}

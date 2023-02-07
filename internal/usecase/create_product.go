package usecase

import "gituhb.com/gothello/live-full-cycle-kafka/internal/entity"

type CreateProductInput struct {
	Name  string
	Price float64
}

type CreateProductOutput struct {
	ID    string
	Name  string
	Price float64
}

type CreateProductUseCase struct {
	ProductRepository entity.ProductRepository
}

func NewCreateProductUseCase(pr entity.ProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{
		ProductRepository: pr,
	}
}

func (u *CreateProductUseCase) Execute(input CreateProductInput) (*CreateProductOutput, error) {
	prod := entity.NewProduct(input.Name, input.Price)

	if err := u.ProductRepository.Insert(prod); err != nil {
		return nil, err
	}

	return &CreateProductOutput{
		ID:    prod.ID,
		Name:  prod.Name,
		Price: prod.Price,
	}, nil
}

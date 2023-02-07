package usecase

import "gituhb.com/gothello/live-full-cycle-kafka/internal/entity"

type ListProductOutput struct {
	ID    string
	Name  string
	Price float64
}

type ListProductUseCase struct {
	ProductRepository entity.ProductRepository
}

func NewListProductUseCase(pr entity.ProductRepository) *ListProductUseCase {
	return &ListProductUseCase{
		ProductRepository: pr,
	}
}

func (u *ListProductUseCase) Execute() ([]*ListProductOutput, error) {
	ps, err := u.ProductRepository.GetAll()
	if err != nil {
		return nil, err
	}

	var res []*ListProductOutput

	for _, p := range ps {

		res = append(res, &ListProductOutput{p.ID, p.Name, p.Price})
	}

	return res, nil
}

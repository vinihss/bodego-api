package product

type DeleteProductUseCase struct {
	repo ProductRepository
}

func NewDeleteProductUseCase(repo ProductRepository) *DeleteProductUseCase {
	return &DeleteProductUseCase{repo: repo}
}

func (uc *DeleteProductUseCase) Execute(id uint) error {
	entity, err := uc.repo.FindByID(id)
	if err != nil {
		return err
	}
	return uc.repo.Delete(entity.GetID())
}

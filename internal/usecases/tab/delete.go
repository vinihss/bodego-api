package tab

type DeleteTabUseCase struct {
	repo TabRepository
}

func NewDeleteTabUseCase(repo TabRepository) *DeleteTabUseCase {
	return &DeleteTabUseCase{repo: repo}
}

func (uc *DeleteTabUseCase) Execute(id uint) error {
	entity, err := uc.repo.FindByID(id)
	if err != nil {
		return err
	}
	return uc.repo.Delete(entity.ID)
}

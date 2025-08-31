package news

type DeleteNewsUseCase struct {
	repo NewsRepository
}

func NewDeleteNewsUseCase(repo NewsRepository) *DeleteNewsUseCase {
	return &DeleteNewsUseCase{repo: repo}
}

func (uc *DeleteNewsUseCase) Execute(id uint) error {
	return uc.repo.Delete(id)
}
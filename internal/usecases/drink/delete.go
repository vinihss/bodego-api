package drink

import "github.com/vinihss/bodego-api/internal/domain/drink"

type DeleteDrinkUseCase struct {
	Repo drink.Repository
}

func (uc *DeleteDrinkUseCase) Execute(id uint) error {
	return uc.Repo.Delete(id)
}

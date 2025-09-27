package drink

import "github.com/vinihss/bodego-api/internal/domain/drink"

type FindDrinkUseCase struct {
	Repo drink.Repository
}

func (uc *FindDrinkUseCase) Execute(id uint) (*drink.Drink, error) {
	return uc.Repo.FindByID(id)
}

func (uc *FindDrinkUseCase) ExecuteAll() ([]drink.Drink, error) {
	return uc.Repo.FindAll()
}

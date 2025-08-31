package drink

import "github.com/vinihss/bodego-api/internal/domain/drink"

type UpdateDrinkUseCase struct {
	Repo drink.Repository
}

func (uc *UpdateDrinkUseCase) Execute(d drink.Drink) (*drink.Drink, error) {
	return uc.Repo.Update(d)
}

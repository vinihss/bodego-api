package drink

import "github.com/vinihss/bodego-api/internal/domain/drink"

type CreateDrinkUseCase struct {
	Repo drink.Repository
}

func (uc *CreateDrinkUseCase) Execute(d drink.Drink) (drink.Drink, error) {
	return uc.Repo.Create(d)
}

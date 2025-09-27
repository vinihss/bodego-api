package http_interfaces_drink

import (
	"github.com/vinihss/bodego-api/internal/usecases/drink"
)

type DrinkController struct {
	createUC *drink.CreateDrinkUseCase
	findUC   *drink.FindDrinkUseCase
	updateUC *drink.UpdateDrinkUseCase
	deleteUC *drink.DeleteDrinkUseCase
}

func NewDrinkController(
	createUC *drink.CreateDrinkUseCase,
	findUC *drink.FindDrinkUseCase,
	updateUC *drink.UpdateDrinkUseCase,
	deleteUC *drink.DeleteDrinkUseCase,
) *DrinkController {
	return &DrinkController{
		createUC: createUC,
		findUC:   findUC,
		updateUC: updateUC,
		deleteUC: deleteUC,
	}
}

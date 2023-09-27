package api

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gitsridhar/myopenapi3/echo/Impl/api/models"
	echo "github.com/labstack/echo/v4"
)

type FruitStore struct {
	Fruits map[int64]models.Fruit
	NextId int64
	Lock   sync.Mutex
}

func NewFruitStore() *FruitStore {
	return &FruitStore{
		Fruits: make(map[int64]models.Fruit),
		NextId: 1000,
	}
}

func sendFruitStoreError(ctx echo.Context, code int, message string) error {
	fruitErr := models.Error{
		Code:    int32(code),
		Message: message,
	}
	err := ctx.JSON(code, fruitErr)
	return err
}

func (f *FruitStore) FindFruits(ctx echo.Context, params models.FindFruitsParams) error {
	f.Lock.Lock()
	defer f.Lock.Unlock()

	var result []models.Fruit

	for _, fruit := range f.Fruits {
		if params.Tags != nil {
			for _, t := range *params.Tags {
				if fruit.Tag != nil && (*fruit.Tag == t) {
					result = append(result, fruit)
				}
			}
		} else {
			result = append(result, fruit)
		}

		if params.Limit != nil {
			l := int(*params.Limit)
			if len(result) >= l {
				break
			}
		}
	}

	return ctx.JSON(http.StatusOK, result)
}

func (f *FruitStore) AddFruit(ctx echo.Context) error {
	var newFruit models.NewFruit
	err := ctx.Bind(&newFruit)
	if err != nil {
		return sendFruitStoreError(ctx, http.StatusBadRequest, "Invalid format for newFruit")
	}

	f.Lock.Lock()
	defer f.Lock.Unlock()

	var fruit models.Fruit
	fruit.Name = newFruit.Name
	fruit.Tag = newFruit.Tag
	fruit.Id = f.NextId
	f.NextId++

	f.Fruits[fruit.Id] = fruit

	err = ctx.JSON(http.StatusCreated, fruit)
	if err != nil {
		return err
	}

	return nil
}

func (f *FruitStore) FindFruitByID(ctx echo.Context, fruitId int64) error {
	f.Lock.Lock()
	defer f.Lock.Unlock()

	fruit, found := f.Pets[fruitId]
	if !found {
		return sendFruitStoreError(ctx, http.StatusNotFound, fmt.Sprintf("Could not find fruit with ID %d", fruitId))
	}
	return ctx.JSON(http.StatusOK, fruit)
}

func (f *fruitStore) DeletePet(ctx echo.Context, id int64) error {
	f.Lock.Lock()
	defer f.Lock.Unlock()

	_, found := f.Pets[id]
	if !found {
		return sendFruitStoreError(ctx, http.StatusNotFound, fmt.Sprintf("Could not find fruit with ID %d", id))
	}
	delete(f.Pets, id)
	return ctx.JSON(http.StatusNoContent)
}

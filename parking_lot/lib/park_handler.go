package lib

import (
	"errors"
	"fmt"
	"strconv"

	"parking_lot/models"
)

type ParkHandler struct{}

func (ParkHandler) ValidateHandler(args ...string) error {
	if len(args) != 2 {
		return errors.New("Invalid parameter number")
	}

	return nil
}

func (ParkHandler) ExecuteHandler(PS *models.ParkingStore, args ...string) (string, error) {
	car := models.Car{
		Number: args[0],
		Color:  args[1],
	}

	if slot, _ := PS.Park(car); slot > 0 {
		return fmt.Sprintf("Allocated slot number: %d", slot), nil
	}

	return "", errors.New("Sorry, parking lot is full")
}

type CreateHandler struct{}

func (CreateHandler) ValidateHandler(args ...string) error {
	if len(args) == 0 {
		return errors.New("No parameter")
	}

	if len(args) > 1 {
		return errors.New("Too much parameter")
	}

	return nil
}

func (CreateHandler) ExecuteHandler(PS *models.ParkingStore, args ...string) (string, error) {
	slot, _ := strconv.Atoi(args[0])

	if slot <= 0 {
		return "", errors.New("Invalid slot count")
	}

	//fmt.Printf("Created a parking lot with %d slots\n", slot)

	PS.Create(slot)

	//fmt.Printf("PS count %d\n", PS.Count())
	return fmt.Sprintf("Created a parking lot with %d slots", slot), nil
}

package lib

import (
	"errors"
	"strconv"

	"parking_lot/models"
)

type LeaveHandler struct{}

func (LeaveHandler) ValidateHandler(args ...string) error {
	if len(args) != 1 {
		return errors.New("Invalid parameter number")
	}

	return nil
}

func (LeaveHandler) ExecuteHandler(PS *models.ParkingStore, args ...string) (string, error) {
	slot, err := strconv.Atoi(args[0])
	if err != nil {
		return "Input is not number", errors.New("Input is not number")
	}

	if ok := PS.Leave(slot); ok {
		return "Slot number " + args[0] + " is free", nil
	}

	return "Slot number " + args[0] + " is not empty", errors.New("Slot number " + args[0] + " is not empty")
}

package lib

import (
	"errors"
	"strings"

	"parking_lot/models"
)

type NumberColorHandler struct{}

func (NumberColorHandler) ValidateHandler(args ...string) error {
	if len(args) != 1 {
		return errors.New("Invalid parameter number")
	}

	return nil
}

func (NumberColorHandler) ExecuteHandler(PS *models.ParkingStore, args ...string) (string, error) {
	var retNumber []string
	if result := PS.Search(models.WithCarColour(args[0])); len(result) > 0 {
		for _, slot := range result {
			Car := PS.GetCarAtSlot(slot)

			if Car != nil {
				retNumber = append(retNumber, Car.Number)
			}

		}
		return strings.Join(retNumber, ", "), nil
	}

	return "Not found", errors.New("Not found")
}

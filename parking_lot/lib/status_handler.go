package lib

import (
	"errors"
	"fmt"
	"strings"

	"parking_lot/models"

	"github.com/thoas/go-funk"
)

type StatusHandler struct{}

func (StatusHandler) ValidateHandler(args ...string) error {
	if len(args) != 0 {
		return errors.New("Invalid parameter")
	}

	return nil
}

func (StatusHandler) ExecuteHandler(PS *models.ParkingStore, args ...string) (string, error) {
	no := 1
	patternHead := "%s\t%s\t%s\n"

	pattern := "%d\t\t%s\t%s\n"
	ret := fmt.Sprintf(patternHead, "Slot No.", "Registration No", "Colour")

	funk.ForEach(*PS, func(p models.ParkingLot) {
		if p.Car != nil && (*p.Car).Number != "" {
			ret += fmt.Sprintf(pattern, no, p.Car.Number, p.Car.Color)
		}

		no++
	})

	return strings.Trim(ret, "\n"), nil
}

package lib

import (
	"bytes"
	"errors"
	"strconv"

	"parking_lot/models"
)

type SlotNumberHandler struct{}

func (SlotNumberHandler) ValidateHandler(args ...string) error {
	if len(args) != 1 {
		return errors.New("Invalid parameter number")
	}

	return nil
}

func (SlotNumberHandler) ExecuteHandler(PS *models.ParkingStore, args ...string) (string, error) {

	if result := PS.Search(models.WithCarNumber(args[0])); len(result) > 0 {
		return arrayToString(result, ","), nil
	}

	return "Not found", errors.New("Not found")
}

type SlotColorHandler struct{}

func (SlotColorHandler) ValidateHandler(args ...string) error {
	if len(args) != 1 {
		return errors.New("Invalid parameter number")
	}

	return nil
}

func (SlotColorHandler) ExecuteHandler(PS *models.ParkingStore, args ...string) (string, error) {

	if result := PS.Search(models.WithCarColour(args[0])); len(result) > 0 {
		return arrayToString(result, ","), nil
	}

	return "Not found", errors.New("Not found")
}

func arrayToString(A []int, delim string) string {
	var buffer bytes.Buffer
	for i := 0; i < len(A); i++ {
		buffer.WriteString(strconv.Itoa(A[i]))
		if i != len(A)-1 {
			buffer.WriteString(delim)
		}
	}

	return buffer.String()
}

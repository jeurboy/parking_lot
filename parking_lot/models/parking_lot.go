package models

import (
	"github.com/thoas/go-funk"
)

type ParkingStore []ParkingLot
type ParkingLot struct {
	Car *Car
}

type CarSearchOption func(Car) bool

func WithCarColour(color string) CarSearchOption {
	return func(c Car) bool {
		return c.Color == color
	}
}
func WithCarNumber(number string) CarSearchOption {
	return func(c Car) bool {
		return c.Number == number
	}
}

func (ps *ParkingStore) Create(i int) bool {
	np := make([]ParkingLot, i)

	*ps = np
	return true
}

func (ps ParkingStore) Count() (int, int) {
	var index int
	var vacant int

	funk.ForEach(ps, func(p ParkingLot) {
		index++
		if p.Car == nil || (*p.Car).Number == "" {
			vacant++
		}
	})

	return index, vacant
}

func (ps *ParkingStore) Leave(slot int) bool {
	if slot-1 < 0 {
		return false
	}

	all, _ := ps.Count()

	if slot > all {
		return false
	}

	(*ps)[slot-1].Car = nil
	return true
}

func (ps *ParkingStore) Park(c Car) (int, bool) {
	var index int

	if index = ps.FindVacantPark(); index <= 0 {
		return 0, false
	}

	(*ps)[index-1] = ParkingLot{
		Car: &c,
	}

	return index, true
}

func (ps ParkingStore) FindVacantPark() int {
	slot := 0
	found := 0

	funk.ForEach(ps, func(p ParkingLot) {
		if found > 0 {
			return
		}
		slot++

		if p.Car == nil || (*p.Car).Number == "" {
			found = slot
		}
	})

	all, _ := ps.Count()
	if slot > all {
		return 0
	}

	return found
}

func (ps ParkingStore) Search(fxs ...CarSearchOption) []int {
	slots := []int{}

	index := 0

	funk.ForEach(ps, func(p ParkingLot) {
		index++
		result := true
		for _, f := range fxs {
			if !f(*p.Car) {
				result = false
			}
		}

		if result {
			slots = append(slots, index)
		}
	})

	return slots
}

func (ps ParkingStore) GetCarAtSlot(slot int) *Car {
	return ps.getCarAtSlot(slot)
}

func (ps ParkingStore) getCarAtSlot(slot int) *Car {
	if PS := ps[slot-1]; PS.Car != nil {
		return ps[slot-1].Car
	}

	return nil
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

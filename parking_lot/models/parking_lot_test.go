package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParkingLotCreate(t *testing.T) {
	assert := assert.New(t)

	var PS ParkingStore
	var all, v int

	assert.True(PS.Create(5), "Fail to create parking slot")

	all, v = PS.Count()
	assert.Equal(5, v, "Vacant parking slot should be 5")
	assert.Equal(5, all, "Parking slot should be 5")
	assert.NotEqual(6, all, "Parking slot should not be 6")
}

func TestParkingLotAdd(t *testing.T) {
	assert := assert.New(t)

	var PS ParkingStore
	var ok bool
	var slot int

	PS.Create(2)
	slot, ok = PS.Park(Car{Number: "KH-1234", Color: "White"})
	assert.True(ok, "Car 1234 can't be added")
	assert.Equal(1, slot, "Car 1234 must be added to slot 1")

	slot, ok = PS.Park(Car{Number: "KH-4567", Color: "White"})
	assert.True(ok, "Car 4567 can't be added")
	assert.Equal(2, slot, "Car 4567 must be added to slot 2")

	slot, ok = PS.Park(Car{Number: "KH-8900", Color: "White"})
	assert.False(ok, "Car can't be added. Empty parking slot")
	assert.Equal(0, slot, "Car 8900 add to slot 0")
}

func TestParkingLotAddFromNotZero(t *testing.T) {
	assert := assert.New(t)

	var PS ParkingStore
	var ok bool
	var slot int

	slot, ok = PS.Park(Car{Number: "KH-8900", Color: "White"})
	assert.False(ok, "Car can't be added. Empty parking slot")
	assert.Equal(0, slot, "Car 8900 add to slot 0")
}

func TestParkingLotLeave(t *testing.T) {
	assert := assert.New(t)

	var PS ParkingStore

	PS.Create(3)
	PS.Park(Car{Number: "KH-1234", Color: "White"})
	PS.Park(Car{Number: "KH-4567", Color: "White"})
	PS.Park(Car{Number: "KH-8900", Color: "White"})

	assert.False(PS.Leave(4), "Cannot leave slot 4 out of 2 parking")
	assert.True(PS.Leave(2), "Can leave slot 2 out of 2 parking")

	assert.Equal(2, PS.FindVacantPark(), "slot 2 should be empty")
}

func TestParkingLotSearchByColour(t *testing.T) {
	assert := assert.New(t)

	var PS ParkingStore
	var result bool
	var resultSlot []int

	PS.Create(7)
	PS.Park(Car{Number: "KH-1234", Color: "White"})
	PS.Park(Car{Number: "KH-4567", Color: "White"})
	PS.Park(Car{Number: "KH-8900", Color: "Black"})
	PS.Park(Car{Number: "KH-1233", Color: "White"})
	PS.Park(Car{Number: "AH-4537", Color: "Black"})
	PS.Park(Car{Number: "BH-1233", Color: "Blue"})
	PS.Park(Car{Number: "RH-4537", Color: "Red"})

	f := WithCarColour("White")
	result = f(Car{Number: "KH-1234", Color: "White"})
	assert.True(result, "Car KH-1234 has white color")

	b := WithCarColour("Black")
	result = b(Car{Number: "KH-1225", Color: "White"})
	assert.False(result, "Car KH-1234 has white color")

	resultSlot = PS.Search(WithCarColour("Black"))
	assert.Equal([]int{3, 5}, resultSlot, "Found car with black colour at 3, 5 parking")
}

func TestParkingLotSearchByNumber(t *testing.T) {
	assert := assert.New(t)

	var PS ParkingStore

	PS.Create(7)
	PS.Park(Car{Number: "KH-1234", Color: "cc"})
	PS.Park(Car{Number: "KH-4567", Color: "vv"})
	PS.Park(Car{Number: "KH-8900", Color: "ff"})
	PS.Park(Car{Number: "KH-1233", Color: "gg"})
	PS.Park(Car{Number: "AH-4537", Color: "hh"})
	PS.Park(Car{Number: "BH-1233", Color: "yy"})
	PS.Park(Car{Number: "RH-4537", Color: "h"})

	resultSlot := PS.Search(WithCarNumber("KH-4567"))
	assert.Equal([]int{2}, resultSlot, "Found car with KH-4567 number at 2 parking")

	resultSlot = PS.Search(WithCarNumber("KH-xxxxx"))
	assert.Equal([]int{}, resultSlot, "Found car with KH-xxxxx number at 0 parking")
}

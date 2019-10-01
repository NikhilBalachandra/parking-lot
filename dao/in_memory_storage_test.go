package dao

import (
	"reflect"
	"testing"
)

func TestInMemoryStorage_Park(t *testing.T) {
	storage := InMemoryStorage{}
	storage.SetSize(6)

	car := Car{
		Color:              "White",
		RegistrationNumber: "KA-01-HH-1234",
	}
	err := storage.Park(1, &car)
	if err != nil {
		t.Errorf("Park() Error parking the car. Error %v", err)
	}
}

func TestInMemoryStorage_ParkShouldNotAllowDuplicateCar(t *testing.T) {
	storage := InMemoryStorage{}
	storage.SetSize(6)
	car := Car{
		Color:              "White",
		RegistrationNumber: "KA-01-HH-1234",
	}
	car1 := Car{
		Color:              "White",
		RegistrationNumber: "KA-01-HH-1234",
	}
	_ = storage.Park(1, &car)
	err := storage.Park(2, &car1)
	if err != ErrDuplicateRegNum {
		t.Errorf("Park() Error got %v want %v", err, ErrDuplicateRegNum)
	}
}
func TestInMemoryStorage_ParkShouldNotAllowOccupiedSlot(t *testing.T) {
	storage := InMemoryStorage{}
	storage.SetSize(6)
	car := Car{
		Color:              "White",
		RegistrationNumber: "KA-01-HH-1234",
	}
	car1 := Car{
		Color:              "White",
		RegistrationNumber: "KA-01-HH-1235",
	}
	_ = storage.Park(1, &car)
	err := storage.Park(1, &car1)
	if err != ErrSlotAlreadyOccupied {
		t.Errorf("Park() Error got %v want %v", err, ErrSlotAlreadyOccupied)
	}
}

func TestInMemoryStorage_ParkShouldNotAllowSlotExceedingCapacity(t *testing.T) {
	storage := InMemoryStorage{}
	storage.SetSize(6)
	car := Car{
		Color:              "White",
		RegistrationNumber: "KA-01-HH-1234",
	}
	err := storage.Park(7, &car)
	if err != ErrSlotExceedsAvailableParking {
		t.Errorf("Park() Error got %v want %v", err, ErrSlotExceedsAvailableParking)
	}
}

func TestInMemoryStorage_Leave(t *testing.T) {
	storage := InMemoryStorage{}
	storage.SetSize(6)
	car := Car{
		Color:              "White",
		RegistrationNumber: "KA-01-HH-1234",
	}
	_ = storage.Park(1, &car)
	car1, err := storage.Leave(1)
	if err != nil {
		t.Errorf("Leave() Error %v", err)
	}
	if car1 != &car {
		t.Errorf("Leave() got %v want %v", car1, car)
	}
}

func TestInMemoryStorage_LeaveErrorOutOnUnoccupiedSlot(t *testing.T) {
	storage := InMemoryStorage{}
	storage.SetSize(6)
	_, err := storage.Leave(1)
	if err != ErrSlotNotOccupied {
		t.Errorf("Leave() Error got %v want %v", err, ErrSlotNotOccupied)
	}
}

func TestInMemoryStorage_LeaveErrorOutOnExceedingAvailableParking(t *testing.T) {
	storage := InMemoryStorage{}
	storage.SetSize(6)
	_, err := storage.Leave(7)
	if err != ErrSlotExceedsAvailableParking {
		t.Errorf("Leave() Error got %v want %v", err, ErrSlotExceedsAvailableParking)
	}
}

func TestInMemoryStorage_RegNumForCarsWithColor(t *testing.T) {
	storage := InMemoryStorage{}
	storage.SetSize(6)
	_ = storage.Park(2, &Car{
		Color:              "White",
		RegistrationNumber: "KA-01-HH-1234",
	})
	_ = storage.Park(2, &Car{
		Color:              "Red",
		RegistrationNumber: "KA-01-HH-1235",
	})
	_ = storage.Park(3, &Car{
		Color:              "White",
		RegistrationNumber: "KA-01-HH-1236",
	})
	regNums := storage.RegNumForCarsWithColor("White")
	expected := []string{"KA-01-HH-1234", "KA-01-HH-1236"}
	if !reflect.DeepEqual(regNums, expected) {
		t.Errorf("() RegNumForCarsWithColor got %v want %v", regNums, expected)
	}
}

func TestInMemoryStorage_RegNumForCarsWithColorNoCar(t *testing.T) {
	storage := InMemoryStorage{}
	storage.SetSize(6)
	regNums := storage.RegNumForCarsWithColor("White")
	expected := []string{}
	if !reflect.DeepEqual(regNums, expected) {
		t.Errorf("() RegNumForCarsWithColor got %v want %v", regNums, expected)
	}
}

func TestInMemoryStorage_SlotNumForCarsWithColor(t *testing.T) {
	storage := InMemoryStorage{}
	storage.SetSize(6)
	_ = storage.Park(1, &Car{
		Color:              "White",
		RegistrationNumber: "KA-01-HH-1234",
	})
	_ = storage.Park(2, &Car{
		Color:              "Red",
		RegistrationNumber: "KA-01-HH-1235",
	})
	_ = storage.Park(3, &Car{
		Color:              "White",
		RegistrationNumber: "KA-01-HH-1236",
	})
	regNums := storage.SlotNumForCarsWithColor("White")
	expected := []int{1, 3}
	if !reflect.DeepEqual(regNums, expected) {
		t.Errorf("() RegNumForCarsWithColor got %v want %v", regNums, expected)
	}
}

func TestInMemoryStorage_SlotNumForCarsWithColorNoCar(t *testing.T) {
	storage := InMemoryStorage{}
	storage.SetSize(6)
	regNums := storage.SlotNumForCarsWithColor("White")
	expected := []int{}
	if !reflect.DeepEqual(regNums, expected) {
		t.Errorf("() RegNumForCarsWithColor got %v want %v", regNums, expected)
	}
}

func TestInMemoryStorage_SlotNumForCarWithRegNum(t *testing.T) {
	storage := InMemoryStorage{}
	storage.SetSize(6)
	_ = storage.Park(2, &Car{
		Color:              "White",
		RegistrationNumber: "KA-01-HH-1234",
	})
	regNum := storage.SlotNumForCarWithRegNum("KA-01-HH-1234")
	expected := 2
	if !reflect.DeepEqual(regNum, expected) {
		t.Errorf("() RegNumForCarsWithColor got %v want %v", regNum, expected)
	}
}

func TestInMemoryStorage_SlotNumForCarWithRegNumNoCar(t *testing.T) {
	storage := InMemoryStorage{}
	storage.SetSize(6)
	regNum := storage.SlotNumForCarWithRegNum("KA-01-HH-1234")
	expected := 0
	if !reflect.DeepEqual(regNum, expected) {
		t.Errorf("() RegNumForCarsWithColor got %v want %v", regNum, expected)
	}
}

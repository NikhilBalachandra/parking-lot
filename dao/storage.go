package dao

import "errors"

var (
	ErrSlotExceedsAvailableParking = errors.New("ERR_SLOT_EXCEEDS_AVAILABLE_PARKING")
	ErrSlotAlreadyOccupied = errors.New("ERR_SLOT_ALREADY_OCCUPIED")
	ErrSlotNotOccupied = errors.New("ERR_SLOT_NOT_OCCUPIED")
	// ErrDuplicateRegNum specifies error that is returned when second car
	// with the same registration number is attempted to be parked.
	ErrDuplicateRegNum = errors.New("ERR_DUPLICATE_REG_NUM")
)

type Status struct {
	SlotNum int
	RegNum string
	Color string
}

// Car is a container struct to hold car details.
type Car struct {
	RegistrationNumber string
	Color              string
}

// Slot is a container struct to hold a car.
type Slot struct {
	ID  int
	Car *Car
}

// Storage interface deals with storing parking related information.
type Storage interface {
	// SetSize allocates and initializes memory
	SetSize(int)
	// Park parks a car. Parking a car occupies a slot.
	Park(slotID int, car *Car) error
	// Leave Un-parks a car. Un-parking a car unoccupies a slot.
	Leave(slotID int) (*Car, error)
	// RegNumForCarsWithColor returns list of cars reg numbers with the color
	RegNumForCarsWithColor(color string) []string
	// SlotNumForCarsWithColor returns list of slot numbers with the car of specified color
	SlotNumForCarsWithColor(color string) []int
	// SlotNumForCarWithRegNum returns slot ID of the car with the specified reg num.
	SlotNumForCarWithRegNum(color string) int
	// Returns status of the each occupied and unoccupied slot.
	Status() []Status
}

package dao

import (
	"parking_lot/common"
)

// index is a helper struct to index specific attributes
// such as colors.
type index struct {
	idx map[string]*common.LinkedHashIntSet
}

func newIndex() index {
	i := index{}
	i.idx = make(map[string]*common.LinkedHashIntSet)
	return i
}

func (i *index) Add(slotID int, key string) error {
	set := common.NewLinkedHashIntSet()
	if _, ok := i.idx[key]; !ok {
		i.idx[key] = set
	}
	return i.idx[key].Add(slotID)
}

func (i *index) Remove(slotID int, key string) error {
	if _, ok := i.idx[key]; ok {
		return i.idx[key].Remove(slotID)
	}
	return common.ErrSetMemberNotExists
}

func (i *index) Exists(key string) bool {
	_, ok := i.idx[key]
	return ok
}

func (i *index) Membership(key string) []int {
	if _, ok := i.idx[key]; !ok {
		i.idx[key] = common.NewLinkedHashIntSet()
		return []int{}
	}
	result := make([]int, 0)
	for _, k := range i.idx[key].Members() {
		result = append(result, k)
	}
	return result
}

type InMemoryStorage struct {
	size          int
	slots         []Slot
	slotsByColor  index // Color-SlotID mapping
	slotsByRegNum index // Registration number - SlotID mapping
}

func (ims *InMemoryStorage) SetSize(size int) {
	ims.slots = make([]Slot, size, size)
	for i := 0; i < size; i++ {
		ims.slots[i] = Slot{
			ID:  i+1,
			Car: nil,
		}
	}
	ims.size = size
	ims.slotsByColor = newIndex()
	ims.slotsByRegNum = newIndex()
}

func (ims *InMemoryStorage) Park(slotID int, car *Car) error {
	if slotID > ims.size {
		return ErrSlotExceedsAvailableParking
	}

	if ims.slots[slotID-1].Car != nil {
		return ErrSlotAlreadyOccupied
	}

	if ims.slotsByRegNum.Exists(car.RegistrationNumber) && len(ims.slotsByRegNum.Membership(car.RegistrationNumber)) > 0 {
		return ErrDuplicateRegNum
	}

	ims.slots[slotID-1].Car = car
	ims.slotsByColor.Add(slotID, car.Color)
	ims.slotsByRegNum.Add(slotID, car.RegistrationNumber)
	return nil
}

func (ims *InMemoryStorage) Leave(slotID int) (*Car, error) {
	if slotID > ims.size {
		return nil, ErrSlotExceedsAvailableParking
	}

	car := ims.slots[slotID-1].Car
	if car == nil {
		return nil, ErrSlotNotOccupied
	}

	ims.slots[slotID-1].Car = nil
	ims.slotsByColor.Remove(slotID, car.Color)
	ims.slotsByRegNum.Remove(slotID, car.RegistrationNumber)
	return car, nil
}

func (ims *InMemoryStorage) RegNumForCarsWithColor(color string) []string {
	slotIDs := ims.slotsByColor.Membership(color)
	regNums := make([]string, 0)
	for _, id := range slotIDs {
		regNums = append(regNums, ims.slots[id - 1].Car.RegistrationNumber)
	}
	return regNums
}

func (ims *InMemoryStorage) SlotNumForCarsWithColor(color string) []int {
	return ims.slotsByColor.Membership(color)
}

func (ims *InMemoryStorage) SlotNumForCarWithRegNum(regNum string) int {
	slotNums := ims.slotsByRegNum.Membership(regNum)
	if len(slotNums) <= 0 {
		return 0
	}
	return slotNums[0]
}

func (ims *InMemoryStorage) Status() []Status {
	result := make([]Status, 0, ims.size)
	for i := 0; i < ims.size; i++ {
		slot := ims.slots[i]
		car := slot.Car
		if car != nil {
			result = append(result, Status{
				SlotNum: slot.ID,
				RegNum:  car.RegistrationNumber,
				Color:   car.Color,
			})
		} else {
			result = append(result, Status{
				SlotNum: slot.ID,
				RegNum:  "",
				Color:   "",
			})
		}
	}
	return result
}
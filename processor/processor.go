package processor

import (
	"errors"
	"fmt"
	"parking_lot/dao"
	"parking_lot/parser"
	"strconv"
	"strings"
	"sync"
)

var (
	// ErrParkingLotSizeInvalid specifies <= 0 size parking lot
	ErrParkingLotSizeInvalid = errors.New("ERR_PARKING_LOT_SIZE_INVALID")
	// ErrParkingLotSizeAlreadySet specifies size of parking lot is already set.
	ErrParkingLotSizeAlreadySet = errors.New("ERR_PARKING_LOT_SIZE_ALREADY_SET")
	// ErrParkingLotSizeNotSet specifies size of parking lot is not set.
	ErrParkingLotSizeNotSet = errors.New("ERR_PARKING_LOT_SIZE_NOT_SET")
	// ErrInvalidSlotID specifies slot ID is either not valid or is out of parking lot bounds.
	ErrInvalidSlotID = errors.New("ERR_INVALID_SLOT_ID")
)

func Process(tokenizer *parser.Tokenizer, mutex *sync.Mutex, allocator Allocator, s dao.Storage) (string, error) {
	command, err := parser.NextCommand(tokenizer)

	if err != nil {
		return "", err
	}

	// Use mutex. Though not necessarily required in this particular example
	// as stdin I suppose writing to stdin wouldn't be concurrent (due to interleaving)
	// Adding it here as in real-world, a parking lot might have multiple entry and exit
	// point which may lead to concurrent access.
	mutex.Lock()
	defer mutex.Unlock()

	// IMPORTANT:
	// In the switch statement below, number of arguments is already
	// validated by parser. Argument slice indices can be safely
	// used without additional bound checks.
	switch command.Type {
	case parser.CommandCreateParkingLot:
		size, err := strconv.ParseInt(command.Arguments[0], 10, 64)
		if err != nil {
			return "", ErrInvalidSlotID
		} else if size <= 0 {
			return "", ErrParkingLotSizeInvalid
		} else if allocator.GetSize() > 0 {
			return "", ErrParkingLotSizeAlreadySet
		}

		allocator.SetSize(int(size))
		s.SetSize(int(size))
		return fmt.Sprintf("Created a parking lot with %d slots\n", size), nil
	case parser.CommandPark:
		if allocator.GetSize() <= 0 {
			return "", ErrParkingLotSizeNotSet
		}
		car := dao.Car{
			RegistrationNumber: command.Arguments[0],
			Color:              command.Arguments[1],
		}
		slotID := allocator.SelectCandidate()
		if slotID == 0 {
			return "Sorry, parking lot is full\n", nil
		}
		err = s.Park(slotID, &car)
		if err != nil {
			return "", err
		}
		allocator.MarkAsAllocated()
		return fmt.Sprintf("Allocated slot number: %d\n", slotID), nil
	case parser.CommandLeave:
		if allocator.GetSize() <= 0 {
			return "", ErrParkingLotSizeNotSet
		}
		slotID, err := strconv.ParseInt(command.Arguments[0], 10, 64)
		if err != nil {
			return "", ErrInvalidSlotID
		}
		_, err = s.Leave(int(slotID))
		if err != nil {
			return "", err
		}
		allocator.MarkAsAvailable(int(slotID))
		return fmt.Sprintf("Slot number %d is free\n", slotID), nil
	case parser.CommandStatus:
		return Format(s.Status()), nil
	case parser.CommandRegNumForCarWithColor:
		regNums := s.RegNumForCarsWithColor(command.Arguments[0])
		if len(regNums) <= 0 {
			return fmt.Sprint("Not found\n"), nil
		}
		return fmt.Sprintf("%s\n", strings.Join(regNums, ", ")), nil
	case parser.CommandSlotNumForCarWithColor:
		slots := s.SlotNumForCarsWithColor(command.Arguments[0])
		if len(slots) <= 0 {
			return fmt.Sprint("Not found\n"), nil
		}
		return fmt.Sprintf("%s\n", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(slots)), ", "), "[]")), nil
	case parser.CommandSlotNumForCarWithRegNum:
		slotID := s.SlotNumForCarWithRegNum(command.Arguments[0])
		if slotID == 0 {
			return fmt.Sprint("Not found\n"), nil
		}
		return fmt.Sprintf("%d\n", slotID), nil
	case parser.CommandUnknown:
		return "", parser.ErrUnknownCommand
	default:
		panic(fmt.Sprintf("Unhandled command type %v", command.Type))
	}
}

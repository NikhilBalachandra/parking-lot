package parser

import (
	"bytes"
	"errors"
	"strings"
)

var (
	// ErrEmptyLineEntry specifies empty line in the input.
	ErrEmptyLineEntry = errors.New("ERR_LINE_STRING_EMPTY")
	ErrUnknownCommand = errors.New("ERR_UNKNOWN_COMMAND")
	ErrIncorrectUsage = errors.New("ERR_INCORRECT_USAGE")
)

type CommandType int

const (
	CommandUnknown CommandType = iota
	CommandCreateParkingLot
	CommandPark
	CommandLeave
	CommandStatus
	CommandRegNumForCarWithColor
	CommandSlotNumForCarWithColor
	CommandSlotNumForCarWithRegNum
)

const (
	// AverageArgumentsPerCommand is the average number of arguments in a
	// single command. Used to eagerly allocate memory.
	AverageArgumentsPerCommand = 3
)

// Type represents a single line from the input.
type Command struct {
	Type      CommandType
	Arguments []string
}

func NewCommand(command CommandType, args []string) Command {
	return Command{
		Type:      command,
		Arguments: args,
	}
}

func NextCommand(t *Tokenizer) (Command, error) {
	token, err := t.NextToken()
	if err != nil {
		return NewCommand(CommandUnknown, nil), err
	} else if len(token) == 0 {
		return NewCommand(CommandUnknown, nil), ErrEmptyLineEntry
	}

	// Split token into command and it's arguments.
	// Copy the values from the token buffer. This is required as underlying
	// memory could be erased / reused.
	cmdAndArgs := bytes.Split(token, []byte(" "))
	cmd := string(cmdAndArgs[0])
	args := make([]string, 0, AverageArgumentsPerCommand)
	for _, arg := range cmdAndArgs[1:] {
		args = append(args, string(arg))
	}

	switch cmd {
	case "create_parking_lot":
		return parseCommandCreateParkingLot(args)
	case "park":
		return parseCommandPark(args)
	case "leave":
		return parseCommandLeave(args)
	case "status":
		return parseCommandStatus(args)
	case "registration_numbers_for_cars_with_colour":
		return parseCommandRegNumForCarWithColor(args)
	case "slot_numbers_for_cars_with_colour":
		return parseCommandSlotNumForCarWithColor(args)
	case "slot_number_for_registration_number":
		return parseCommandSlotNumForCarWithRegNum(args)
	default:
		return NewCommand(CommandUnknown, args), ErrUnknownCommand
	}
}

// parseCommandCreateParkingLot contains logic to parse create_parking_lot command.
// Example: "create_parking_lot 6"
func parseCommandCreateParkingLot(args []string) (Command, error) {
	if len(args) != 1 {
		return NewCommand(CommandCreateParkingLot, args), ErrIncorrectUsage
	}
	return NewCommand(CommandCreateParkingLot, args), nil
}

// parseCommandPark contains logic to parse park command.
// Examples:
//   1) "park KA-01-HH-1234 White"
//   2) "park KA-01-HH-1234 Crimson Red"
func parseCommandPark(args []string) (Command, error) {
	if len(args) < 2 {
		return NewCommand(CommandPark, args), ErrIncorrectUsage
	}

	// Join Colors separated with space (Eg: "Crimson Red") into single argument.
	color := strings.Join(args[1:], " ")
	return NewCommand(CommandPark, []string{args[0], color}), nil
}

// parseCommandLeave contains logic to parse leave command.
// Example: "leave 4"
func parseCommandLeave(args []string) (Command, error) {
	if len(args) != 1 {
		return NewCommand(CommandLeave, args), ErrIncorrectUsage
	}
	return NewCommand(CommandLeave, args), nil
}

// parseCommandStatus contains logic to parse status command.
// Example: "status"
func parseCommandStatus(args []string) (Command, error) {
	if len(args) != 0 {
		return NewCommand(CommandStatus, args), ErrIncorrectUsage
	}
	return NewCommand(CommandStatus, nil), nil
}

// parseCommandRegNumForCarWithColor contains logic to parse registration_numbers_for_cars_with_colour command.
// Examples:
//   1) "registration_numbers_for_cars_with_colour White"
//   2) "registration_numbers_for_cars_with_colour Crimson Red"
func parseCommandRegNumForCarWithColor(args []string) (Command, error) {
	if len(args) < 1 {
		return NewCommand(CommandRegNumForCarWithColor, args), ErrIncorrectUsage
	}
	// Join Colors separated with space (Eg: "Crimson Red") into single argument.
	color := strings.Join(args[0:], " ")
	return NewCommand(CommandRegNumForCarWithColor, []string{color}), nil
}

// parseCommandSlotNumForCarWithColor contains logic to parse slot_numbers_for_cars_with_colour command.
// Examples:
//   1) "slot_numbers_for_cars_with_colour White"
//   2) "slot_numbers_for_cars_with_colour Crimson Red"
func parseCommandSlotNumForCarWithColor(args []string) (Command, error) {
	if len(args) < 1 {
		return NewCommand(CommandSlotNumForCarWithColor, args), ErrIncorrectUsage
	}
	// Join Colors separated with space (Eg: "Crimson Red") into single argument.
	color := strings.Join(args[0:], " ")
	return NewCommand(CommandSlotNumForCarWithColor, []string{color}), nil
}

// parseCommandSlotNumForCarWithRegNum contains logic to parse slot_number_for_registration_number command.
// Example: "slot_number_for_registration_number KA-01-HH-1234"
func parseCommandSlotNumForCarWithRegNum(args []string) (Command, error) {
	if len(args) != 1 {
		return NewCommand(CommandSlotNumForCarWithRegNum, args), ErrIncorrectUsage
	}
	return NewCommand(CommandSlotNumForCarWithRegNum, args), nil
}

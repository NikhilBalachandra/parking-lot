package parser

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

//func TestNewCommandFromBufferCopiesUnderlyingMemory(t *testing.T) {
//	b := []byte("park KA-01-HH-1234 White")
//	cmd, _ := NewCommandFromBuffer(b)
//
//	if cmd.Name != "park" {
//		t.Errorf("NewCommandFromBuffer() got = %v, want %v", cmd.Name, "park")
//	}
//
//	// Modify underlying memory
//	b[0] = 10
//	b[1] = 20
//
//	expectedArguments := []string{"KA-01-HH-1234", "White"}
//	if !reflect.DeepEqual(cmd.Arguments, expectedArguments) {
//		t.Errorf("NewCommandFromBuffer() got = %v, want %v", cmd.Arguments, expectedArguments)
//	}
//}

func TestNextCommand(t *testing.T) {
	tests := []struct {
		name        string
		tokenizer   Tokenizer
		want        Command
		wantErr     bool
		wantErrType error
	}{
		{
			name: "Returns io.EOF on EOF", tokenizer: NewTokenizer(strings.NewReader("")),
			want: NewCommand(CommandUnknown, nil), wantErr: true, wantErrType: io.EOF,
		},
		{
			name: "Returns ErrEmptyLineEntry on just new line", tokenizer: NewTokenizer(strings.NewReader("\n")),
			want: NewCommand(CommandUnknown, nil), wantErr: true, wantErrType: ErrEmptyLineEntry,
		},
		{
			name: "Parse create_parking_lot", tokenizer: NewTokenizer(strings.NewReader("create_parking_lot 6\n")),
			want: NewCommand(CommandCreateParkingLot, []string{"6"}), wantErr: false,
		},
		{
			name: "Fails create_parking_lot without arg", tokenizer: NewTokenizer(strings.NewReader("create_parking_lot\n")),
			want: NewCommand(CommandCreateParkingLot, []string{}), wantErr: true, wantErrType: ErrIncorrectUsage,
		},
		{
			name: "Fails create_parking_lot with more args", tokenizer: NewTokenizer(strings.NewReader("create_parking_lot 5 6\n")),
			want: NewCommand(CommandCreateParkingLot, []string{"5", "6"}), wantErr: true, wantErrType: ErrIncorrectUsage,
		},
		{
			name: "Parse park", tokenizer: NewTokenizer(strings.NewReader("park KA-01-HH-1234 White\n")),
			want: NewCommand(CommandPark, []string{"KA-01-HH-1234", "White"}), wantErr: false,
		},
		{
			name: "Parse park with two world color", tokenizer: NewTokenizer(strings.NewReader("park KA-01-HH-1234 Crimson Red\n")),
			want: NewCommand(CommandPark, []string{"KA-01-HH-1234", "Crimson Red"}), wantErr: false,
		},
		{
			name: "Fails park without arg", tokenizer: NewTokenizer(strings.NewReader("park\n")),
			want: NewCommand(CommandPark, []string{}), wantErr: true, wantErrType: ErrIncorrectUsage,
		},
		{
			name: "Fails park without two arg", tokenizer: NewTokenizer(strings.NewReader("park KA-01-HH-1234\n")),
			want: NewCommand(CommandPark, []string{"KA-01-HH-1234"}), wantErr: true, wantErrType: ErrIncorrectUsage,
		},
		{
			name: "Parse status", tokenizer: NewTokenizer(strings.NewReader("status\n")),
			want: NewCommand(CommandStatus, nil), wantErr: false,
		},
		{
			name: "Fail status with arg", tokenizer: NewTokenizer(strings.NewReader("status 4\n")),
			want: NewCommand(CommandStatus, []string{"4"}), wantErr: true, wantErrType: ErrIncorrectUsage,
		},
		{
			name: "Parse registration_numbers_for_cars_with_colour", tokenizer: NewTokenizer(strings.NewReader("registration_numbers_for_cars_with_colour White\n")),
			want: NewCommand(CommandRegNumForCarWithColor, []string{"White"}), wantErr: false,
		},
		{
			name: "Parse registration_numbers_for_cars_with_colour with two word color", tokenizer: NewTokenizer(strings.NewReader("registration_numbers_for_cars_with_colour Dark Blue\n")),
			want: NewCommand(CommandRegNumForCarWithColor, []string{"Dark Blue"}), wantErr: false,
		},
		{
			name: "Fail registration_numbers_for_cars_with_colour without arg", tokenizer: NewTokenizer(strings.NewReader("registration_numbers_for_cars_with_colour\n")),
			want: NewCommand(CommandRegNumForCarWithColor, []string{}), wantErr: true, wantErrType: ErrIncorrectUsage,
		},
		{
			name: "Parse slot_numbers_for_cars_with_colour", tokenizer: NewTokenizer(strings.NewReader("slot_numbers_for_cars_with_colour White\n")),
			want: NewCommand(CommandSlotNumForCarWithColor, []string{"White"}), wantErr: false,
		},
		{
			name: "Parse slot_numbers_for_cars_with_colour with two word color", tokenizer: NewTokenizer(strings.NewReader("slot_numbers_for_cars_with_colour Royal Blue\n")),
			want: NewCommand(CommandSlotNumForCarWithColor, []string{"Royal Blue"}), wantErr: false,
		},
		{
			name: "Fail slot_numbers_for_cars_with_colour without arg", tokenizer: NewTokenizer(strings.NewReader("slot_numbers_for_cars_with_colour\n")),
			want: NewCommand(CommandSlotNumForCarWithColor, []string{}), wantErr: true, wantErrType: ErrIncorrectUsage,
		},
		{
			name: "Parse slot_number_for_registration_number", tokenizer: NewTokenizer(strings.NewReader("slot_number_for_registration_number KA-01-HH-1234\n")),
			want: NewCommand(CommandSlotNumForCarWithRegNum, []string{"KA-01-HH-1234"}), wantErr: false,
		},
		{
			name: "Fail slot_number_for_registration_number with two arg", tokenizer: NewTokenizer(strings.NewReader("slot_number_for_registration_number KA-01-HH-1234 KA-01-HH-1235\n")),
			want: NewCommand(CommandSlotNumForCarWithRegNum, []string{"KA-01-HH-1234", "KA-01-HH-1235"}), wantErr: true, wantErrType: ErrIncorrectUsage,
		},
		{
			name: "Fail slot_number_for_registration_number without arg", tokenizer: NewTokenizer(strings.NewReader("slot_number_for_registration_number\n")),
			want: NewCommand(CommandSlotNumForCarWithRegNum, []string{}), wantErr: true, wantErrType: ErrIncorrectUsage,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NextCommand(&tt.tokenizer)
			if (err != nil) != tt.wantErr {
				t.Errorf("NextCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && err != nil && err != tt.wantErrType {
				t.Errorf("NextCommand() got = %v, want %v", err, tt.wantErrType)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NextCommand() got = %v, want %v", got, tt.want)
			}
		})
	}
}

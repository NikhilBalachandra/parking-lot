package processor

import (
	"fmt"
	"parking_lot/dao"
	"strings"
)

func Format(status []dao.Status) string {
	builder := strings.Builder{}
	builder.WriteString("Slot No.    Registration No    Colour\n")
	for _, entry := range status {
		if entry.RegNum == "" || entry.Color == "" {
			continue
		}
		builder.WriteString(fmt.Sprintf("%-11d %-18s %s\n", entry.SlotNum, entry.RegNum, entry.Color))
	}
	return builder.String()
}

package formatType

import (
	"fmt"
)

func Indicator(s string) (string, bool) {
	if s[0:4] != "0002" {
		fmt.Println("Not Thai QR Payment")
		return "", false
	} else {
		fmt.Printf("\n%s : %s %s %s\n", "[Payload Format Indicator]", s[0:2], s[2:4], s[4:6])

		fmt.Printf("%-15s: %s\n", "Tag ID", s[0:2])
		fmt.Printf("%-15s: %s\n", "Length", s[2:4])
		fmt.Printf("%-15s: %s\n", "Value", s[4:6])
		fmt.Printf("%-15s: %s\n", "Description", "Release version of QR")
		return s[6:], true
	}
}

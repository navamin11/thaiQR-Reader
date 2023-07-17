package formatType

import (
	"fmt"
)

func InitiationMethod(s string) string {
	fmt.Printf("\n%s : %s %s %s\n", "[Point of Initiation method]", s[0:2], s[2:4], s[4:6])
	fmt.Printf("%-15s: %s\n", "Tag ID", s[0:2])
	fmt.Printf("%-15s: %s\n", "Length", s[2:4])
	fmt.Printf("%-15s: %s\n", "Value", indicatesQR(s[4:6]))
	fmt.Printf("%-15s: %s\n", "Description", "Initiation method (Type of QR)")

	return s[6:]
}

func indicatesQR(code string) string {
	var des string
	switch code {
	case "11":
		des = "Static QR Payment"
	case "12":
		des = "Dynamic QR Payment"
	default:
		des = "Unknow"
	}

	return fmt.Sprintf("%s [%s]", code, des)
}

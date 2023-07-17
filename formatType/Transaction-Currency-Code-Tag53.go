package formatType

import (
	"fmt"
	"strconv"
)

func TransactionCurrencyCodeTag53(s string) string {
	if s[0:2] != "53" {
		return s
	}

	v, _ := strconv.Atoi(s[2:4])

	fmt.Printf("\n%s : %s %s %s\n", "[Transaction Currency Code]", s[0:2], s[2:4], s[4:(4+v)])
	fmt.Printf("%-15s: %s\n", "Tag ID", s[0:2])
	fmt.Printf("%-15s: %s\n", "Length", s[2:4])
	fmt.Printf("%-15s: %s\n", "Value", s[4:(4+v)])
	fmt.Printf("%-15s: %s\n", "Description", "The transaction currency shall contain 3-digit numeric representation of the currency as defined by ISO 4217")

	return s[(4 + v):]
}

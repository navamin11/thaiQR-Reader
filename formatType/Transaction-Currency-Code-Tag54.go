package formatType

import (
	"fmt"
	"strconv"
)

func TransactionAmountTag54(s string) string {
	if s[0:2] != "54" {
		return s
	}

	v, _ := strconv.Atoi(s[2:4])

	fmt.Printf("\n%s : %s %s %s\n", "[Transaction Amount]", s[0:2], s[2:4], s[4:(4+v)])
	fmt.Printf("%-15s: %s\n", "Tag ID", s[0:2])
	fmt.Printf("%-15s: %s\n", "Length", s[2:4])
	fmt.Printf("%-15s: %s\n", "Value", s[4:(4+v)])
	fmt.Printf("%-15s: %s\n", "Description", "Amount will be provided with decimal sign which can be parsed to decimal type value")

	return s[(4 + v):]
}

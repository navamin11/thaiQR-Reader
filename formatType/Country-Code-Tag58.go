package formatType

import (
	"fmt"
	"strconv"
)

func CountryCodeTag58(s string) string {
	if s[0:2] != "58" {
		return s
	}

	v, _ := strconv.Atoi(s[2:4])

	fmt.Printf("\n%s : %s %s %s\n", "[Country Code]", s[0:2], s[2:4], s[4:(4+v)])
	fmt.Printf("%-15s: %s\n", "Tag ID", s[0:2])
	fmt.Printf("%-15s: %s\n", "Length", s[2:4])
	fmt.Printf("%-15s: %s\n", "Value", s[4:(4+v)])
	fmt.Printf("%-15s: %s\n", "Description", "Country Code should indicate the country in which the merchant transacts. It should have a value as defined by ISO 3166-1 alpha 2")

	return s[(4 + v):]
}

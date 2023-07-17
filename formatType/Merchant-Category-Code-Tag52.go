package formatType

import (
	"fmt"
	"strconv"
)

func MerchantCategoryCodeTag52(s string) string {
	if s[0:2] != "52" {
		return s
	}

	v, _ := strconv.Atoi(s[2:4])

	fmt.Printf("\n%s : %s %s %s\n", "[Merchant Category Code]", s[0:2], s[2:4], s[4:(4+v)])
	fmt.Printf("%-15s: %s\n", "Tag ID", s[0:2])
	fmt.Printf("%-15s: %s\n", "Length", s[2:4])
	fmt.Printf("%-15s: %s\n", "Value", s[4:(4+v)])
	fmt.Printf("%-15s: %s\n", "Description", "This field should indicate Merchant Category Code, MCC, as defined by ISO 18245:2003.	If the field is not necessary, use the value “0000”.")

	return s[(4 + v):]
}

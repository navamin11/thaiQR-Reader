package formatType

import (
	"fmt"
	"strconv"
)

func AdditionalDataFieldTemplateTag62(s string) string {
	if s[0:2] != "62" {
		return s
	}

	v, _ := strconv.Atoi(s[2:4])

	fmt.Printf("\n%s : %s %s %s\n", "[Postal Code]", s[0:2], s[2:4], s[4:(4+v)])
	fmt.Printf("%-15s: %s\n", "Tag ID", s[0:2])
	fmt.Printf("%-15s: %s\n", "Length", s[2:4])
	fmt.Printf("%-15s: %s\n", "Value", s[4:(4+v)])
	fmt.Printf("%-15s: %s\n", "Description", "")

	return s[(4 + v):]
}

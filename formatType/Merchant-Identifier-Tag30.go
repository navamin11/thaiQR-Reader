package formatType

import (
	"fmt"
	"strconv"
)

func MerchantIdentifierPromptPayTag30(s string) string {
	if s[0:2] != "30" {
		return s
	}

	v, _ := strconv.Atoi(s[2:4])

	fmt.Printf("\n%s : %s %s %s\n", "[Merchant identifier - Reserved for PromptPay - Bill Payment]", s[0:2], s[2:4], s[4:(4+v)])
	fmt.Printf("%-15s: %s\n", "Tag ID", s[0:2])
	fmt.Printf("%-15s: %s\n", "Length", s[2:4])
	fmt.Printf("%-15s: %s\n", "Value", s[4:(4+v)])

	v_AID, _ := strconv.Atoi(s[6:8])

	fmt.Printf("\n%-15s: %s\n", "Sub Tag ID", s[4:6])
	fmt.Printf("%-15s: %s\n", "Length", s[6:8])
	fmt.Printf("%-15s: %s\n", "Value", s[8:(8+v_AID)])
	fmt.Printf("%-15s: %s - %s\n", "Description", "AID", checkAIDTag30(s[8:(8+v_AID)]))

	newString := s[(8 + v_AID):]

	v_PPBillerID, _ := strconv.Atoi(newString[2:4])
	des, remark := checkPromptPayIDTag30(newString[0:2])

	fmt.Printf("\n%-15s: %s [%s]\n", "Sub Tag ID", newString[0:2], des)
	fmt.Printf("%-15s: %s\n", "Length", newString[2:4])
	fmt.Printf("%-15s: %s\n", "Value", newString[4:(4+v_PPBillerID)])
	fmt.Printf("%-15s: %s\n", "Description", remark)

	newString = newString[(4 + v_PPBillerID):]

	v_PPReference1, _ := strconv.Atoi(newString[2:4])
	des, remark = checkPromptPayIDTag30(newString[0:2])

	fmt.Printf("\n%-15s: %s [%s]\n", "Sub Tag ID", newString[0:2], des)
	fmt.Printf("%-15s: %s\n", "Length", newString[2:4])
	fmt.Printf("%-15s: %s\n", "Value", newString[4:(4+v_PPReference1)])
	fmt.Printf("%-15s: %s\n", "Description", remark)

	newString = newString[(4 + v_PPReference1):]

	if newString[0:2] != "03" {
		return newString
	}

	v_PPReference2, _ := strconv.Atoi(newString[2:4])
	des, remark = checkPromptPayIDTag30(newString[0:2])

	fmt.Printf("\n%-15s: %s [%s]\n", "Sub Tag ID", newString[0:2], des)
	fmt.Printf("%-15s: %s\n", "Length", newString[2:4])
	fmt.Printf("%-15s: %s\n", "Value", newString[4:(4+v_PPReference2)])
	fmt.Printf("%-15s: %s\n", "Description", remark)

	return newString[(4 + v_PPReference2):]
}

func checkAIDTag30(code string) string {
	var des string
	switch code {
	case "A000000677010112", "a000000677010112":
		des = "domestic merchant"
	case "A000000677012006", "a000000677012006":
		des = "cross-border merchant"
	default:
		des = "Unknow"
	}

	return fmt.Sprintf("%s [%s]", code, des)
}

func checkPromptPayIDTag30(code string) (string, string) {
	var des, remark string
	switch code {
	case "01":
		des = "Biller ID"
		remark = "National ID/Tax ID + Suffix"
	case "02":
		des = "Reference 1"
	case "03":
		des = "Reference 2"
	default:
		des = "Unknow"
	}
	return des, remark
}

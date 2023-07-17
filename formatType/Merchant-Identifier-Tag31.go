package formatType

import (
	"fmt"
	"strconv"
)

func MerchantIdentifierPromptPayTag31(s string) string {
	if s[0:2] != "31" {
		return s
	}

	v, _ := strconv.Atoi(s[2:4])

	fmt.Printf("\n%s : %s %s %s\n", "[Merchant identifier - Reserved for Payment Innovation (API) - 2 use cases]", s[0:2], s[2:4], s[4:(4+v)])
	fmt.Printf("%-15s: %s\n", "Tag ID", s[0:2])
	fmt.Printf("%-15s: %s\n", "Length", s[2:4])
	fmt.Printf("%-15s: %s\n", "Value", s[4:(4+v)])

	v_AID, _ := strconv.Atoi(s[6:8])

	fmt.Printf("\n%-15s: %s\n", "Sub Tag ID", s[4:6])
	fmt.Printf("%-15s: %s\n", "Length", s[6:8])
	fmt.Printf("%-15s: %s\n", "Value", s[8:(8+v_AID)])
	fmt.Printf("%-15s: %s - %s\n", "Description", "AID", checkAIDTag31(s[8:(8+v_AID)]))

	newString := s[(8 + v_AID):]

	if s[8:(8+v_AID)] == "A000000677012004" || s[8:(8+v_AID)] == "a000000677012004" {

		v_PPAPIID, _ := strconv.Atoi(newString[2:4])
		des, remark := checkPromptPayIDTag31(newString[0:2])

		fmt.Printf("\n%-15s: %s [%s]\n", "Sub Tag ID", newString[0:2], des)
		fmt.Printf("%-15s: %s\n", "Length", newString[2:4])
		fmt.Printf("%-15s: %s\n", "Value", newString[4:(4+v_PPAPIID)])
		fmt.Printf("%-15s: %s\n", "Description", remark)

		newString = newString[(4 + v_PPAPIID):]

		v_PPServiceProviderID, _ := strconv.Atoi(newString[2:4])
		des, remark = checkPromptPayIDTag31(newString[0:2])

		fmt.Printf("\n%-15s: %s [%s]\n", "Sub Tag ID", newString[0:2], des)
		fmt.Printf("%-15s: %s\n", "Length", newString[2:4])
		fmt.Printf("%-15s: %s\n", "Value", newString[4:(4+v_PPServiceProviderID)])
		fmt.Printf("%-15s: %s\n", "Description", remark)

		newString = newString[(4 + v_PPServiceProviderID):]

		v_PPTransactionReference, _ := strconv.Atoi(newString[2:4])
		des, remark = checkPromptPayIDTag31(newString[0:2])

		fmt.Printf("\n%-15s: %s [%s]\n", "Sub Tag ID", newString[0:2], des)
		fmt.Printf("%-15s: %s\n", "Length", newString[2:4])
		fmt.Printf("%-15s: %s\n", "Value", newString[4:(4+v_PPTransactionReference)])
		fmt.Printf("%-15s: %s\n", "Description", remark)

		return newString[(4 + v_PPTransactionReference):]
	} else {
		v_PPAcquirerID, _ := strconv.Atoi(newString[2:4])
		des, remark := checkPromptPayIDTag31(newString[0:2])

		fmt.Printf("\n%-15s: %s [%s]\n", "Sub Tag ID", newString[0:2], des)
		fmt.Printf("%-15s: %s\n", "Length", newString[2:4])
		fmt.Printf("%-15s: %s\n", "Value", newString[4:(4+v_PPAcquirerID)])
		fmt.Printf("%-15s: %s\n", "Description", remark)

		newString = newString[(4 + v_PPAcquirerID):]
		if newString[0:2] != "02" {
			return newString
		}

		v_PPAcquirerspecific, _ := strconv.Atoi(newString[2:4])
		des, remark = checkPromptPayIDTag31(newString[0:2])

		fmt.Printf("\n%-15s: %s [%s]\n", "Sub Tag ID", newString[0:2], des)
		fmt.Printf("%-15s: %s\n", "Length", newString[2:4])
		fmt.Printf("%-15s: %s\n", "Value", newString[4:(4+v_PPAcquirerspecific)])
		fmt.Printf("%-15s: %s\n", "Description", remark)

		return newString[(4 + v_PPAcquirerspecific):]
	}
}

func checkAIDTag31(code string) string {
	var des string
	switch code {
	case "A000000677012004", "a000000677012004":
		des = "Payment Innovation [Standard API]"
	case "A000000677010113", "a000000677010113":
		des = "Payment Innovation [Acquirer specific innovation]"
	default:
		des = "Unknow"
	}

	return fmt.Sprintf("%s [%s]", code, des)
}

func checkPromptPayIDTag31(code string) (string, string) {
	var des, remark string
	switch code {
	case "01":
		des = "API ID"
		remark = "\"001\"=Transaction verification API Note: The following sub tags will vary based on this API ID"
	case "02":
		des = "Service Provider ID"
	case "03":
		des = "Transaction Reference"
	default:
		des = "Unknow"
	}
	return des, remark
}

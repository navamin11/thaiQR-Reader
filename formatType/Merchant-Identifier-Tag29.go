package formatType

import (
	"fmt"
	"strconv"
)

func MerchantIdentifierPromptPayTag29(s string) string {
	if s[0:2] != "29" {
		return s
	}

	v, _ := strconv.Atoi(s[2:4])

	fmt.Printf("\n%s : %s %s %s\n", "[Merchant identifier - Reserved for PromptPay - Credit Transfer with PromptPayID]", s[0:2], s[2:4], s[4:(4+v)])
	fmt.Printf("%-15s: %s\n", "Tag ID", s[0:2])
	fmt.Printf("%-15s: %s\n", "Length", s[2:4])
	fmt.Printf("%-15s: %s\n", "Value", s[4:(4+v)])

	v_AID, _ := strconv.Atoi(s[6:8])

	fmt.Printf("\n%-15s: %s\n", "Sub Tag ID", s[4:6])
	fmt.Printf("%-15s: %s\n", "Length", s[6:8])
	fmt.Printf("%-15s: %s\n", "Value", s[8:(8+v_AID)])
	fmt.Printf("%-15s: %s - %s\n", "Description", "AID", checkAIDTag29(s[8:(8+v_AID)]))

	newString := s[(8 + v_AID):]

	v_PPID, _ := strconv.Atoi(newString[2:4])
	des, remark := checkPromptPayIDTag29(newString[0:2], s[8:(8+v_AID)])

	fmt.Printf("\n%-15s: %s [%s]\n", "Sub Tag ID", newString[0:2], des)
	fmt.Printf("%-15s: %s\n", "Length", newString[2:4])
	fmt.Printf("%-15s: %s\n", "Value", newString[4:(4+v_PPID)])
	fmt.Printf("%-15s: %s\n", "Description", remark)

	return newString[(4 + v_PPID):]
}

func checkAIDTag29(code string) string {
	var des string
	switch code {
	case "A000000677010111":
		des = "merchant-presented QR"
	case "A000000677010114":
		des = "customer-presented QR"
	default:
		des = "Unknow"
	}

	return fmt.Sprintf("%s [%s]", code, des)
}

func checkPromptPayIDTag29(code, aid string) (string, string) {
	var des, remark string
	switch code {
	case "01":
		des = "Mobile Number"
		remark = "e.g. 0066XXXXXXXXX"
	case "02":
		des = "National ID or Tax ID"
	case "03":
		des = "E-Wallet ID"
	case "04":
		des = "Bank Account"
		remark = "Reserved for future use; Bank code (3 digit) + account no."
	case "05":
		if aid == "A000000677010114" {
			des = "OTA"
			remark = "mandatory if AID = \"A000000677010114\""
		} else {
			des = "Unknow"
		}
	default:
		des = "Unknow"
	}
	return des, remark
}

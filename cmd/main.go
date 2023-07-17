package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"thaiQR/formatType"
)

func YesNoPrompt(label string, def bool) (string, bool) {
	choices := "N"
	if !def {
		choices = "n"
	}

	r := bufio.NewReader(os.Stdin)
	var s string

	for {
		fmt.Fprintf(os.Stderr, "\n%s (%s to Exit...): ", label, choices)
		s, _ = r.ReadString('\n')
		s = strings.TrimSpace(s)
		if s == "" {
			return s, def
		}
		s = strings.ToLower(s)
		if s == "n" || s == "no" {
			return s, true
		} else {
			return s, false
		}
	}
}

func main() {
	for {
		input, ok := YesNoPrompt("Enter Thai QR Code", true)
		if ok {
			fmt.Println("Good bye!")
			break
		} else {
			input = strings.TrimSuffix(input, "\n")
			input, valid := formatType.Indicator(input)
			if valid {
				input = formatType.InitiationMethod(input)
				input = formatType.MerchantIdentifierVisa(input)
				input = formatType.MerchantIdentifierMastercard(input)
				input = formatType.MerchantIdentifierPromptPayTag29(input)
				input = formatType.MerchantIdentifierPromptPayTag30(input)
				input = formatType.MerchantIdentifierPromptPayTag31(input)
				input = formatType.MerchantCategoryCodeTag52(input)
				input = formatType.TransactionCurrencyCodeTag53(input)
				input = formatType.TransactionAmountTag54(input)
				input = formatType.CountryCodeTag58(input)
				input = formatType.MerchantNameTag59(input)
				input = formatType.MerchantCityTag60(input)
				input = formatType.PostalCodeTag61(input)
				input = formatType.AdditionalDataFieldTemplateTag62(input)
				formatType.CRCTag63(input)
			}
		}
	}
}

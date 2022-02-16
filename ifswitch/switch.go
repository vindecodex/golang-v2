package ifswitch

import "strings"

const (
	AUSTRALIA   string = "AU"
	BELGIUM            = "BE"
	CANADA             = "CA"
	DENMARK            = "DK"
	EGYPT              = "EG"
	FINLAND            = "FI"
	GERMANY            = "DE"
	HONGKONG           = "HK"
	INDIA              = "IN"
	JAPAN              = "JP"
	MADAGASCAR         = "MG"
	NETHERLANDS        = "NL"
	PHILIPPINES        = "PH"
	RUSSIA             = "RU"
	SINGAPORE          = "SG"
	THAILAND           = "TH"
	UKRAINE            = "UA"
	VIETNAM            = "VN"
)

func Countries2LetterCode(country string) string {
	switch strings.ToLower(country) {
	case "australia":
		return AUSTRALIA
	case "belgium":
		return BELGIUM
	case "canada":
		return CANADA
	case "denmark":
		return DENMARK
	case "egypt":
		return EGYPT
	case "finland":
		return FINLAND
	case "germany":
		return GERMANY
	case "hongkong":
		return HONGKONG
	case "india":
		return INDIA
	case "japan":
		return JAPAN
	case "madagascar":
		return MADAGASCAR
	case "netherlands":
		return NETHERLANDS
	case "philippines":
		return PHILIPPINES
	case "russia":
		return RUSSIA
	case "singapore":
		return SINGAPORE
	case "thailand":
		return THAILAND
	case "ukraine":
		return UKRAINE
	case "vietnam":
		return VIETNAM
	default:
		return country + " not listed"
	}
}

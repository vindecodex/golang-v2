package ifswitch

import "strings"

const (
	MALE   = "male"
	FEMALE = "female"
)

func SheOrHe(gender string) string {
	if strings.ToLower(gender) == MALE {
		return "He"
	}
	if strings.ToLower(gender) == FEMALE {
		return "She"
	}
	return "Unknown gender"
}

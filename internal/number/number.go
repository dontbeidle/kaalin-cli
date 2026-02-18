package number

import (
	"fmt"
	"math"
	"strings"
)

var onesLat = []string{"nol", "bir", "eki", "úsh", "tórt", "bes", "altı", "jeti", "segiz", "toǵız"}
var teensLat = []string{"on bir", "on eki", "on úsh", "on tórt", "on bes", "on altı", "on jeti", "on segiz", "on toǵız"}
var tensLat = []string{"on", "jigirma", "otız", "qırıq", "eliw", "alpıs", "jetpis", "seksen", "toqsan"}
var thousandsLat = []string{"", "mıń", "million", "milliard", "trillion", "kvadrillion", "kvintillion", "sekstilion", "septillion", "oktillion", "nonillion"}

const hundredLat = "júz"
const minusLat = "minus"
const putinLat = "pútin"

var onesCyr = []string{"ноль", "бир", "еки", "үш", "төрт", "бес", "алты", "жети", "сегиз", "тоғыз"}
var teensCyr = []string{"он бир", "он еки", "он үш", "он төрт", "он бес", "он алты", "он жети", "он сегиз", "он тоғыз"}
var tensCyr = []string{"он", "жигирма", "отыз", "қырық", "елиў", "алпыс", "жетпис", "сексен", "тоқсан"}
var thousandsCyr = []string{"", "мың", "миллион", "миллиард", "триллион", "квадриллион", "квинтиллион", "секстиллион", "септиллион", "октиллион", "нониллион"}

const hundredCyr = "жүз"
const minusCyr = "минус"
const putinCyr = "пүтін"

// Maximum value: 10^30
const maxValue = 1e30

// ToWord converts a number to its Karakalpak word representation.
// script should be "lat" (default) or "cyr".
func ToWord(number float64, script string) (string, error) {
	if math.IsInf(number, 0) || math.IsNaN(number) {
		return "", fmt.Errorf("\"%.0f\" is not a valid number", number)
	}

	negative := false
	if number < 0 {
		negative = true
		number = -number
	}

	if number >= maxValue {
		return "", fmt.Errorf("number exceeds maximum allowed value (max: 10^30)")
	}

	isCyr := script == "cyr"

	// Split into integer and fractional parts
	intPart := int64(number)
	fracStr := extractFraction(number)

	var result string
	if fracStr != "" {
		// Fractional number
		intWord := convertInteger(intPart, isCyr)
		fracWord, denomWord := convertFraction(fracStr, isCyr)
		putin := putinLat
		if isCyr {
			putin = putinCyr
		}
		result = intWord + " " + putin + " " + denomWord + " " + fracWord
	} else {
		result = convertInteger(intPart, isCyr)
	}

	if negative {
		prefix := minusLat
		if isCyr {
			prefix = minusCyr
		}
		result = prefix + " " + result
	}

	return result, nil
}

func extractFraction(number float64) string {
	s := fmt.Sprintf("%.10f", number)
	parts := strings.SplitN(s, ".", 2)
	if len(parts) < 2 {
		return ""
	}
	frac := strings.TrimRight(parts[1], "0")
	return frac
}

func convertInteger(n int64, isCyr bool) string {
	ones := onesLat
	teens := teensLat
	tens := tensLat
	thousands := thousandsLat
	hundred := hundredLat
	if isCyr {
		ones = onesCyr
		teens = teensCyr
		tens = tensCyr
		thousands = thousandsCyr
		hundred = hundredCyr
	}

	if n == 0 {
		return ones[0]
	}

	// Special case: exactly 100 → "júz" / "жүз"
	if n == 100 {
		return hundred
	}

	// Special case: exactly 1000 → "mıń" / "мың"
	if n == 1000 {
		return thousands[1]
	}

	var parts []string
	groupIndex := 0

	for n > 0 {
		group := int(n % 1000)
		n /= 1000

		if group != 0 {
			groupWord := convertHundreds(group, ones, teens, tens, hundred)

			if groupIndex > 0 {
				suffix := thousands[groupIndex]
				// Special case: only for thousands (groupIndex==1), group==1 → just "mıń"
				if group == 1 && groupIndex == 1 {
					groupWord = suffix
				} else {
					groupWord = groupWord + " " + suffix
				}
			}

			parts = append([]string{groupWord}, parts...)
		}

		groupIndex++
		if groupIndex >= len(thousands) {
			break
		}
	}

	return strings.Join(parts, " ")
}

func convertHundreds(n int, ones, teens, tens []string, hundred string) string {
	if n == 0 {
		return ""
	}

	var parts []string

	h := n / 100
	remainder := n % 100

	if h > 0 {
		parts = append(parts, ones[h]+" "+hundred)
	}

	if remainder > 0 {
		if remainder <= 9 {
			parts = append(parts, ones[remainder])
		} else if remainder == 10 {
			parts = append(parts, tens[0])
		} else if remainder >= 11 && remainder <= 19 {
			parts = append(parts, teens[remainder-11])
		} else {
			t := remainder / 10
			u := remainder % 10
			if u == 0 {
				parts = append(parts, tens[t-1])
			} else {
				parts = append(parts, tens[t-1]+" "+ones[u])
			}
		}
	}

	return strings.Join(parts, " ")
}

func convertFraction(fracStr string, isCyr bool) (string, string) {
	// Parse the fractional digits as an integer
	fracNum := int64(0)
	for _, c := range fracStr {
		fracNum = fracNum*10 + int64(c-'0')
	}

	// Denominator is 10^len(fracStr)
	denomPow := len(fracStr)
	denomWord := getDenominator(denomPow, isCyr)

	fracWord := convertInteger(fracNum, isCyr)
	return fracWord, denomWord
}

func getDenominator(power int, isCyr bool) string {
	if isCyr {
		return getDenominatorCyr(power)
	}
	return getDenominatorLat(power)
}

func getDenominatorLat(power int) string {
	switch power {
	case 1:
		return "onnan"
	case 2:
		return "júzden"
	case 3:
		return "mıńnan"
	case 4:
		return "on mıńnan"
	case 5:
		return "júz mıńnan"
	case 6:
		return "millionnan"
	case 7:
		return "on millionnan"
	case 8:
		return "júz millionnan"
	case 9:
		return "milliardtan"
	default:
		return "onnan"
	}
}

func getDenominatorCyr(power int) string {
	switch power {
	case 1:
		return "оннан"
	case 2:
		return "жүзден"
	case 3:
		return "мыңнан"
	case 4:
		return "он мыңнан"
	case 5:
		return "жүз мыңнан"
	case 6:
		return "миллионнан"
	case 7:
		return "он миллионнан"
	case 8:
		return "жүз миллионнан"
	case 9:
		return "миллиардтан"
	default:
		return "оннан"
	}
}

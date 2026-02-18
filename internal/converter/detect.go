package converter

import "unicode"

// DetectScript analyzes text and returns "cyrillic", "latin", or "unknown".
func DetectScript(text string) string {
	var cyrCount, latCount int

	for _, r := range text {
		if isCyrillic(r) {
			cyrCount++
		} else if isLatin(r) {
			latCount++
		}
	}

	if cyrCount == 0 && latCount == 0 {
		return "unknown"
	}
	if cyrCount >= latCount {
		return "cyrillic"
	}
	return "latin"
}

func isCyrillic(r rune) bool {
	return unicode.Is(unicode.Cyrillic, r)
}

func isLatin(r rune) bool {
	if unicode.Is(unicode.Latin, r) {
		return true
	}
	// Additional Karakalpak Latin characters with diacritics
	switch r {
	case 'á', 'Á', 'ó', 'Ó', 'ú', 'Ú', 'ǵ', 'Ǵ', 'ń', 'Ń', 'ı', 'Í':
		return true
	}
	return false
}

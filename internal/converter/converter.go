package converter

import (
	"strings"
	"unicode/utf8"
)

// Latin2Cyrillic converts Karakalpak Latin text to Cyrillic.
func Latin2Cyrillic(text string) string {
	var b strings.Builder
	b.Grow(len(text))

	runes := []rune(text)
	i := 0

	for i < len(runes) {
		// Try 2-character combinations first (uppercase)
		if i+1 < len(runes) {
			pair := string(runes[i : i+2])

			found := false
			for _, m := range latToCyrMultiUpper {
				if pair == m.Latin {
					b.WriteString(m.Cyrillic)
					i += 2
					found = true
					break
				}
			}
			if found {
				continue
			}

			for _, m := range latToCyrMultiLower {
				if pair == m.Latin {
					b.WriteString(m.Cyrillic)
					i += 2
					found = true
					break
				}
			}
			if found {
				continue
			}
		}

		// Single character mapping
		r := runes[i]
		if repl, ok := latToCyrUpper[r]; ok {
			b.WriteString(repl)
		} else if repl, ok := latToCyrLower[r]; ok {
			b.WriteString(repl)
		} else {
			b.WriteRune(r)
		}
		i++
	}

	return b.String()
}

// Cyrillic2Latin converts Karakalpak Cyrillic text to Latin.
func Cyrillic2Latin(text string) string {
	// Apply special rules before conversion (only when NOT at word start)
	text = applySpecialRules(text)

	var b strings.Builder
	b.Grow(len(text))

	for _, r := range text {
		if repl, ok := cyrToLatUpper[r]; ok {
			b.WriteString(repl)
		} else if repl, ok := cyrToLatLower[r]; ok {
			b.WriteString(repl)
		} else {
			b.WriteRune(r)
		}
	}

	return b.String()
}

// applySpecialRules handles ьи→yi, ьо→yo, ъе→ye transformations.
// These rules apply only when NOT at word start.
func applySpecialRules(text string) string {
	runes := []rune(text)
	var b strings.Builder
	b.Grow(len(text))

	i := 0
	for i < len(runes) {
		if i > 0 && i+1 < len(runes) && !isWordBoundary(runes[i-1]) {
			pair := string(runes[i : i+2])
			switch pair {
			case "ьи":
				b.WriteString("yi")
				i += 2
				continue
			case "ьо":
				b.WriteString("yo")
				i += 2
				continue
			case "ъе":
				b.WriteString("ye")
				i += 2
				continue
			}
		}

		// Check if current rune has mapping (pass through for special rules at word start)
		b.WriteRune(runes[i])
		i++
	}

	return b.String()
}

func isWordBoundary(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r' ||
		!utf8.ValidRune(r)
}

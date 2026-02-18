package strutil

import "strings"

// Upper converts text to uppercase with Karakalpak-specific handling.
// Converts dotless-i (ı) to Í before applying standard uppercase.
func Upper(text string) string {
	text = strings.ReplaceAll(text, "ı", "Í")
	return strings.ToUpper(text)
}

// Lower converts text to lowercase with Karakalpak-specific handling.
// Converts Í to dotless-i (ı) before applying standard lowercase.
func Lower(text string) string {
	text = strings.ReplaceAll(text, "Í", "ı")
	return strings.ToLower(text)
}

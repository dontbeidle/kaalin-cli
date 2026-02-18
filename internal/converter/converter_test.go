package converter

import "testing"

func TestLatin2Cyrillic(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"basic greeting", "Assalawma áleykum", "Ассалаўма әлейкум"},
		{"hello", "Sálem", "Сәлем"},
		{"sh digraph", "sháyir", "шәйир"},
		{"ch digraph", "chaqqan", "чаққан"},
		{"ya digraph", "yaman", "яман"},
		{"yu digraph", "yurt", "юрт"},
		{"yo digraph", "yoq", "ёқ"},
		{"uppercase SH", "Shahar", "Шаҳар"},
		{"uppercase CH", "Chet", "Чет"},
		{"uppercase YA", "Yaman", "Яман"},
		{"uppercase YU", "Yurt", "Юрт"},
		{"numbers passthrough", "abc 123 def", "абц 123 деф"},
		{"empty string", "", ""},
		{"punctuation passthrough", "Sálem, dun'ya!", "Сәлем, дун'я!"},
		{"mixed content", "Hello 123 Sálem!", "Ҳелло 123 Сәлем!"},
		{"all caps SH", "SHAHAR", "ШАҲАР"},
		{"special chars", "ǵarri", "ғарри"},
		{"dotless i", "qırıq", "қырық"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Latin2Cyrillic(tt.input)
			if got != tt.want {
				t.Errorf("Latin2Cyrillic(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestCyrillic2Latin(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"basic greeting", "Ассалаўма әлейкум", "Assalawma áleykum"},
		{"hello", "Сәлем", "Sálem"},
		{"ш to sh", "шәйир", "sháyir"},
		{"ч to ch", "чаққан", "chaqqan"},
		{"я to ya", "яман", "yaman"},
		{"ю to yu", "юрт", "yurt"},
		{"ё to yo", "ёқ", "yoq"},
		{"uppercase", "СӘЛЕМ", "SÁLEM"},
		{"numbers passthrough", "абц 123 деф", "abc 123 def"},
		{"empty string", "", ""},
		{"punctuation", "Сәлем, дуньа!", "Sálem, duna!"},
		{"ъ removed", "объект", "obyekt"},
		{"ь removed", "кальций", "kalciy"},
		{"щ to sh", "Щ", "Sh"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Cyrillic2Latin(tt.input)
			if got != tt.want {
				t.Errorf("Cyrillic2Latin(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestSpecialRules(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"ьи mid-word", "обьиктив", "obyiktiv"},
		{"ьо mid-word", "обьор", "obyor"},
		{"ъе mid-word", "объект", "obyekt"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Cyrillic2Latin(tt.input)
			if got != tt.want {
				t.Errorf("Cyrillic2Latin(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestMultiCharMapping(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"sh", "ш"},
		{"ch", "ч"},
		{"ya", "я"},
		{"yu", "ю"},
		{"Sh", "Ш"},
		{"Ch", "Ч"},
		{"Ya", "Я"},
		{"Yu", "Ю"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := Latin2Cyrillic(tt.input)
			if got != tt.want {
				t.Errorf("Latin2Cyrillic(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestPassthrough(t *testing.T) {
	inputs := []string{
		"123",
		"!@#$%^&*()",
		"   ",
		"\n\t",
	}

	for _, input := range inputs {
		t.Run(input, func(t *testing.T) {
			if got := Latin2Cyrillic(input); got != input {
				t.Errorf("Latin2Cyrillic(%q) = %q, want passthrough", input, got)
			}
			if got := Cyrillic2Latin(input); got != input {
				t.Errorf("Cyrillic2Latin(%q) = %q, want passthrough", input, got)
			}
		})
	}
}

func TestDetectScript(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"cyrillic text", "Сәлем", "cyrillic"},
		{"latin text", "Sálem", "latin"},
		{"numbers only", "12345", "unknown"},
		{"empty string", "", "unknown"},
		{"mixed more cyrillic", "Сәлем hello", "cyrillic"},
		{"mixed more latin", "Сa hello world", "latin"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DetectScript(tt.input)
			if got != tt.want {
				t.Errorf("DetectScript(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

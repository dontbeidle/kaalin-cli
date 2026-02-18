package strutil

import "testing"

func TestUpper(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"basic", "assalawma áleykum", "ASSALAWMA ÁLEYKUM"},
		{"with dotless i", "qırıq", "QÍRÍQ"},
		{"empty", "", ""},
		{"already upper", "SÁLEM", "SÁLEM"},
		{"mixed", "Sálem Dun'ya", "SÁLEM DUN'YA"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Upper(tt.input)
			if got != tt.want {
				t.Errorf("Upper(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestLower(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"basic", "ASSALAWMA ÁLEYKUM", "assalawma áleykum"},
		{"with Í", "QÍRÍK", "qırık"},
		{"empty", "", ""},
		{"already lower", "sálem", "sálem"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Lower(tt.input)
			if got != tt.want {
				t.Errorf("Lower(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestDotlessI(t *testing.T) {
	// ı (dotless i) should become Í when uppercased
	got := Upper("ı")
	if got != "Í" {
		t.Errorf("Upper(ı) = %q, want Í", got)
	}

	// Í should become ı (dotless i) when lowercased
	got = Lower("Í")
	if got != "ı" {
		t.Errorf("Lower(Í) = %q, want ı", got)
	}
}

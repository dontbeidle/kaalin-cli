package number

import "testing"

func TestBasicNumbers(t *testing.T) {
	tests := []struct {
		name   string
		input  float64
		script string
		want   string
	}{
		{"zero", 0, "lat", "nol"},
		{"one", 1, "lat", "bir"},
		{"five", 5, "lat", "bes"},
		{"ten", 10, "lat", "on"},
		{"eleven", 11, "lat", "on bir"},
		{"nineteen", 19, "lat", "on toǵız"},
		{"twenty", 20, "lat", "jigirma"},
		{"hundred", 100, "lat", "júz"},
		{"thousand", 1000, "lat", "mıń"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToWord(tt.input, tt.script)
			if err != nil {
				t.Fatalf("ToWord(%v, %q) returned error: %v", tt.input, tt.script, err)
			}
			if got != tt.want {
				t.Errorf("ToWord(%v, %q) = %q, want %q", tt.input, tt.script, got, tt.want)
			}
		})
	}
}

func TestComplexNumbers(t *testing.T) {
	tests := []struct {
		name   string
		input  float64
		script string
		want   string
	}{
		{"123", 123, "lat", "bir júz jigirma úsh"},
		{"999", 999, "lat", "toǵız júz toqsan toǵız"},
		{"1001", 1001, "lat", "mıń bir"},
		{"1100", 1100, "lat", "mıń bir júz"},
		{"2000", 2000, "lat", "eki mıń"},
		{"10000", 10000, "lat", "on mıń"},
		{"1000000", 1000000, "lat", "bir million"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToWord(tt.input, tt.script)
			if err != nil {
				t.Fatalf("ToWord(%v, %q) returned error: %v", tt.input, tt.script, err)
			}
			if got != tt.want {
				t.Errorf("ToWord(%v, %q) = %q, want %q", tt.input, tt.script, got, tt.want)
			}
		})
	}
}

func TestFloat(t *testing.T) {
	got, err := ToWord(12.75, "lat")
	if err != nil {
		t.Fatalf("ToWord(12.75, lat) returned error: %v", err)
	}
	want := "on eki pútin júzden jetpis bes"
	if got != want {
		t.Errorf("ToWord(12.75, lat) = %q, want %q", got, want)
	}
}

func TestNegative(t *testing.T) {
	got, err := ToWord(-5, "lat")
	if err != nil {
		t.Fatalf("ToWord(-5, lat) returned error: %v", err)
	}
	want := "minus bes"
	if got != want {
		t.Errorf("ToWord(-5, lat) = %q, want %q", got, want)
	}
}

func TestCyrillic(t *testing.T) {
	tests := []struct {
		name  string
		input float64
		want  string
	}{
		{"zero cyr", 0, "ноль"},
		{"123 cyr", 123, "бир жүз жигирма үш"},
		{"999 cyr", 999, "тоғыз жүз тоқсан тоғыз"},
		{"1000 cyr", 1000, "мың"},
		{"100 cyr", 100, "жүз"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToWord(tt.input, "cyr")
			if err != nil {
				t.Fatalf("ToWord(%v, cyr) returned error: %v", tt.input, err)
			}
			if got != tt.want {
				t.Errorf("ToWord(%v, cyr) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestOverflow(t *testing.T) {
	_, err := ToWord(1e31, "lat")
	if err == nil {
		t.Error("ToWord(1e31) should return error for overflow")
	}
}

func TestSpecialCases(t *testing.T) {
	tests := []struct {
		name  string
		input float64
		want  string
	}{
		{"100 is júz", 100, "júz"},
		{"1000 is mıń", 1000, "mıń"},
		{"200", 200, "eki júz"},
		{"300", 300, "úsh júz"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToWord(tt.input, "lat")
			if err != nil {
				t.Fatalf("ToWord(%v, lat) returned error: %v", tt.input, err)
			}
			if got != tt.want {
				t.Errorf("ToWord(%v, lat) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

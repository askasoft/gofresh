package fresh

import (
	"testing"
)

func TestParseDate(t *testing.T) {
	tests := []struct {
		s string
		w string
	}{
		{"2020-01-02", "2020-01-02"},
		{"2022-02-02", "2022-02-02"},
		{"2023-03-02", "2023-03-02"},
	}

	for i, tt := range tests {
		ta, err := ParseDate(tt.s)
		if err != nil {
			t.Fatalf("[%d] ParseDate(%q) = %v", i, tt.s, err)
		}

		sa := ta.String()
		if sa != tt.w {
			t.Fatalf("[%d] ParseDate(%q) = %v, want %q", i, tt.s, sa, tt.w)
		}
	}
}

func TestParseTime(t *testing.T) {
	tests := []struct {
		s string
		w string
	}{
		{"2020-01-02T03:04:05Z", "2020-01-02T03:04:05Z"},
		{"2020-01-02T03:04:05+08:00", "2020-01-01T19:04:05Z"},
	}

	for i, tt := range tests {
		ta, err := ParseTime(tt.s)
		if err != nil {
			t.Fatalf("[%d] ParseTime(%q) = %v", i, tt.s, err)
		}

		sa := ta.String()
		if sa != tt.w {
			t.Fatalf("[%d] ParseTime(%q) = %v, want %q", i, tt.s, sa, tt.w)
		}
	}
}

func TestParseTimeSpent(t *testing.T) {
	cs := []struct {
		s string
		w TimeSpent
	}{
		{"09:00", 540},
		{"08:00", 480},
		{"360", 360},
	}

	for i, c := range cs {
		a, err := ParseTimeSpent(c.s)
		if err != nil || a != c.w {
			t.Errorf("[%d] ParseTimeSpent(%q) = (%d, %v), want %d", i, c.s, a, err, c.w)
		}
	}
}

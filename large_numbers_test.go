// +build !386

package numerizer

import "testing"

func TestParseLarge(t *testing.T) {
	var tests = []struct {
		in   string
		want int
	}{
		{"five billion", 5000000000},
		{"six trillion", 6000000000000},
	}
	for _, tt := range tests {
		have, err := Parse(tt.in)
		if err != nil {
			t.Fatalf("Parse(%q) %v", tt.in, err)
		} else if have != tt.want {
			t.Errorf("Parse(%q)\nhave %d\nwant %d", tt.in, have, tt.want)
		}
	}
}

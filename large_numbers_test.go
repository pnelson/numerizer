// +build !386

package numerizer

import "testing"

func TestParseLarge(t *testing.T) {
	var tests = []struct {
		in   string
		want int
	}{
		{"six billion", 6000000000},
		{"seven trillion", 7000000000000},
		{"eight quadrillion", 8000000000000000},
		{"nine quintillion two hundred twenty-three quadrillion three hundred seventy-two trillion thirty-six billion eight hundred fifty-four million seven hundred seventy-five thousand eight hundred seven", 9223372036854775807},
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

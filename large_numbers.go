// +build !386

package numerizer

// supports up to max signed 64 bit integer
// 2^63-1 = 9223372036854775807
var largeNumbers = map[string]int{
	"thousand":    1000,
	"million":     1000000,
	"billion":     1000000000,
	"trillion":    1000000000000,
	"quadrillion": 1000000000000000,
	"quintillion": 1000000000000000000,
}

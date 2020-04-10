package floatutil

import (
	"fmt"
	"github.com/spf13/cast"
	"math"
)

// RoundTo4thDecimal round off to the 4th decimal place
func RoundTo4thDecimal(f float64) float64 {
	f = math.Round(f*10000) / 10000.0
	string := fmt.Sprintf("%.4f", f)
	return cast.ToFloat64(string)
}

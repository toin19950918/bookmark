package strutil

import (
	"github.com/spf13/cast"
	"math"
	"math/rand"
	"time"
)

// const option
const (
	OptionAlphbet = 0
	OptionNumber  = 1
	OptionSymbol  = 2
)

var (
	alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	number   = []rune("0123456789")
	symbol   = []rune("!@#$%^&*")
)

// RandStr return random string
func RandStr(length int, options ...int) string {
	rand.Seed(time.Now().UnixNano())
	var letter []rune
	result := ""

	if len(options) == 0 {
		return result
	}

	for _, option := range options {
		if option == OptionAlphbet {
			letter = append(letter, alphabet...)
		} else if option == OptionNumber {
			letter = append(letter, number...)
		} else if option == OptionSymbol {
			letter = append(letter, symbol...)
		}
	}

	for i := 0; i < length; i++ {
		result += string(letter[rand.Intn(len(letter))])
	}
	return result
}

func Reverse(a []interface{}) {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
}

func StringInSlice(str string, slice []string) bool {
	for _, value := range slice {
		if str == value {
			return true
		}
	}
	return false
}

func AddZero(num int) string {
	if num < 10 {
		return "0" + cast.ToString(num)
	}
	return cast.ToString(num)
}

func FixPrecision(origin int, digit int) float64 {
	return cast.ToFloat64(origin) * math.Pow10(-1*digit)
}

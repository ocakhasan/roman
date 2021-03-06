// package roman includes all functions which are used to
// calculate roman numeral of an integer and
// find all the roman numerals in a range

package roman

import (
	"strconv"
	"sync"

	"github.com/ocakhasan/roman/pkg/handler/structure"
)

var (
	nums    = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	symbols = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
)

// ConvertIntegerToRoman converts given input
// to roman numeral.
func ConvertIntegerToRoman(input int) string {
	var (
		i      = len(nums) - 1
		result string
	)

	for input > 0 {
		division := input / nums[i]
		input = input % nums[i]

		for division > 0 {
			result += symbols[i]
			division = division - 1
		}

		i = i - 1
	}

	return result
}

// NumeralRange returns all the roman numerals of
// the numbers between minValue and maxValue.
// It handles all the operations in parallel to make it faster.
// NumeralRange will have maxValue-minValue+1 goroutines to calculate
// the roman numeral conversions.
func NumeralRange(minValue, maxValue int) []structure.RomanResponse {
	var (
		length = maxValue - minValue + 1
		output = make([]structure.RomanResponse, length)
		wg     sync.WaitGroup
	)

	wg.Add(length)

	for i := range output {
		go func(_number int, _i int) { // calculate and append to result in background
			output[_i] = structure.RomanResponse{
				Input:  strconv.Itoa(_number),
				Output: ConvertIntegerToRoman(_number),
			}

			wg.Done()
		}(minValue+i, i)
	}

	wg.Wait() // wait all of goroutines to stop.

	return output
}

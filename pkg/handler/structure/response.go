// package structure defines the structures used in
// http requests and responses

package structure

// RomanResponse is used to return the roman numeral
// conversion of an integer
type RomanResponse struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

// RangeResponse is used to return the roman numerals
// of all integers in given range.
type RangeResponse struct {
	Conversions []RomanResponse `json:"conversions"`
}

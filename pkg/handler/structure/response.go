package structure

type RomanResponse struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

type RangeResponse struct {
	Conversions []RomanResponse `json:"conversions"`
}

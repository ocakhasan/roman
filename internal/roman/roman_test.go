package roman

import (
	"reflect"
	"testing"

	"github.com/ocakhasan/roman/pkg/handler/structure"
)

func TestConvertIntegerToRoman(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  string
	}{
		{
			name:  "conversion of 1",
			input: 1,
			want:  "I",
		},
		{
			name:  "conversion of 3",
			input: 3,
			want:  "III",
		},
		{
			name:  "conversion of 4",
			input: 4,
			want:  "IV",
		},
		{
			name:  "conversion of 9",
			input: 9,
			want:  "IX",
		},
		{
			name:  "conversion of 10",
			input: 10,
			want:  "X",
		},
		{
			name:  "conversion of 40",
			input: 40,
			want:  "XL",
		},
		{
			name:  "conversion of 44",
			input: 44,
			want:  "XLIV",
		},
		{
			name:  "conversion of 49",
			input: 49,
			want:  "XLIX",
		},
		{
			name:  "conversion of 50",
			input: 50,
			want:  "L",
		},
		{
			name:  "conversion of 90",
			input: 90,
			want:  "XC",
		},
		{
			name:  "conversion of 100",
			input: 100,
			want:  "C",
		},
		{
			name:  "conversion of 400",
			input: 400,
			want:  "CD",
		},
		{
			name:  "conversion of 500",
			input: 500,
			want:  "D",
		},
		{
			name:  "conversion of 900",
			input: 900,
			want:  "CM",
		},
		{
			name:  "conversion of 1000",
			input: 1000,
			want:  "M",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertIntegerToRoman(tt.input); got != tt.want {
				t.Errorf("ConvertIntegerToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumeralRange(t *testing.T) {
	type args struct {
		minValue int
		maxValue int
	}
	tests := []struct {
		name string
		args args
		want []structure.RomanResponse
	}{
		{
			name: "range from 1 to 10",
			args: args{
				minValue: 1,
				maxValue: 10,
			},
			want: []structure.RomanResponse{
				{
					Input:  "1",
					Output: "I",
				},
				{
					Input:  "2",
					Output: "II",
				},
				{
					Input:  "3",
					Output: "III",
				},
				{
					Input:  "4",
					Output: "IV",
				},
				{
					Input:  "5",
					Output: "V",
				},
				{
					Input:  "6",
					Output: "VI",
				},
				{
					Input:  "7",
					Output: "VII",
				},
				{
					Input:  "8",
					Output: "VIII",
				},
				{
					Input:  "9",
					Output: "IX",
				},
				{
					Input:  "10",
					Output: "X",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumeralRange(tt.args.minValue, tt.args.maxValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NumeralRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

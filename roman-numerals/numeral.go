package romannumerals

import "strings"

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

type RomanNumerals []RomanNumeral

var allRomanNumerals = RomanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func (r RomanNumerals) GenerateRomanArabicMap() map[string]uint16 {
	romanArabicMap := map[string]uint16{}
	for _, numeral := range r {
		romanArabicMap[numeral.Symbol] = numeral.Value
	}
	return romanArabicMap
}

var romanArabicMap = allRomanNumerals.GenerateRomanArabicMap()

func ConvertToRoman(arabic uint16) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) (total uint16) {
	for i := 0; i < len(roman); i++ {
		symbol := string(roman[i])
		value := romanArabicMap[symbol]

		if i+1 < len(roman) {
			nextSymbol := string(roman[i+1])
			nextValue := romanArabicMap[nextSymbol]

			if nextValue > value {
				value = nextValue - value
				i++
			}
		}

		total += value
	}
	return

	// for _, symbols := range windowedRoman(roman).Symbols() {
	// 	total += romanArabicMap[symbols]
	// }
	// return
}

// type windowedRoman string

// func (w windowedRoman) Symbols() (symbols []string) {
// 	for i := 0; i < len(w); i++ {
// 		if i+1 < len(w) {
// 			symbol := string([]byte{w[i], w[i+1]})
// 			_, exists := romanArabicMap[symbol]
// 			if exists {
// 				symbols = append(symbols, symbol)
// 				i++
// 				continue
// 			}
// 		}

// 		symbol := string(w[i])
// 		symbols = append(symbols, symbol)
// 	}
// 	return
// }

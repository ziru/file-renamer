package main

import (
	"fmt"
	"regexp"
)

var (
	chToArabicMapping = "零一二三四五六七八九"
	valueMapping      map[string]int
	unitMapping       map[string]int
	pattern           *regexp.Regexp
)

func init() {
	unitMapping = make(map[string]int)
	unitMapping["十"] = 10
	unitMapping["百"] = 100
	unitMapping["千"] = 1000

	valueMapping = make(map[string]int)
	idx := 0
	for _, val := range chToArabicMapping {
		valueMapping[string(val)] = idx
		idx++
	}

	pattern = regexp.MustCompile("^[零一二三四五六七八九十]?[十百千]?")
}

// ConvertChineseNumberToArabicNumber converts the input string with chinese numbers into arabic numbers.
// Note: the conversion is simple and only works with small numbers.
func ConvertChineseNumberToArabicNumber(in string) string {
	val := 0
	for {
		comp := pattern.FindString(in)
		runes := []rune(comp)
		if len(runes) == 0 {
			break
		}
		if len(runes) == 1 {
			if unit, ok := unitMapping[comp]; ok {
				if val > 0 {
					val *= unit
				} else {
					val = unit
				}
			} else {
				val += ConvertSingleComponent(comp)
			}
		} else if len(runes) == 2 {
			// ex: 三十，五百
			val += ConvertSingleComponent(string(runes[0])) * unitMapping[string(runes[1])]
		}
		in = in[len(comp):]
	}
	return fmt.Sprintf("%d", val)
}

func ConvertSingleComponent(in string) int {
	return valueMapping[in]
}

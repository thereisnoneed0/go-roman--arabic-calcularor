package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var romanNumeralMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func main() {
	fmt.Println(calculator("X  /  I V"))
	fmt.Println(calculator("X"))
	// fmt.Println(calculator("18 -  5"))

}
func calculator(expression string) interface{} {
	trimmedExpression := strings.ReplaceAll(expression, " ", "")

	if len(trimmedExpression) < 3 {
		return " wrong number of operands or operators"
	}

	operatorRegex := regexp.MustCompile(`[+\-*/]`)
	operator := operatorRegex.FindString(trimmedExpression)

	parts := strings.Split(trimmedExpression, operator)

	var operands []int
	for _, part := range parts {
		if num, err := strconv.Atoi(part); err == nil {
			operands = append(operands, num)
		} else {
			operands = append(operands, convertRomanToInteger(part))
		}
	}

	result := applyOperatorToOperands(operands, operator)

	if isRoman(trimmedExpression) {
		return convertIntegerToRoman(result)
	}

	return result
}
func applyOperatorToOperands(s []int, o string) int {
	switch o {
	case "+":
		return s[0] + s[1]
	case "-":
		return s[0] - s[1]
	case "/":
		return s[0] / s[1]
	case "*":
		return s[0] * s[1]
	}
	return 0
}

func convertRomanToInteger(s string) int {
	var result int
	for i := 0; i < len(s); {
		currentNumeralValue := romanNumeralMap[s[i:i+1]]

		if i+1 < len(s) {
			nextNumeralValue := romanNumeralMap[s[i+1:i+2]]
			if nextNumeralValue > currentNumeralValue {
				result -= currentNumeralValue
			} else {
				result += currentNumeralValue
			}
		} else {
			result += currentNumeralValue
		}
		i++
	}
	return result
}

func convertIntegerToRoman(num int) string {
	var result, prevKey string
	var remaining int
	remaining = num

	romanKeys := make([]string, 0, len(romanNumeralMap))
	for k := range romanNumeralMap {
		romanKeys = append(romanKeys, k)
	}

	sort.Slice(romanKeys, func(i, j int) bool {
		return romanNumeralMap[romanKeys[i]] < romanNumeralMap[romanKeys[j]]
	})

	for remaining > 0 {
		for _, key := range romanKeys {
			if remaining == romanNumeralMap[key] {
				result += key
				remaining -= romanNumeralMap[key]
				break
			} else if remaining < romanNumeralMap[key] {
				result += prevKey
				remaining -= romanNumeralMap[prevKey]
				break
			} else {
				prevKey = key
			}
		}
	}

	return result
}

func isRoman(s string) bool {
	romanRegex := regexp.MustCompile(`[IVXLCDM]+`)
	matched := romanRegex.FindString(s)

	if matched == "" {
		return false
	}

	return true
}

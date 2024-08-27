package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func isValidMapping(mapping map[rune]int, s string) (int, bool) {
	var numStr string
	for _, char := range s {
		numStr += strconv.Itoa(mapping[char])
	}

	if len(numStr) > 1 && numStr[0] == '0' {
		return 0, false
	}
	num, _ := strconv.Atoi(numStr)
	return num, true
}

func solveCryptarithm(letters []rune, uniqueLetters []rune, index int, usedDigits []bool, mapping map[rune]int, operands []string, operator string) bool {
	if index == len(uniqueLetters) {

		values := make([]int, len(operands))
		for i, operand := range operands {
			if value, valid := isValidMapping(mapping, operand); valid {
				values[i] = value
			} else {
				return false
			}
		}

		var result int
		switch operator {
		case "+":
			result = values[0] + values[1]
		case "-":
			result = values[0] - values[1]
		}

		return result == values[2]
	}

	for digit := 0; digit <= 9; digit++ {
		if usedDigits[digit] {
			continue
		}
		mapping[uniqueLetters[index]] = digit
		usedDigits[digit] = true
		if solveCryptarithm(letters, uniqueLetters, index+1, usedDigits, mapping, operands, operator) {
			return true
		}
		usedDigits[digit] = false
	}

	return false
}

func main() {
	input := "ABD - AD = DKL"
	input = strings.ReplaceAll(input, " ", "")

	var operator string
	if strings.Contains(input, "+") {
		operator = "+"
	} else if strings.Contains(input, "-") {
		operator = "-"
	}

	operands := strings.Split(input, operator)
	if len(operands) != 2 {
		fmt.Println("Input tidak valid")
		return
	}
	operandsAndResult := strings.Split(operands[1], "=")
	if len(operandsAndResult) != 2 {
		fmt.Println("Input tidak valid")
		return
	}
	operands = []string{operands[0], operandsAndResult[0], operandsAndResult[1]}

	lettersMap := make(map[rune]bool)
	for _, char := range input {
		if unicode.IsLetter(char) {
			lettersMap[char] = true
		}
	}
	letters := make([]rune, 0, len(lettersMap))
	for letter := range lettersMap {
		letters = append(letters, letter)
	}

	mapping := make(map[rune]int)
	usedDigits := make([]bool, 10)
	if solveCryptarithm(letters, letters, 0, usedDigits, mapping, operands, operator) {
		for i, operand := range operands {
			if i > 0 {
				fmt.Print(operator)
			}
			for _, char := range operand {
				fmt.Print(mapping[char])
			}
			if i == 1 {
				fmt.Print("=")
			}
		}
		fmt.Println()
	} else {
		fmt.Println("Tidak ada solusi yang ditemukan")
	}
}

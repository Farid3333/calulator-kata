package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romans2Arabic = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

func operation(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		return 0
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите операцию:")
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(fmt.Sprintf("Случилась ошибка при чтении текста: %v\n", err))
	}
	text = strings.TrimSpace(text)
	var values []string
	splitStrings := strings.Split(text, " ")
	for _, part := range splitStrings {
		if part != "" {
			values = append(values, part)
		}
	}

	if len(values) != 3 {
		panic("Ошибка: должна быть одна операция (+, -, *, /) с 2 операндами и между ними пробел")
	}
	num1Str, op, num2Str := values[0], values[1], values[2]

	_, isNum1Romans := romans2Arabic[num1Str]
	_, isNum2Romans := romans2Arabic[num2Str]

	_, err1 := strconv.Atoi(num1Str)
	_, err2 := strconv.Atoi(num2Str)

	if !isNum1Romans && err1 != nil || !isNum2Romans && err2 != nil {
		panic("Ошибка: операнды не являются числами")
	}

	if isNum1Romans && !isNum2Romans || !isNum1Romans && isNum2Romans {
		panic("Ошибка: нельзя смешивать римские и арабские числа.")
	}

	if op != "+" && op != "-" && op != "*" && op != "/" {
		panic("Ошибка: неизвестная операция, можно только +, -, *, /")
	}

	var num1, num2 int
	if isNum1Romans && isNum2Romans {
		num1 = romans2Arabic[num1Str]
		num2 = romans2Arabic[num2Str]
	} else {
		num1, _ = strconv.Atoi(num1Str)
		num2, _ = strconv.Atoi(num2Str)
	}

	if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
		panic("Ошибка: числа должны быть в диапазоне от 1 до 10.")
	}

	resultInt := operation(num1, num2, op)
	var resultString string
	if isNum1Romans {
		resultString = arabicToRoman(resultInt)
	} else {
		resultString = strconv.Itoa(resultInt)
	}
	fmt.Println(resultString)
}

func arabicToRoman(num int) string {
	if num < 1 {
		panic("Ошибка: в римской системе нет отрицательных чисел или нуля.")
	}

	result := ""
	values := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i, value := range values {
		for num >= value {
			num -= value
			result += symbols[i]
		}
	}
	return result
}

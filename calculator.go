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
		fmt.Printf("случалась ошибка при чтении текста; %v\n", err)
	}
	text = strings.TrimSpace(text)
	var values []string
	splitStrings := strings.Split(text, " ")
	for _, part := range splitStrings {
		if part != "" {
			values = append(values, part)
		}
	}

	num1Str, op, num2Str := values[0], values[1], values[2]

	_, isNum1Romans := romans2Arabic[num1Str]
	_, isNum2Romans := romans2Arabic[num2Str]

	_, isArabicNum1 := strconv.Atoi(num1Str)
	_, isArabicNum2 := strconv.Atoi(num2Str)

	if !isNum1Romans && isArabicNum1 != nil || !isNum2Romans && isArabicNum2 != nil {
		fmt.Println("Ошибка: операнды не являются числами")
		return
	}

	if isNum1Romans && !isNum2Romans || !isNum1Romans && isNum2Romans {
		fmt.Println("Ошибка: нельзя смешивать римские и арабские числа.")
		return
	}

	if op != "+" && op != "-" && op != "*" && op != "/" {
		fmt.Println("Ошибка: неизвестная операция, можно только +, -, *, /")
		return
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
		fmt.Println("Ошибка: числа должны быть в диапазоне от 1 до 10.")
		return
	}

	result := operation(num1, num2, op)
	if isNum1Romans {
		if result < 1 {
			fmt.Println("Ошибка: в римской системе нет отрицательных чисел или нуля.")
			return
		}
		fmt.Println(arabicToRoman(result))
	} else {
		fmt.Println(result)
	}
}

func arabicToRoman(num int) string {
	if num < 1 {
		return "Exception"
	}

	result := ""
	values := []int{10, 9, 5, 4, 1}
	symbols := []string{"X", "IX", "V", "IV", "I"}

	for i, value := range values {
		for num >= value {
			num -= value
			result += symbols[i]
		}
	}
	return result
}

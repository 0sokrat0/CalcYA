package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

// Calc вычисляет выражение, учитывая приоритет операторов и скобки.
func Calc(expression string) (float64, error) {
	index := 0
	return parseExpression(expression, &index)
}

// parseExpression обрабатывает сложение и вычитание.
func parseExpression(expression string, index *int) (float64, error) {
	result, err := parseTerm(expression, index) // Сначала обрабатываем умножение/деление
	if err != nil {
		return 0, err
	}

	for *index < len(expression) {
		char := expression[*index]

		if char == ' ' { // Пропускаем пробелы
			*index++
			continue
		}

		if char == '+' || char == '-' {
			*index++                                      // Переходим к следующему символу
			nextTerm, err := parseTerm(expression, index) // Обрабатываем следующий терм
			if err != nil {
				return 0, err
			}
			if char == '+' {
				result += nextTerm
			} else {
				result -= nextTerm
			}
		} else {
			break // Если нет '+' или '-', завершаем обработку
		}
	}

	return result, nil
}

// parseTerm обрабатывает умножение и деление.
func parseTerm(expression string, index *int) (float64, error) {
	result, err := parseFactor(expression, index) // Сначала обрабатываем числа и скобки
	if err != nil {
		return 0, err
	}

	for *index < len(expression) {
		char := expression[*index]

		if char == ' ' { // Пропускаем пробелы
			*index++
			continue
		}

		if char == '*' || char == '/' {
			*index++                                          // Переходим к следующему символу
			nextFactor, err := parseFactor(expression, index) // Обрабатываем следующий множитель
			if err != nil {
				return 0, err
			}
			if char == '*' {
				result *= nextFactor
			} else {
				if nextFactor == 0 {
					return 0, errors.New("деление на ноль")
				}
				result /= nextFactor
			}
		} else {
			break // Если нет '*' или '/', завершаем обработку
		}
	}

	return result, nil
}

// parseFactor обрабатывает числа и выражения в скобках.
func parseFactor(expression string, index *int) (float64, error) {
	if *index < len(expression) && expression[*index] == '(' {
		*index++                                          // Пропускаем '('
		result, err := parseExpression(expression, index) // Обрабатываем выражение внутри скобок
		if err != nil {
			return 0, err
		}
		if *index >= len(expression) || expression[*index] != ')' {
			return 0, errors.New("отсутствует закрывающая скобка")
		}
		*index++ // Пропускаем ')'
		return result, nil
	}

	return readNumber(expression, index) // Если это число, читаем его
}

// readNumber считывает число из строки.
func readNumber(expression string, index *int) (float64, error) {
	start := *index
	for *index < len(expression) && (unicode.IsDigit(rune(expression[*index])) || expression[*index] == '.') {
		*index++
	}
	if start == *index {
		return 0, errors.New("ожидалось число")
	}
	return strconv.ParseFloat(expression[start:*index], 64)
}

func main() {
	expression := "2+2*2"
	result, err := Calc(expression)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Результат: %f\n", result) // Ожидается: 6.000000
	}
}

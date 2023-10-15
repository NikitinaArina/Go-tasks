package tasks

import (
	"math"
	"strings"
	"unicode"
)

/*
https://leetcode.com/problems/string-to-integer-atoi/description/

Убираем пробелы в начале и конце строки
Если пустая строка, то возвращаем 0
Проверяем на знак и удаляем его из строки
Пробегаемся циклом:

	Проверяем является ли символ цифрой
	Если да, то увеличиваем нум в 10 раз и прибавляем (текущая руна - '0') '0' - позволяет преобразовать руну в число
*/
func MyAtoi(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}

	sign := 1
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	var num int
	for _, c := range s {
		if !unicode.IsDigit(c) {
			break
		}
		num = num*10 + int(c-'0')
		if sign*num <= math.MinInt32 {
			return math.MinInt32
		}
		if sign*num >= math.MaxInt32 {
			return math.MaxInt32
		}
	}

	return int(sign * num)
}

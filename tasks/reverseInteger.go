package tasks

import "math"

/*
https://leetcode.com/problems/reverse-integer/description/
Пока х != 0:

	переменную темп увеличиваем в 10 раз и прибавляем остаток от деления

	проверяем подходит ли получившееся число под ограничение если да, то возвращаем 0
	иначе делим х на 10
*/
func Reverse(x int) int {
	temp := 0

	for x != 0 {
		temp = temp*10 + x%10
		if temp > math.MaxInt32 || temp < math.MinInt32 {
			return 0
		}
		x = x / 10
	}

	return temp
}

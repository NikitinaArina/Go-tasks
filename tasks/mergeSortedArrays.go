package tasks

/*
https://leetcode.com/problems/merge-sorted-array/description/

Три итеративных переменных, две для 1-ого массива, одна для 2-ого массива
Начинаем цикл с конца массива, если значение из 1-ого массива больше, то перезаписываем его в конец массива
Иначе значение из 2-ого массива перезаписываем в конец 1-ого
*/
func Merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	k := m + n - 1

	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}

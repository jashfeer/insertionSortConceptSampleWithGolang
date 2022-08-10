package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) []int{
	for i := 0; i < n; i++ {
		nums1=append(nums1, nums2[i]) 
	}
	for i := 1; i < m+n; i++ {
		key := nums1[i]
		j := i - 1
		for j >= 0 && nums1[j] > key {
			nums1[j+1] = nums1[j]
			j--
		}
		nums1[j+1] = key
	}
	return nums1
}

func main() {
	array1 := []int{10, 30, 4, 29, 50,0,4,3,5,6, 4, 5}
	array2 := []int{20, 40, 3, 2,30,3,2,4,5,6,7,8,9,2,1,3,4, 5, 6, 7}
	fmt.Println(merge(array1, len(array1), array2, len(array2))) 

}

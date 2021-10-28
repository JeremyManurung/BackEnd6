package main

import "fmt"

func BinarySearch(array []int, x int) {
	i, j := 0, len(array)-1
	if x > array[len(array)-1] || x < array[0] {
		fmt.Println(-1)
		return
	} else if array[i] == x {
		fmt.Println(i)
		return
	} else if array[j] == x {
		fmt.Println(j)
		return
	} else {
		for i+1 != j {
			if array[(i+j)/2] == x {
				fmt.Println((i + j) / 2)
				return
			} else if array[(i+j)/2] < x {
				i = (i + j) / 2
			} else {
				j = (i + j) / 2
			}
		}
	}

	fmt.Println(-1)
	return
}

func main() {

	BinarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3) // 2

	BinarySearch([]int{1, 2, 3, 5, 6, 8, 10}, 5) // 3

	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 53) // 6

	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 100) // -1

}
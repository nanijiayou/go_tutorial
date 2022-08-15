package slice

import "fmt"

func TestSlice() {
	fmt.Println("============= slice ===============")
	// 1.
	var arr1 [10]int
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	// 2.
	arr2 := [10]int{1, 2, 3, 4}
	for index, value := range arr2 {
		fmt.Println("index = ", index, " value = ", value)
	}

	// 3. 申明一个切片， 并初始化为 1， 2， 3， 4
	arr3 := []int{1, 2, 3, 4}
	for index, value := range arr3 {
		fmt.Println("index = ", index, " value = ", value)
	}

	// 4.
	slice1 := make([]int, 3)
	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)

	// 5. 判空
	slice2 := make([]int, 0)
	if slice2 == nil {
		fmt.Println("slice2 is empty")
	} else {
		fmt.Println("slice2 is not empty")
	}

	// 操作
	// 1. len 和 cap属性（和stl vector一致)
	slice3 := make([]int, 3, 5)
	fmt.Printf("len = %d, cap = %d, slice = %v \n", len(slice3), cap(slice3), slice3)

	// 2. append
	slice3 = append(slice3, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v \n", len(slice3), cap(slice3), slice3)
	slice3 = append(slice3, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v \n", len(slice3), cap(slice3), slice3)
	slice3 = append(slice3, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v \n", len(slice3), cap(slice3), slice3)
	slice3 = append(slice3, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v \n", len(slice3), cap(slice3), slice3)
}

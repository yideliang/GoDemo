package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	var a = "abc"
	fmt.Println(a)

	_, numb, strs := number()
	fmt.Println(numb, strs)

	primenumber()
	//fmt.Println("\n")

	res := fibonacci(10)
	fmt.Println(res)

	structtojson()

	//s := []string{"flow", "flower", "flight"}
	var s = []string{"flow", "flower", "flight"}
	result := longestCommonPrefix(s)
	fmt.Println(result)
	fmt.Println("asdfads")

	var nums = []int{-2,1,-3,4,-1,2,1,-5,4}
	res1 := maxSubArray(nums)
	fmt.Println(res1)

	var nums1 = []int{3,2,3,1,2,4,5,5,6}
	k := 4
	res2 := findKthLargest(nums1, k-1)
	fmt.Println(res2)
}

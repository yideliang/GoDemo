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
}
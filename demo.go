package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string
	ID int
	Course []string
}

func structtojson(){
	group := Student{
		Name: "Dennis",
		ID: 19941024 ,
		Course: []string{"math", "science", "chinese"},
	}
	b, err := json.Marshal(group)
	if err != nil{
		fmt.Println("json.marshal failed, err:", err)
		return
	}
	fmt.Printf("%s\n", string(b))
}

func number()(int,int,string){
	var a = 1
	b, c := 2, "dennis"
	return a, b, c
}

func primenumber(){
	var i, j int
	for i=2; i<100; i++{
		for j=2; j<=(i/j); j++{
			if i%j==0 {
				break
			}
		}
		if j>(i/j) {
			fmt.Printf("%d ", i)
		}
	}
}

func fibonacci(n int) int{
	if n < 2{
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

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
}




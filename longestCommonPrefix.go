package main

import "fmt"

func main(){
	//s := []string{"flow", "flower", "flight"}
	var s = []string{"flow", "flower", "flight"}
	result := longestCommonPrefix(s)
	fmt.Println(result)
	fmt.Println("asdfads")
}

func findShortestArray(s []string) string{
	if len(s) == 0{
		return ""
	}
	shortest := s[0]
	for _, v := range s{
		if len(v) < len(shortest){
			if len(v) == 0{
				return ""
			} else {
				shortest = v
			}
		}
	}
	return shortest
}

func longestCommonPrefix(strs []string) string {
	short := findShortestArray(strs)
	if len(short) == 0{
		return ""
	}
	for i, v := range short{
		for j:=0; j<len(strs); j++{
			if strs[j][i] != byte(v){
				return strs[j][:i]
			}
		}
	}
	return short
}

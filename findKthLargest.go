package main

func findKthLargest(nums []int, k int) int{
	for i:=0; i<len(nums); i++{
		for j:=i+1; j<len(nums); j++{
			if nums[i] < nums[j]{
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums[k]
}

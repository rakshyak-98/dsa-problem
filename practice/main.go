package main

import "fmt"


func removeDuplicates(nums []int) int {
	writIndex := 0;

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i - 1]{
			nums[writIndex] = nums[i - 1]
			writIndex++
		}
	}

	nums[writIndex] = nums[len(nums) - 1]
	return writIndex  + 1
}



func main(){
	arr := []int{1,1,2}
	fmt.Println(removeDuplicates(arr))
}
package main

import "fmt"

func reverse(s []int32) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main(){
	months := [...]string{1: "January" , 12: "December"}
	fmt.Println(months)
	fmt.Println(len(months))
	fmt.Println(months[1:13])
	fmt.Println(months[:])
	//------------------------------------------------
	var s []int    // len(s) == 0, s == nil
	fmt.Println(s == nil)
	s = nil        // len(s) == 0, s == nil
	s = []int(nil) // len(s) == 0, s == nil
	s = []int{}    // len(s) == 0, s != nil
	fmt.Println(s)
	//------------------------------------------------
	nums:=[]int32{1,2,3,4,5,6,7,8,9}
	fmt.Println(nums)
	nums=append(nums, 10,11)
	fmt.Println(nums)
	reverse(nums)
	fmt.Println(nums)
	//------------------------------------------------
	nums2:=[10]int32{0,1,2,3,4,5,6,7,8,9}
	fmt.Println(nums2)
	reverse(nums2[:])
	fmt.Println("### ",nums2)
	//------------------------------------------------
	nums3:=make([]int32, 5, 6)
	fmt.Println("nums3 ",nums3)
	nums3[0]=1
	fmt.Println("nums3 ",nums3)
	nums3=append(nums3, 10,11)
	fmt.Println("nums3 ",nums3)
	fmt.Println("cap ",cap(nums3))
	fmt.Println("len ",len(nums3))
}

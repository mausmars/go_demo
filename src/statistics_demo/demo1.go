package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	probability:=[]int32{10,10,10,10,10}

	counts:=[]int32{0,0,0,0,0}
	randomCount:=10000

	//count1:=make([]int32)
	for i:=0;i<randomCount;i++{
		index:=ratioRandom(probability)
		counts[index]+=1
	}
	realProbability:=make([]float32,5)
	for i,count :=range counts{
		realProbability[i]=float32(count)/float32(randomCount)
	}
	fmt.Println(realProbability)
}

func ratioRandom(array []int32) int{
	total :=0
	for _,v := range array{
		total+=int(v)
	}
	if total <= 0 {
		return -1
	}
	rand.Seed(time.Now().UnixNano())
	n:=rand.Intn(total)
	total = -1
	for index,rate := range array{
		total += int(rate)
		if n <= total {
			return index
		}
	}
	return -1
}
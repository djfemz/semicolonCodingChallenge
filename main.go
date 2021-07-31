package main

import (
	"fmt"

)


func main(){
	output:= calculateIndex([]int{1,2,3,4,5}, 5)
	fmt.Println(output)
}

func calculateIndex(inputSlice []int, target int) (holdingArray []int)  {
	for i:=0; i<len(inputSlice);i++{
		for j := i+1; j < len(inputSlice); j++ {
				if target-inputSlice[i]==inputSlice[j]&&len(holdingArray)<1{
					holdingArray=append(holdingArray, i,j)
				}
			}
		}
		return 
}
	

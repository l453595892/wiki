package main

import (
	"math"
	"fmt"
)

type BaseData struct {
	Data []int
}

func (d *BaseData)BinarySearch(k int) int{
	left, right, mid := 1, len(d.Data), 0
	for {
		mid = int(math.Floor(float64((left + right) / 2)))
		if d.Data[mid] > k {
			right = mid - 1
		} else if d.Data[mid] < k {
			left = mid + 1
		} else {
			break
		}
		if left > right {
			mid = -1
			break
		}
	}
	return mid
}

func main(){
	d1:=BaseData{Data:[]int{1,2,3,8,5,6,7,4,9,10}}
	fmt.Println(d1.BinarySearch(7))
}



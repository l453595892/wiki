package main

import (
	"fmt"
)

func main() {
	fmt.Println(getNum(100, 1))
}

func pow(x, n int) int {
	ret := 1
	for n != 0 {
		if n%2 != 0 {
			ret = ret * x
		}
		n /= 2
		x = x * x
	}
	return ret
}

func getCount(n int64) (int64, int) {
	var count int64 = 0
	var length int
	for {
		length++
		count += n & 1
		n >>= 1
		if n == 0 {
			break
		}
	}
	return count, length
}

var slice []int64

func getNum(k, m int) int64 {
	countk, lengthk := getCount(int64(k))
	for n := k; n <= k+pow(2, lengthk-2); n++ {
		countn, lengthn := getCount(int64(n))
		if (countk == countn) &&
			((int64(lengthk) - countk) == (int64(lengthn) - countn)) &&
			(n > k) {
			slice = append(slice, int64(n))
		}
	}
	fmt.Println(slice)
	if m > len(slice) {
		return -1
	}
	return slice[m-1]
}

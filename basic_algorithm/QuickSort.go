package main

import "fmt"

func main() {
	numbers := []int{6, 2, 7, 7, 3, 8, 9}
	//fmt.Println(numbers)
	QuickSort(numbers)

	fmt.Println(numbers)
}

func QuickSort(values []int) {
	length := len(values)
	if length <= 1 {
		return
	}
	mid, i := values[0], 1    // 取第一个元素作为分水岭，i下标初始为1，即分水岭右侧的第一个元素的下标
	head, tail := 0, length-1 // 头尾的下标

	// 如果头和尾没有相遇，就会一直触发交换
	for head < tail {
		if values[i] > mid {
			// 如果分水岭右侧的元素大于分水岭，就将右侧的尾部元素和分水岭右侧元素交换
			values[i], values[tail] = values[tail], values[i]
			tail-- // 尾下标左移一位
		} else {
			// 如果分水岭右侧的元素小于等于分水岭，就将分水岭右移一位
			values[i], values[head] = values[head], values[i]
			head++ // 头下标右移一位
			i++    // i下标右移一位
		}
	}

	// 分水岭左右的元素递归做同样处理
	QuickSort(values[:head])
	QuickSort(values[head+1:])
}

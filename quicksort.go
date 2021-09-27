// 实现快排升序

package main

import "fmt"

var orgArray = []int{3, 1, 2, 2, 19, 8, 7, 10, 5, 2}
var c = 0

//方法一：利用思想，直接创建左右新数组，append到尾部。有不断创建数组的损耗
func quicksort(arr []int) []int {

	c = c + 1
	fmt.Println(c, arr)
	// if c > 10 {
	// 	return orgArray
	// }

	if len(arr) <= 1 {
		return arr
	}

	var base = 0
	leftArr := []int{}
	rightArr := []int{}
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[base] {
			leftArr = append(leftArr, arr[i])
		} else {
			rightArr = append(rightArr, arr[i])
		}
	}
	fmt.Println(c, leftArr, rightArr)
	// return orgArray
	return append(append(quicksort(leftArr), arr[0]), quicksort(rightArr)...)
}

//方法二： 经典的在同一个数组交换法，效率高
func quicksort2(arr []int) []int {
	fmt.Println("classic quicksort , perfermance begin", arr)
	if len(arr) < 2 {
		return arr
	}
	//取第一个做pivolt
	_quicksort2(arr, 0, len(arr)-1)
	fmt.Println("classic quicksort , perfermance end", arr)
	return arr
}

func _quicksort2(arr []int, left int, right int) {
	if right-left < 1 {
		return
	}

	i := left + 1
	j := right
	pivot := arr[left] //取第一个做参考
	for i <= j {
		if arr[j] < pivot { //升序，从右边找出比参考小的数，待交换
			//j要交换
		} else {
			j--
		}
		if arr[i] > pivot {
			//i要交换
		} else {
			i++
		}
		//交换
		if arr[j] < pivot && arr[i] > pivot && i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
		fmt.Println(i, j, arr)
	}
	//交换完一轮, 要将参考点换到分界位置
	arr[left], arr[j] = arr[j], arr[left]
	fmt.Println(arr, left, j-1, i, right)
	// 对左边、右边进行递归
	_quicksort2(arr, left, j-1)
	_quicksort2(arr, i, right)
}

//方法三：在同一个数组内交换，更清晰的写法
func quicksort3(arr []int) {
	_quicksort3(arr, 0, len(arr)-1)
}
func _quicksort3(arr []int, left int, right int) {
	if right-left < 1 { //数组没有2个元素，直接返回
		return
	}
	i := left + 1
	j := right
	pivot := arr[left]
	for i <= j {
		for arr[i] <= pivot {
			i++
		}
		for arr[j] >= pivot {
			j--
		}
		//这时候要交换了
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
		fmt.Println(i, j, arr)
	}
	//把pivot 换到分界位置
	arr[left], arr[j] = arr[j], arr[left]
	fmt.Println(arr, left, j-1, i, right)
	_quicksort3(arr, left, j-1)
	_quicksort3(arr, i, right)
	fmt.Println(arr)
}

// func quick(i, m, j)
func main() {
	// fmt.Println("quicksort up")
	// sortArray := quicksort(orgArray)
	// fmt.Println(orgArray)
	// fmt.Println(sortArray)

	quicksort3(orgArray)
}
//结论:最优和清晰的是quicksort3方法

// 快排：最坏时间复杂度 O(logN^2)  在顺序表情况下，因为这时候根本没有折半操作
// 最优时间复杂度O(NlogN)  ，折半操作 2^x = N , x = log2^N  ,  每次操作时间复杂度N   那么Nlog2^N
// 空间复杂度， 其实就是递归造成的树高度 最优 log2^N, 最坏N
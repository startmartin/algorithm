// 实现冒泡升序排序

package main

import "fmt"

var orgArray = []int{3,1,19,8,7,10,5,2}
func bubblesort(arr []int)  {
	arrLen := len(arr)
	for i := 0; i < len(arr) - 1; i++ {
		for j := 0; j < arrLen - 1 - i; j++ {
			if arr[j] > arr[j + 1]{	//升序
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
		fmt.Println(i, orgArray)
	}
}
func main(){
	fmt.Println("bubble up")
	bubblesort(orgArray)
	fmt.Println(orgArray)
}
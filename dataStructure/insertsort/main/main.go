package main

import "fmt"

func main() {
	s := [7]int{10, 2, 3, -5, 4, 29, 79}
	insertSort(&s)
	fmt.Println(s)
}

// 插入排序法
func insertSort(arr *[7]int) {
	for i := 1; i < len(arr); i++ { //这里len(arr)不需要-1,因为有可能最后一个是最大的,减1最后一个不会去跟有序数组进行比较
		insertVal := arr[i] //这里会保存无序数组里边元素
		insertIndex := i - 1

		//从大到小,这里会在有序数组里边进行循环比较找到合适位置
		for insertIndex >= 0 && arr[insertIndex] > insertVal {
			arr[insertIndex+1] = arr[insertIndex] //把有序数组里的元素往后移
			insertIndex--                         //条件成立要把位置让出来给insertVal
		}

		//如果比有序数组里的元素小就没必要交换位置，这里是从大到小排序
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}

		//fmt.Println(*arr)

		// fmt.Printf("arr %d = %v\n", i, *arr)
	}

	//推导过程
	// //给第2个元素找到合适位置插入
	// insertVal := arr[1]
	// insertIndex := 0 //1-1

	// //从大到小
	// for insertIndex >= 0 && arr[insertIndex] < insertVal {
	// 	arr[insertIndex+1] = arr[insertIndex] //把有序数组里的元素往后移
	// 	insertIndex--
	// }

	// //如果比有序数组里的元素小就没必要交换位置，这里是从大到小排序
	// if insertIndex+1 != 1 {
	// 	arr[insertIndex+1] = insertVal
	// }

	// //第一次插入后的arr
	// fmt.Println("arr 1 = ", *arr)

	//给第3个元素找到合适位置插入
	// insertVal := arr[2]
	// insertIndex := 1 //2-1

	// //从大到小
	// for insertIndex >= 0 && arr[insertIndex] < insertVal {
	// 	arr[insertIndex+1] = arr[insertIndex] //把有序数组里的元素往后移
	// 	insertIndex--
	// }

	// //如果比有序数组里的元素小就没必要交换位置，这里是从大到小排序
	// if insertIndex+1 != 2 {
	// 	arr[insertIndex+1] = insertVal
	// }

	// //第2次插入后的arr
	// fmt.Println("arr 2 = ", *arr)
}

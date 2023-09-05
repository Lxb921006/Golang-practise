package main

import "fmt"

func main() {
	s := [7]int{2, 3, 10, -5, 4, 29, 79}
	selectSort(&s)
	fmt.Println("s = ", s)
}

//普通的排序
// func SelectSort(a *[5]int) {
// 	for i := 0; i < len(a); i++ { //这层循环作用是读取每个元素跟下一层循环元素做比较
// 		for t := 0; t < len(a)-1; t++ { //这层循环作用是比较每一个元素
// 			if a[t] < a[t+1] {
// 				a[t], a[t+1] = a[t+1], a[t]
// 			}
// 		}
// 	}
// }

// 选择排序法
func selectSort(a *[7]int) {
	//推导
	//假设第1个元素是最大的
	// max := a[0]
	// maxIndex := 0
	// for i := 0 + 1; i < len(a); i++ { //这里跳过max,开始循环的跟max进行比较,条件成立就先保存起来
	// 	if max < a[i] {
	// 		max = a[i] //先不需要交换位置,核心思想就是最大减少交换位置次数
	// 		maxIndex = i
	// 	}
	// }

	// if maxIndex != 0 {
	// 	a[0], a[maxIndex] = a[maxIndex], a[0]
	// }

	//假设第2个元素是最大的
	// max := a[1]
	// maxIndex := 1
	// for i := 1 + 1; i < len(a); i++ {
	// 	if max < a[i] {
	// 		max = a[i] //先不需要交换位置,核心思想就是最大减少交换位置次数
	// 		maxIndex = i
	// 	}
	// }

	// if maxIndex != 1 {
	// 	a[0], a[maxIndex] = a[maxIndex], a[0]
	// }

	for j := 0; j < len(a); j++ {
		max := a[j]
		maxIndex := j
		for i := j + 1; i < len(a); i++ { //忽略掉跟自己的对比
			//从大到小
			if max < a[i] {
				max = a[i] //先不需要交换位置,核心思想就是最大减少交换位置次数
				maxIndex = i
			}
		}

		if maxIndex != j { //如果都没有找到比最开始定义的max大就不交换位置
			a[j], a[maxIndex] = a[maxIndex], a[j]
		}

		// fmt.Printf("第%d = %v\n", j, *a)
	}

}

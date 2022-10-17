package main

import (
	"fmt"
)

//快速排序法,是对冒泡排序法的改进
//小到大排序
//选择一个中轴(也不一定要用这个,也可用最后一个)(最小的下标+最大的下标)/2
//数组里的每个元素跟中轴比 arr[0+(len(arr)-1)/2]
//把比中轴小的数放到中轴的左边 left, 比中轴大的放右边 right
//然后左右两边的元素继续如此类推 - 递归
//最终就得到一个从小到大的有序数组
func main() {
	s := []int{0, 5, 1, -4, 100, 3}
	fmt.Println("s 1  = ", s)
	FastSort(s, 0, len(s)-1)
	fmt.Println("s 2  = ", s)
}

func FastSort(arr []int, left, right int) {
	l := left
	r := right
	pivot := arr[(left+right)/2]
	temp := 0

	//这里是分解arr, 把比pivot小的放左边, 比pivot大的放右边
	for l < r {
		//把比pivot小的元素放到左边
		for arr[l] < pivot { //小到大
			l++
		}
		//把比pivot大的元素放到右边
		for arr[r] > pivot { //小到大
			r--
		}
		//如果这里成立就表示分解完成
		if l >= r {
			break
		}
		//交换位置
		temp = arr[l]
		arr[l], arr[r] = arr[r], temp
		//优化：防止在递归之后的排序退出不了最外层的for循环
		if arr[l] == pivot {
			r--
		}
		if arr[r] == pivot {
			l++
		}
	}

	//这里的作用就是给结束递归设置的条件
	if l == r {
		l++
		r--
	}

	//向左递归
	if left < r {
		FastSort(arr, left, r)
	}
	//向右递归
	if right > l {
		FastSort(arr, l, right)
	}

}

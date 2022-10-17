package main

import "fmt"

//地图必须是同一个
//i,j表示对地图的哪个坐标进行测试
func Maze(td *[8][7]int, i, j int) bool {
	//这里规定到达到6，5坐标就表示走通
	if td[6][5] == 2 { //这个是退出条件
		fmt.Println("1 = ", i, j)
		return true
	} else {
		if td[i][j] == 0 { //看看这个坐标是否可以探路
			//先假设这个坐标可以探路
			//改成下右上左
			td[i][j] = 2          // 通了就把坐标改成2
			if Maze(td, i+1, j) { //下
				return true
			} else if Maze(td, i, j+1) { //右
				return true
			} else if Maze(td, i-1, j) { //上
				return true
			} else if Maze(td, i, j-1) { //左
				return true
			} else { //死路，回溯
				td[i][j] = 3
				fmt.Println(i, j)
				return false
			}
		} else { //说明这个坐标不能探路因为为1
			return false
		}
	}
}

func main() {
	//元素1表示墙
	//没有走过为0
	//元素2表示通路
	//3表示走过的死路
	m := [8][7]int{}
	//先设置围墙,即二维数组最外层元素设为1
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if i == 0 || i == len(m)-1 {
				m[i][j] = 1
			} else {
				if j == 0 || j == len(m[i])-1 {
					m[i][j] = 1
				}
			}
		}
	}

	m[3][1] = 1
	m[3][2] = 1
	m[1][2] = 1
	m[2][2] = 1

	for _, v1 := range m {
		for _, v2 := range v1 {
			fmt.Printf("%d ", v2)
		}
		fmt.Println()
	}
	Maze(&m, 1, 1)
	fmt.Println("迷宫")
	for _, v1 := range m {
		for _, v2 := range v1 {
			fmt.Printf("%d ", v2)
		}
		fmt.Println()
	}
}

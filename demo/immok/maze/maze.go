package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}

	var row, col int
	fmt.Fscanf(file, "%d %d\r\n", &row, &col) // 读取文件

	//content,_ := ioutil.ReadAll(file)
	//fmt.Println(string(content))
	//fmt.Printf("%s\n","-----------")
	maze := make([][]int, row)
	//fmt.Println("the maze is:",maze)

	for i := range maze {
		maze[i] = make([]int, col) // 每一行有col列
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
		fmt.Fscanf(file, "\r\n")
	}
	return maze
}

type point struct {
	i, j int
}

var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	// 上下越界
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	// 左右越界
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))

	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	Q := []point{start} // 初始值

	// 队列非空执行
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)

			// maze at next is 0
			// and steps at next is 0
			// and next != start
			val, ok := next.at(maze)
			if !ok || val == 1 { // 剔除墙或者边界
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 { // 不等于0说明走过了
				continue
			}

			if next == start { // 回到原点
				continue
			}

			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1

			Q = append(Q, next)
		}
	}

	return steps
}

func main() {
	filename := "immok/maze/maze.in"

	maze := readMaze(filename)
	/*for _,row := range maze {
		for _,val := range row {
			fmt.Printf("%d ",val)
		}
		fmt.Println()
	}*/

	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}

}

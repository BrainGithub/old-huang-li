package main

import "fmt"

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if n == 0 || grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}

	type point struct {
		x int
		y int
		step int
	}
	var q = make(chan point, n*n)
	q <- point{0,0,1}
	grid[0][0] = -1

	dx := []int {-1,-1,-1,0,0,1,1,1}
	dy := []int {-1,0,1,-1,1,-1,0,1}

	for len(q) > 0 {
		pt := <- q
		for i:=0; i<len(dx); i++ {
			xx := pt.x + dx[i]
			yy := pt.y + dy[i]
			if xx == n-1 && yy == n-1 {
				return pt.step + 1
			}
			if (xx >= 0 && xx < n) && (yy >= 0 && yy < n) && grid[xx][yy] == 0 {
				grid[xx][yy] = -1
				q <- point{xx, yy, pt.step + 1}
			}
		}
	}

	return -1
}

func test_shortestPathBinaryMatrix() {
	grid := [][]int{{0,0,0},{1,1,0},{1,1,0}}
	n := shortestPathBinaryMatrix(grid)
	fmt.Println(n)
}


func findOrder(numCourses int, prerequisites [][]int) []int {
	var indegree = make([]int, numCourses)
	var adj = map[int][]int{}
	var res = []int{}
	for i:=0; i<numCourses; i++ {
		adj[i] = []int {}
	}
	for _, item := range prerequisites {
		second, first := item[0], item[1]
		adj[first] = append(adj[first], second)
		indegree[second] += 1
	}
	q := []int{}
	for i:=0; i<numCourses; i++ {
		if indegree[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		course := q[0]
		q = q[1:]
		res = append(res, course)
		for _, successor := range adj[course] {
			indegree[successor] -= 1
			if indegree[successor] == 0 {
				q = append(q, successor)
			}
		}
	}

	if len(res) != numCourses {
		res = []int {}
	}
	return res
}


func test_findOrder() {
	res := findOrder(2, [][]int{[]int{1, 0},})
	fmt.Printf("%+v", res)
}

func testDefer(a * int) {
	fmt.Println(*a)
}

func main() {
	a := 1
	defer testDefer(&a)
	a += 1
	panic("aaa")
	test_findOrder()
}




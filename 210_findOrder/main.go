package main

import "fmt"

// 210. 课程表 II
// 拓扑排序 和207 课程表基本相似
func findOrder(numCourses int, prerequisites [][]int) []int {
	edges := make([][]int, numCourses)  //邻接表
	indegree := make([]int, numCourses) //入度
	var q []int
	for _, v := range prerequisites {
		edges[v[1]] = append(edges[v[1]], v[0])
		indegree[v[0]]++
	}
	for k, v := range indegree {
		if v == 0 {
			q = append(q, k)
		}
	}
	//fmt.Println("edges:", edges)
	//fmt.Println("indegree:", indegree)
	//fmt.Println("q:", q)
	var ans []int
	for len(q) != 0 {
		front := q[0]
		q = q[1:]
		ans = append(ans, front)
		for _, v := range edges[front] {
			indegree[v]--
			if indegree[v] == 0 {
				q = append(q, v)
			}
		}
	}
	if len(ans) == numCourses {
		return ans
	}
	return nil
}

func main() {
	numCourses := 2
	prerequisites := [][]int{{1, 0}}
	fmt.Println(findOrder(numCourses, prerequisites))
	numCourses = 4
	prerequisites = [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	fmt.Println(findOrder(numCourses, prerequisites))
}

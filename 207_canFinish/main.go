package main

import "fmt"

// 其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi
// 拓扑排序
func canFinish(numCourses int, prerequisites [][]int) bool {
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
	count := 0
	for len(q) != 0 {
		front := q[0]
		q = q[1:]
		count++
		for _, v := range edges[front] {
			indegree[v]--
			if indegree[v] == 0 {
				q = append(q, v)
			}
		}
	}
	return count == numCourses
}

func main() {
	numCourses := 2
	prerequisites := [][]int{{1, 0}}
	fmt.Println(canFinish(numCourses, prerequisites))
	numCourses = 2
	prerequisites = [][]int{{1, 0}, {0, 1}}
	fmt.Println(canFinish(numCourses, prerequisites))
}

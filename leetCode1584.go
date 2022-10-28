package main

import (
	"fmt"
	"sort"
)

type UnionFind struct {
	parent []int
	size   []int
}

func newUnionFind(numOfNodes int) *UnionFind {
	// makeSet
	parent := make([]int, numOfNodes)
	size := make([]int, numOfNodes)
	for i := 0; i < numOfNodes; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &UnionFind{
		parent: parent,
		size:   size,
	}
}

func (uf *UnionFind) find(node int) int {
	for node != uf.parent[node] {
		uf.parent[node] = uf.parent[uf.parent[node]]
		node = uf.parent[node]
	}
	return node
}

func (uf *UnionFind) union(node1, node2 int) bool {
	root1 := uf.find(node1)
	root2 := uf.find(node2)

	if root1 == root2 {
		return false
	}

	if uf.size[root1] > uf.size[root2] {
		uf.parent[root2] = root1
		uf.size[root1] += 1
	} else {
		uf.parent[root1] = root2
		uf.size[root2] += 1
	}
	return true
}

type edgeInfo struct {
	distance int
	node1    int
	node2    int
}

func minCostConnectPoints(points [][]int) int {
	numOfNodes := len(points)

	edges := []*edgeInfo{}
	for i := 0; i < numOfNodes; i++ {
		for j := i + 1; j < numOfNodes; j++ {
			distance := abs(points[i][0]-points[j][0]) + abs(points[i][1]-points[j][1])
			edges = append(edges, &edgeInfo{distance: distance, node1: i, node2: j})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].distance < edges[j].distance
	})

	uf := newUnionFind(numOfNodes)

	cost := 0
	edgeCount := 0
	for _, edge := range edges {
		if uf.union(edge.node1, edge.node2) {
			cost += edge.distance
			edgeCount += 1
			if edgeCount == numOfNodes-1 {
				return cost
			}
		}
	}
	return cost
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	points := [][]int{
		{0, 0},
		{2, 2},
		{3, 10},
		{5, 2},
		{7, 0},
	}

	fmt.Println(minCostConnectPoints(points))
}

// output: 20

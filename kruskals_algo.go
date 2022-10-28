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

func kruskalsAlgo(edges [][]int, numOfNodes int) int {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][0] < edges[j][0]
	})

	uf := newUnionFind(numOfNodes)

	minWeight := 0
	edgeCount := 0
	MSTedges := [][]int{}
	for _, edge := range edges {
		weight, node1, node2 := edge[0], edge[1], edge[2]
		if uf.union(node1, node2) {
			minWeight += weight
			edgeCount += 1
			MSTedges = append(MSTedges, []int{node1, node2})
			if edgeCount == numOfNodes-1 {
				fmt.Println("edges in MST", MSTedges)
				return minWeight
			}
		}
	}
	return minWeight
}

func main() {
	edges := [][]int{
		// weight, node1, node2
		{2, 0, 2},
		{6, 0, 3},
		{5, 1, 2},
		{1, 1, 4},
		{2, 2, 3},
		{3, 3, 4},
	}
	numOfNodes := 5
	fmt.Println("minimum total weight:", kruskalsAlgo(edges, numOfNodes))
}

/*
output:
edges in MST [[1 4] [0 2] [2 3] [3 4]]
minimum total weight: 8
*/

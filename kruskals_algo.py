class UnionFind:
    def __init__(self, numOfNodes):
        self.parent = self.makeSet(numOfNodes)
        self.size = [1 for _ in range(numOfNodes)]
    
    def makeSet(self, numOfNodes):
        return [x for x in range(numOfNodes)]
    
    def find(self, node):
        while node != self.parent[node]:
            self.parent[node] = self.parent[self.parent[node]]
            node = self.parent[node]
        return node
    
    def union(self, node1, node2):
        root1 = self.find(node1)
        root2 = self.find(node2)
        if root1 == root2:
            return False
     
        if self.size[root1] > self.size[root2]:
            self.parent[root2] = root1
            self.size[root1] += 1
        else:
            self.parent[root1] = root2
            self.size[root2] += 1
                
        return True
        
def kruskalsAlgo(edges, numOfNodes):
    # sort by weight
    edges.sort()
    edgeCount = 0
    MSTedges = []
    uf = UnionFind(numOfNodes)

    minWeight = 0
    for weight, node1, node2 in edges:
        if uf.union(node1, node2):
            minWeight += weight
            edgeCount += 1
            MSTedges.append((node1, node2))
            if edgeCount == numOfNodes - 1:
                print("edges in MST", MSTedges)
                return minWeight
  
    return minWeight

if __name__ == "__main__":
    edges = [
        # weight, node1, node2
        [2, 0, 2],
        [6, 0, 3],
        [5, 1, 2],
        [1, 1, 4],
        [2, 2, 3],
        [3, 3, 4],
    ]
    numOfNodes = 5
    print("minimum total weight", kruskalsAlgo(edges, numOfNodes))
            
"""
output:
('edges in MST', [(1, 4), (0, 2), (2, 3), (3, 4)])
('minimum total weight', 8)
"""
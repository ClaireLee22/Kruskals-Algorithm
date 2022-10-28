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
        
        
class Solution(object):
    def minCostConnectPoints(self, points):
        """
        :type points: List[List[int]]
        :rtype: int
        """
        numOfNodes = len(points) 
        edges = []
        for i in range(numOfNodes):
            for j in range(i+1, numOfNodes):
                distance = abs(points[i][0]-points[j][0]) + abs(points[i][1]-points[j][1])
                edges.append((distance, i, j))
        
        # sort based on distance
        edges.sort()
        
        # Kruskal's algorithm
        cost = 0
        edgeCount = 0
        
        uf = UnionFind(numOfNodes)
        for distance, node1, node2 in edges:
            if uf.union(node1, node2):
                cost += distance
                edgeCount += 1
                if edgeCount == numOfNodes - 1:
                    return cost
        
        return cost
            
if __name__ == "__main__":
    points = [[0,0],[2,2],[3,10],[5,2],[7,0]]
    solution = Solution()
    print(solution.minCostConnectPoints(points))

# output: 20
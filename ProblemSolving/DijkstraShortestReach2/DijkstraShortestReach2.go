package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io/ioutil"
	"os"
)

//PRIORITY QUEUE FROM https://pkg.go.dev/container/heap#example-package-PriorityQueue

//holds information about the current edge
type edge struct {
	currentNode int // the node where we are in (current node label).
	cost        int // the cost needed to reach this node.
	from        int // the node from where this node was reached
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*edge

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*edge)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

// calculates the cost of traveling to a node
func cost(currentNode edge, target []int) edge {
	cost := currentNode.cost + target[1]
	return edge{currentNode: target[0], cost: cost, from: currentNode.currentNode}
}

// // find the shortest path using the dijkstra algorithm
// func dijkstra(graph map[int][][]int, start int, goal int) edge {
//   // if the queue is not empty
//   for len(pq) > 0 {
//     currentEdge := heap.Pop(&pq).(*edge) //take the current shortest traveled distance node
//     // if this node havn't been visited yet
//     if !nodesVisited[currentEdge.currentNode] {
//       neighbors := graph[currentEdge.currentNode] // get its neighbors
//       for _, val := range neighbors {
//         // test if the current edge have been used
//         _, ok := visitedEdge[getVisitedEdgeKey(val[0], val[1])]
//         // if not
//         if !ok {
//           cost := cost(*currentEdge, val)                       // calculate the cost of using this edge
//           heap.Push(&pq, &cost)                                 // put it in the priority queue
//           nodesVisited[currentEdge.currentNode] = true          // mark node as visited
//           visitedEdge[getVisitedEdgeKey(val[0], val[1])] = true // mark edge as used
//         }
//       }

//       cacheControl(*currentEdge)

//       if currentEdge.currentNode == goal {
//         return *currentEdge
//       }
//     }
//   }
//   return edge{currentNode: goal, cost: -1, from: -1} // no path to goal found
// }

func initializCache(cache []int) {
	for i := 0; i < len(cache); i++ {
		cache[i] = -1
	}
}

// find the shortest path using the dijkstra algorithm
func dijkstra(graph map[int][][]int, start int, nodes int) []int {
	pq := make(PriorityQueue, 0)  // priority queu
	var cache = make([]int, 3001) // to save already calculated distances
	//initializes the cache to a very high value the indices in the cache array represents the graph nodes
	initializCache(cache)

	nodesVisited := make([]bool, nodes+1) // keeps track of the visited nodes

	heap.Init(&pq) //initializes the priority queue
	// creates the starting edge and put it in the queue
	e := &edge{currentNode: start, cost: 0, from: 0}
	heap.Push(&pq, e)

	// while the queue is not empty
	for len(pq) > 0 {
		currentEdge := heap.Pop(&pq).(*edge) // take the current shortest distance to this node
		// the this edge have not been procesed before or if its cost is more than the edge in the cache just continue
		if cache[currentEdge.currentNode] != -1 && cache[currentEdge.currentNode] < currentEdge.cost {
			continue
		}
		// if node is not visited
		if !nodesVisited[currentEdge.currentNode] {
			neighbors := graph[currentEdge.currentNode] // get it's neighbors
			for _, val := range neighbors {
				cost := cost(*currentEdge, val) // calculates the cost
				// if the edge is not yet procesed and the cost in the cache is greter the the cost of traveling this edge
				if cache[cost.currentNode] == -1 || cache[cost.currentNode] > cost.cost {
					heap.Push(&pq, &cost)               // add the edge to the heap
					cache[cost.currentNode] = cost.cost // update the cost for this node in the cache
				}
				nodesVisited[currentEdge.currentNode] = true // mark node as visited
			}
		}
	}
	return cache
}

// IO FUNTIONS

// This function converts an array of bytes (a part of it) to an intenger using the following logic
// The numbers (from 0 to 9) are represented in bytes by the numbers from 48 to 57 where
// 0=48, 1=49 ... 9=57 one can get any number by substracting '0' (which is 48) from any byte
// for example in bytes 9 is represented by 57 by substractig 48 we have 57-48 = 8
// With for numbers with more than one digit each digit is represented by one individual byte
// 95 for example is represented as [57 53] do get the decimal representation we substract 48 from 57 and
// multiply the result by 10 and then add the result of substracting 48 from 53
// [(57-48)*10 + (53-48)] = [(9)*10 + (5)] = [90+5] = 95
// This formula applies for any number no matter it number of digits
// in bytes the new line character is 10 and the space is 32 so, when this numbers are found it means the complete number have read
func readInt(input []byte, start int) (int, int) {
	var result, index int
	for i := start; i < len(input); i++ {
		index = i
		if input[i] == 32 || input[i] == 10 {
			break
		}
		result = (result * 10) + int(input[i]-'0')
	}
	return result, index + 1
}

// Adds elements to the adjacency List
// Recieves as parameters the adjacency List and an slice of 3 elements where
// the index 0 is the currentNode, in index 1 is the targetNode and in index 3 the cost
// of using the edge
func addToAddjList(adjList map[int][][]int, node1, node2, node3 int) {
	val, ok := adjList[node1]
	if ok {
		newEdge := [2]int{node2, node3}
		val = append(val, newEdge[:])
		adjList[node1] = val
	} else {
		var newEdgeSlice [][]int
		newEdge := [2]int{node2, node3}
		newEdgeSlice = append(newEdgeSlice, newEdge[:])
		adjList[node1] = newEdgeSlice
	}
}

// cordinates the addjacency list creation by reading the input from a reader
// input must be three integers separated by spaces the first two integers indicate
// the node labels (from, to) and the third is the cost of using the edge
// recieves as parameters a reader from where input will be read and the number of edges
// the graph has
func buildAdjList(input []byte, start, edges int) (map[int][][]int, int) {
	var adjList = make(map[int][][]int)

	for i := 0; i < edges; i++ {
		node1, next := readInt(input, start)
		node2, next := readInt(input, next)
		node3, next := readInt(input, next)
		addToAddjList(adjList, node1, node2, node3)
		addToAddjList(adjList, node2, node1, node3)
		start = next
	}
	return adjList, start
}

// parse the part of the input that describe the number of nodes, edges
// and the description of the cost of going from one node to an other using an edge
// this last part is put into an adjacency list
// retuns an adjacency list, the number of nodes in the graph, the number of edges, the starting node
// and the next byte to read (to know from where start to parse the next test case) in the input array
func parseInput(input []byte, start int) (map[int][][]int, int, int, int, int) {
	nodes, index := readInt(input, start)
	edges, index := readInt(input, index)
	adjList, index := buildAdjList(input, index, edges)
	startNode, index := readInt(input, index)

	return adjList, nodes, edges, startNode, index
}

func main() {
	bytes, _ := ioutil.ReadAll(os.Stdin)
	testCase, index := readInt(bytes, 0)

	writer := bufio.NewWriterSize(os.Stdout, 16*1024*1024)

	for i := 0; i < testCase; i++ {

		adjList, nodes, _, startNode, readIndex := parseInput(bytes, index)

		index = readIndex

		result := dijkstra(adjList, startNode, nodes)
		for i := 1; i <= nodes; i++ {
			if i != startNode {
				fmt.Fprintf(writer, "%d ", result[i])
			}
		}

		fmt.Fprintf(writer, "\n")
	}
	writer.Flush()
}

# graph

[![Go](https://github.com/timHau/graph/actions/workflows/go.yml/badge.svg)](https://github.com/timHau/graph/actions/workflows/go.yml)

A simple graph library for Go.

## Supported Algorithms

- Bellman-Ford Algorithm
- Dijkstras Algorithm
- Breadth-first search
- Depth-first search
- Topological Sort
- Cycle Detection
- Prim Algorithm (MST)
- Floyd-Warshall Algorithm
- Kosaraju Algorithm
- Laplacian Matrix 
- Hamiltonian Path Detection (Via DP)

## Create a graph

There are multiple ways to create a graph. The first and straight forward option is to pass in an `EdgeList`

```go 
edgeList := []graph.Edge{
	{0, 1, 1},
	{1, 2, 1},
	{2, 0, 1},
}
g := graph.FromEdgeList(edgeList)
```

or you can first create the graph and then add Edges to it. The nodes will be created automatically

```go
g2 := graph.NewGraph()
g2.AddEdge(0, 1, 1)
g2.AddEdge(1, 2, 1)
g2.AddEdge(2, 0, 1)
```

or you can pass the Adjaceny Matrix to create a graph, the float values of the matrix represent the weights of the edges. This Method returns an error if the matrix is not square.

```go
g, err := graph.FromAdjMat([]float64{
    0, 2, 1, 0, 0, 0,
    3, 0, 0, 1, 1, 0,
    1, 0, 0, 0, 1, 0,
    0, 9, 0, 0, 1, 1,
    0, 1, 1, 1, 0, 1,
    0, 0, 0, 1, 1, 0,
})
```

Internally the Graph uses its Adjacency List as a data structure, so you can just create a new graph from an Adjacency List

```go
g3 := graph.FromAdjList(graph.AdjList{
    0: []graph.WeightTuple{{1, 1}, {2, 1}},
    1: []graph.WeightTuple{{0, 3}},
    2: []graph.WeightTuple{{1, 2}},
})
```

## Installing 
```sh
go get github.com/timHau/graph@v0.1.2
```

## Example usage
```sh
go run example/main.go
```

## Testing 
```sh
go test -v .
```

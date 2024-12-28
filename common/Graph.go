package common

type Graph[K comparable, T any] struct {
    Verticies map[K]*Vertex[K, T]
    Hash func(T) K
}

type Vertex[K comparable, T any] struct {
    Val T
    Edges map[K]*Edge[K, T] 
}

type Edge[K comparable, T any] struct {
    Weight int
    Vertex *Vertex[K, T]
}

func StringHash(val string) string {
    return val
}

func (this *Graph[K, T]) AddVertex(val T) {
    key := this.Hash(val)
    if _, ok := this.Verticies[key]; ok { return }
    this.Verticies[key] = &Vertex[K, T]{Val: val, Edges: map[K]*Edge[K, T]{}}
}

func (this *Graph[K, T]) AddEdge(src, dest T, weight int) {
    srcKey, destKey := this.Hash(src), this.Hash(dest)
    if _, ok := this.Verticies[srcKey]; !ok { return }
    if _, ok := this.Verticies[destKey]; !ok { return }

    this.Verticies[srcKey].Edges[destKey] = &Edge[K, T]{
        Weight: weight, 
        Vertex: this.Verticies[destKey],
    }
}

func (this *Graph[K, T]) Vertex(val T) *Vertex[K, T] {
    vert, ok := this.Verticies[this.Hash(val)]
    if !ok { return nil }
    return vert
}

func NewGraph[K comparable, T any](hash func(T) K, opts ...func(*Graph[K, T])) *Graph[K, T] {
    new := &Graph[K, T]{
        Verticies: map[K]*Vertex[K, T]{},
        Hash: hash,
    } 
    
    for _, opt := range opts {
        opt(new)
    }

    return new
}

// Return an iterator through the graph using a depth first search.
// The visit callback passes in each value and their depth in the search
func DFS[K comparable, T any](gph Graph[K, T], start T, visit func(T, int) bool) {
    startVert := gph.Vertex(start)
    if startVert == nil { return }

    type node struct {
        v *Vertex[K, T]
        depth int
    }
    stack := []node{}
    visited := map[K]bool{}

    stack = append(stack, node{startVert, 0})

    for len(stack) > 0 {
        length := len(stack)
        currNode := stack[length-1]
        stack = stack[:length-1]

        if _, ok := visited[gph.Hash(currNode.v.Val)]; ok { continue }

        if stop := visit(currNode.v.Val, currNode.depth); stop { break }
        visited[gph.Hash(currNode.v.Val)] = true

        for _, edge := range currNode.v.Edges {
            stack = append(stack, node{edge.Vertex, currNode.depth+1})
        }
    }
}

func BFS[K comparable, T any](gph Graph[K, T], start T, visit func(T, int) bool) {
    startVert := gph.Vertex(start)
    if startVert == nil { return }

    type node struct {
        v *Vertex[K, T]
        depth int
    }
    stack := []node{}
    visited := map[K]bool{}

    stack = append(stack, node{startVert, 0})

    for len(stack) > 0 {
        currNode := stack[0]
        stack = stack[1:]

        if _, ok := visited[gph.Hash(currNode.v.Val)]; ok { continue }

        if stop := visit(currNode.v.Val, currNode.depth); stop { break }
        visited[gph.Hash(currNode.v.Val)] = true

        for _, edge := range currNode.v.Edges {
            stack = append(stack, node{edge.Vertex, currNode.depth+1})
        }
    }
}

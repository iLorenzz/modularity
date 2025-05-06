package main

type Graph struct{
	vertices []int64
	edges map[int64][]int64
	number_of_egdes int64
	in_degrees map[int64]int
}

func new_Graph() *Graph{
	return &Graph{
		edges: make(map[int64][]int64),
		number_of_egdes: 0,
		in_degrees: make(map[int64]int),
	}
}

func (g *Graph) add_Edge(u, v int64){
	g.edges[u] = append(g.edges[u], v)
	g.set_number_of_edges()
	
}

func (g *Graph) set_number_of_edges() {
	g.number_of_egdes++
}

func (g *Graph) get_number_of_edges() int64 {
	return g.number_of_egdes
}

func(g *Graph) set_vertices(u, v int64){
	if g.vertices == nil{
		g.vertices = append(g.vertices, u, v)
	}
	
	count := 0
	for _, vertice := range g.vertices{
		if vertice == u{
			count++
		}
		
		if count == 1{
			continue
		}
	}
	
	if count == 0{
		g.vertices = append(g.vertices, u)
	}
	
	count = 0
	for _, vertice := range g.vertices{
		if vertice == v{
			count++
		}
		
		if count == 1{
			continue
		}
	}
	
	if count == 0{
		g.vertices = append(g.vertices, v)
	}
}
	
func (g *Graph) compute_in_degrees() {
	g.in_degrees = make(map[int64]int)
	for _, neighbors := range g.edges {
		for _, v := range neighbors {
			g.in_degrees[v]++
		}
	}
}
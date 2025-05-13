package main 

type Undirected_Graph struct{
	edges map[int64][]int64
	number_of_egdes int64
}

func new_Undirected_Graph() *Undirected_Graph{
	return &Undirected_Graph{
		edges: make(map[int64][]int64),
		number_of_egdes: 0,
	}
}

func (ug *Undirected_Graph) add_edges(u, v int64){
	ug.edges[u] = append(ug.edges[u], v)
	ug.edges[v] = append(ug.edges[v], u)
	
	ug.number_of_egdes++
}

func (ug *Undirected_Graph) get_number_of_edges() int64{
	return ug.number_of_egdes
}

func (ug *Undirected_Graph) graph_degree() int64{
	var degree int64 = 0
	
	for node := range ug.edges{
		degree += int64(len(ug.edges[node]))
	}
	
	return degree
}

func (ug *Undirected_Graph) add_edges_if_is_community(node int64, edges map[int64][]int64){
	ug.edges[node] = edges[node]
}

func (ug *Undirected_Graph) set_number_of_edges_if_is_community() {
	visited := make(map[[2]int64]bool)
	ug.number_of_egdes = 0

	for u, neighbors := range ug.edges {
		for _, v := range neighbors {
			if _, exists := ug.edges[v]; exists {
				edge := [2]int64{min(u, v), max(u, v)}
				if !visited[edge] {
					visited[edge] = true
					ug.number_of_egdes++
				}
			}
		}
	}
}
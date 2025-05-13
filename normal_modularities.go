package main

import (
	"fmt"
	"math"
)

func modularity(g *Directed_Graph, lambda_resolution float64, communities ...[]map[int64]struct{}) ([]float64){
	var gen_modularity []float64
	m := g.get_number_of_edges()
	
	g.compute_in_degrees()
	adj_map := make(map[int64]map[int64]struct{})
		for u, neighbors := range g.edges {
			adj_map[u] = make(map[int64]struct{})
			for _, v := range neighbors {
				adj_map[u][v] = struct{}{}
			}
		}
	fmt.Println(m)
	
	for _, community := range communities{
		var modularity float64 = 0
		
		for _, node_i := range g.vertices{
			var ki_out float64
			
			if g.edges[int64(node_i)] != nil{
				ki_out = float64(len(g.edges[int64(node_i)]))
			} else{
				ki_out = 0
			}
			
			for _, node_j := range g.vertices{
				
				kj_in := float64(g.in_degrees[(int64(node_j))])
				A_ij := 0.0
				if _, ok := adj_map[node_i][node_j]; ok {
					A_ij = 1.0
				}
				
				delta_ij := float64(is_same_community(int64(node_i), int64(node_j), community))
				
				expected := (ki_out*kj_in) / float64(m)
				
				modularity += ((A_ij) - lambda_resolution * expected) * delta_ij
				//fmt.Println(modularity)
			}
		}
		modularity = modularity/ float64(m)
		gen_modularity = append(gen_modularity, modularity)
	}
	
	return gen_modularity
}

func alternative_modularity(ug *Undirected_Graph, lambda_resolution float64, 
	gen_communities ...[]*Undirected_Graph,
	) []float64{
	
	var gen_modularity []float64
	m := float64(ug.get_number_of_edges())
	//fmt.Println(m)
	
	for _, communities := range gen_communities{
		var modularity float64 = 0.0
		
		for _, community := range communities{
			Lc := float64(community.get_number_of_edges())
			//fmt.Println(Lc)
			kc := float64(community.graph_degree())
			fmt.Println(kc)
			
			//fmt.Println(math.Pow(float64(kc) / 2.0*float64(m), 2.0))
			base := kc / (2.0*m)
			fmt.Println(base)
			
			modularity += (Lc / m - lambda_resolution * (math.Pow(base, 2.0)))
		}
		gen_modularity = append(gen_modularity, modularity)
	}
	
	return gen_modularity
	
} 
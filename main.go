package main

import (
	"fmt"
	"log"
	//"slices"
)

func is_same_community(u, v int64, communities []map[int64]struct{}) int64{
	
	for _, key := range communities{
		if _, exists_v := key[v]; exists_v{
			if _, exists_u := key[u]; exists_u{
				return 1
			}
		}
	}
	
	return 0
}

func modularity(g *Graph, lambda_resolution float64, communities ...[]map[int64]struct{}) ([]float64){
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

func main(){
	graph_scale_free, err := read_graph_form_file("scale_free_network.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	
	sf_all_communities, err := read_communities("out")
	if err != nil{
		log.Fatal(err)
	}
	
	boolean := is_same_community(93, 40, sf_all_communities[69])
	fmt.Println(boolean)
	
	
	//modularities := modularity(graph, 1.0, all_communities...)
	//fmt.Println(modularities[9])
		
	g_congress, err := read_graph_form_file("congress.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
		
	fmt.Println(g_congress)
		
	congress_all_communities, err := read_communities("community_gen")
	if err != nil{
		log.Fatal(err)
	}
		
	congress_modularities := paralell_modularity(g_congress, 1.0, congress_all_communities...)
	for index, modularity := range congress_modularities{
		fmt.Println(index+1, ": ", modularity)
	}
		
	congress_higher_modularity_measure := congress_modularities[0]
	var index_of_higher int
		
	for index, modularity_measure := range congress_modularities{
		if modularity_measure > congress_higher_modularity_measure{
		 	congress_higher_modularity_measure = modularity_measure
			index_of_higher = index
		}
	}
		
	fmt.Println("Best modularity measure = ", congress_higher_modularity_measure)
	fmt.Println("From community ", index_of_higher+1)
	
	scale_free_modularities := paralell_modularity(graph_scale_free, 1.0, sf_all_communities...)
	
	for index, modularity := range scale_free_modularities{
		fmt.Println(index+1, ": ", modularity)
	}
	
	sf_higher_modularity_measure := scale_free_modularities[0]
		
	for index, modularity_measure := range scale_free_modularities{
		if modularity_measure > sf_higher_modularity_measure{
		 	sf_higher_modularity_measure = modularity_measure
			index_of_higher = index
		}
	}
	
	fmt.Println("Best modularity measure = ", sf_higher_modularity_measure)
	fmt.Println("From community ", index_of_higher+1)
}






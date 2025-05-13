package main

import (
	"fmt"
	"log"
	"time"
	//"slices"
)

func main(){
	
	graph_scale_free, err := read_directed_graph_form_file("scale_free_network.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	
	sf_all_communities, err := read_communities("./sf_graph/out")
	if err != nil{
		log.Fatal(err)
	}
	
	
	sf_modularities := paralell_modularity(graph_scale_free, 1.0, sf_all_communities...)
	for index, modularity := range sf_modularities{
		fmt.Println(index+1, ": ", modularity)
	}
	
	//start_congress := time.Now()
	g_congress, err := read_directed_graph_form_file("./congress_graph/congress.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
		
	fmt.Println(g_congress)
		
	congress_all_communities, err := read_communities("./congress_graph/congress_communities_gen")
	if err != nil{
		log.Fatal(err)
	}
	
	start_congress := time.Now()
	congress_modularities := paralell_modularity(g_congress, 1.0, congress_all_communities...)
	end_congress := time.Now()
	elapsed_congress := end_congress.Sub(start_congress)
	
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
	
	/* 
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
	
	*/
	//start_caveman := time.Now()
	caveman, err := read_undirected_graph("./caveman_graph/caveman_graph.txt")
	if err != nil{
		log.Fatal(err)
	}
	
	caveman_all_communities, err := read_communities_as_graphs(caveman, "./caveman_graph/caveman_communities_gen")
	if err != nil{
		log.Fatal(err)
	}
	
	start_caveman := time.Now()
	caveman_modularities := parallel_alternative_modularity(caveman, 1.0, caveman_all_communities...)
	end_caveman := time.Now()
	elapsed_caveman := end_caveman.Sub(start_caveman)
	
	for i, modularity := range caveman_modularities{
		fmt.Println(i+1, ": ", modularity)
	}
	
	fmt.Println(" ")
	
	fmt.Println("Times:")
	fmt.Println("method node by node: ", elapsed_congress)
	fmt.Println("method from communities: ", elapsed_caveman)
}






package main

import (
	"bufio"
	//"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func read_undirected_graph(file_path string) (*Undirected_Graph, error){
	file, err := os.Open(file_path)
	if err != nil{
		return nil, err
	}
	
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	
	ug := new_Undirected_Graph()
	
	for scanner.Scan(){
		line := scanner.Text()
		node := strings.Split(line, " ")
		
		u, err1 := strconv.Atoi(node[0])
		v, err2 := strconv.Atoi(node[1])
		
		if err1 != nil || err2 != nil{
			continue
		}
		
		ug.add_edges(int64(u), int64(v))
	}
	
	return ug, nil
}

func read_communities_as_graphs(graph_father *Undirected_Graph, dir string) ([][]*Undirected_Graph, error){
	entries, err := os.ReadDir(dir)
	if err != nil{
		return nil, err
	}
	
	var gen_communities [][]*Undirected_Graph
	
	for _, entry := range entries{
		file_number := strings.Split(entry.Name(), "_")
		gen_of_file := strings.Split(file_number[0], ".")
		gen, err := strconv.Atoi(gen_of_file[0])
		
		if err != nil{
			return nil, err
		}
		
		file_path := filepath.Join(dir, entry.Name())
		
		file, err := os.Open(file_path)
		if err != nil{
			return nil, err
		}
		
		defer file.Close()
		
		scanner := bufio.NewScanner(file)
		
		var communities []*Undirected_Graph
		
		for scanner.Scan(){
			line := scanner.Text()
			
			community, err := new_community(graph_father, line)
			if err != nil{
				return nil, err
			}
			communities = append(communities, community)
		}
		
		for len(gen_communities) < gen{
			gen_communities = append(gen_communities, nil)
		}
		gen_communities[gen-1] = communities
	}
	
	return gen_communities, nil
}

func new_community(graph_father *Undirected_Graph, line string) (*Undirected_Graph, error){
	ug := new_Undirected_Graph()
	str_nodes := strings.Split(line, " ")
	var nodes []int64
	
	for _, str_node := range str_nodes{
		value, err := strconv.Atoi(str_node)
		if err != nil{
			return nil, err
		}
		
		nodes = append(nodes, int64(value))
	}
	
	for _, node := range nodes{
		ug.add_edges_if_is_community(node, graph_father.edges)
	}
	
	ug.set_number_of_edges_if_is_community()
	
	return ug, nil
}
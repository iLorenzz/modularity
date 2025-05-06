package main

import (
	"bufio"
	"path/filepath"
	"os"
	"strconv"
	"strings"
)

func read_graph_form_file(file_path string) (*Graph, error){
	file, err := os.Open(file_path)
	if err != nil{
		return nil, err
	}
	
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	
	graph := new_Graph()
	
	for scanner.Scan(){
		line := scanner.Text()
		parts := strings.Split(line, " ")
		
		u, err1 := strconv.Atoi(parts[0])
		v, err2 := strconv.Atoi(parts[1])
		
		if err1 != nil || err2 != nil{
			continue
		}
		
		graph.set_vertices(int64(u), int64(v))
		graph.add_Edge(int64(u), int64(v))
	}
	
	if err := scanner.Err(); err != nil{
		return nil, err
	}
	
	return graph, nil
}
	
	func read_communities(dir string) ([][]map[int64]struct{}, error) {
    entries, err := os.ReadDir(dir)
    if err != nil {
        return nil, err
    }

    all_communities := [][]map[int64]struct{}{}

    for _, entry := range entries {
        gen_file := strings.Split(entry.Name(), ".")
        gen, err := strconv.Atoi(gen_file[0])
        if err != nil {
            return nil, err
        }

        file, err := os.Open(filepath.Join(dir, entry.Name()))
        if err != nil {
            return nil, err
        }

        scanner := bufio.NewScanner(file)
        
        var communities []map[int64]struct{}
        
        for scanner.Scan() {
            line := scanner.Text()
            parts := strings.Fields(line)

            community := make(map[int64]struct{})
            for _, value := range parts {
                node, err := strconv.ParseInt(value, 10, 64)
                if err != nil {
                    file.Close()
                    return nil, err
                }
                community[node] = struct{}{}
            }
            communities = append(communities, community)
        }
        file.Close()

        for len(all_communities) < gen {
            all_communities = append(all_communities, nil)
        }
        all_communities[gen-1] = communities
    }

    return all_communities, nil
}
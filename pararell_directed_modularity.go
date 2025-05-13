package main

import (
	"context"
	"runtime"
)

type Modularity_job struct{
	index int
	community []map[int64]struct{}
	modularity_result float64
	g *Directed_Graph
	lambda_resolution float64
	in_degrees        map[int64]int 
}

func (job *Modularity_job) execute(){
	m := job.g.get_number_of_edges()
	
	//job.g.compute_in_degrees()
	
	adj_map := make(map[int64]map[int64]struct{})
	for u, neighbors := range job.g.edges {
		adj_map[u] = make(map[int64]struct{})
		for _, v := range neighbors {
			adj_map[u][v] = struct{}{}
		}
	}
	
		var modularity float64 = 0
		
		for _, node_i := range job.g.vertices{
			var ki_out float64
			
			if job.g.edges[int64(node_i)] != nil{
				ki_out = float64(len(job.g.edges[int64(node_i)]))
			} else{
				ki_out = 0
			}
			
			for _, node_j := range job.g.vertices{
				
				kj_in := float64(job.g.in_degrees[(int64(node_j))])
				A_ij := 0.0
				if _, ok := adj_map[node_i][node_j]; ok {
					A_ij = 1.0
				}
				
				delta_ij := float64(is_same_community(int64(node_i), int64(node_j), job.community))
				
				expected := (ki_out*kj_in) / float64(m)
				
				modularity += ((A_ij) - job.lambda_resolution * expected) * delta_ij
			}
		}	
		job.modularity_result = modularity/ float64(m)
}

func do_parallel(ctx context.Context, inputs <- chan *Modularity_job, output chan <- *Modularity_job){
	for{
		select{
			case job, ok := <- inputs:
			if !ok{
				return
			}
			
			job.execute()
			output <- job
			
			case <- ctx.Done():
				return
		}
	}
}

func paralell_modularity(g *Directed_Graph, lambda_resolution float64, communities...[]map[int64]struct{}) []float64{
	ctx := context.Background()
	max := runtime.NumCPU() + 2
	queue := make(chan *Modularity_job, max)
	output := make(chan *Modularity_job, len(communities))
	defer close(output)
	
	g.compute_in_degrees()
	in_degrees := g.in_degrees
	
	for i := 0; i < max; i++{
		go do_parallel(ctx, queue, output)
	}
	
	jobs := make([]*Modularity_job, len(communities))
	for i, community := range communities{
		jobs[i] = &Modularity_job{
			index: i,
			community: community,
			g: g,
			lambda_resolution: lambda_resolution,
			in_degrees: in_degrees,
		}
	}
	
	go func(){
		for _, job := range jobs{
			queue <- job
		}
		close(queue)
	}()
	
	modularity_results := make([]float64, len(communities))
	for i := 0; i < len(communities); i++{
		select{
			case job := <- output:
				modularity_results[job.index] = job.modularity_result
			case <- ctx.Done():
				break
		}
	}
	return modularity_results
}
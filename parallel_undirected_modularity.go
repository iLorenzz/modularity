package main

import (
	"context"
	"math"
	"runtime"
)

type Alternative_Modularity_job struct{
	index int
	communities []*Undirected_Graph
	modularity_result float64
	ug *Undirected_Graph
	lambda_resolution float64
}

func (job *Alternative_Modularity_job) execute(){
	m := float64(job.ug.get_number_of_edges())
	var modularity float64 = 0.0
	
	for _, community := range job.communities{
		Lc := float64(community.get_number_of_edges())
		kc := float64(community.graph_degree())
		
		base := kc / (2.0*m)
		modularity += (Lc / m - job.lambda_resolution * (math.Pow(base, 2.0)))
	}
	
	job.modularity_result = modularity
}

func a_do_parallel(ctx context.Context, inputs <- chan *Alternative_Modularity_job, 
	output chan <- *Alternative_Modularity_job){
		for {
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
	
func parallel_alternative_modularity(ug *Undirected_Graph, lambda_resolution float64, 
	gen_communities ...[]*Undirected_Graph) []float64{
		ctx := context.Background()
		max := runtime.NumCPU() + 2
		queue := make(chan *Alternative_Modularity_job, max)
		output := make(chan *Alternative_Modularity_job, len(gen_communities))
		defer close(output)
		
		for i := 0; i < max; i++{
			go a_do_parallel(ctx, queue, output)
		}
		
		jobs := make([]*Alternative_Modularity_job, len(gen_communities))
		for i, communities := range gen_communities{
			jobs[i] = &Alternative_Modularity_job{
				index: i,
				communities: communities,
				ug: ug,
				lambda_resolution: lambda_resolution,
			}
		}
		
		go func(){
			for _, job := range jobs{
				queue <- job
			}
			close(queue)
		}()
		
		modularity_results := make([]float64, len(gen_communities))
		for i := 0; i < len(gen_communities); i++{
			select{
				case job := <- output:
					modularity_results[job.index] = job.modularity_result
				case <- ctx.Done():
					break
			}
		}
		
		return modularity_results
}
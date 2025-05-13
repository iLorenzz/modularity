package main

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

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
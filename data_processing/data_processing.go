package dataprocessing

func DataProcessing(elements []int) []int {

	type payload struct {
		pos int
		el  int
	}

	workerOut := make(chan payload, len(elements))

	// Send elements
	for k, v := range elements {
		//result = append(result, v*2)
		go func(k, v int) {
			workerOut <- payload{pos: k, el: v * 2}
		}(k, v)
	}

	// Receive results
	result := make([]int, len(elements))
	for i := 0; i < len(elements); i++ {
		resEl := <-workerOut
		result[resEl.pos] = resEl.el
	}

	return result
}

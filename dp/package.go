package dp

// 0-1 package
func ZeroOnePackage(weights []int, values []int, size int) int {
	n := len(weights)
	mem := make([][]int, n)

	for i := range n {
		mem[i] = make([]int, size+1)
	}

	// When idx is 0, only can use one item
	for i := weights[0]; i <= size; i++ {
		mem[0][i] = values[0]
	}

	// Dp start

	// The outer is item
	for i := 1; i < n; i++ {
		for j := 0; j <= size; j++ {
			if j < weights[i] {
				mem[i][j] = mem[i-1][j]
			} else {
				mem[i][j] = max(mem[i-1][j], mem[i-1][j-weights[i]]+values[i])
			}
		}
	}

	// The outer is weight
	// for j := 0; j <= size; j++ {
	// 	for i := 1; i < n; i++ {
	// 		if j < weights[i] {
	// 			mem[i][j] = mem[i-1][j]
	// 		} else {
	// 			mem[i][j] = max(mem[i-1][j], mem[i-1][j-weights[i]]+values[i])
	// 		}
	// 	}
	// }

	return mem[n-1][size]
}

func ZeroOnePackagePro(weights []int, values []int, size int) int {
	n := len(weights)
	mem := make([]int, size+1)

	for i := 0; i < n; i++ {
		// In order to avoid repeat computing
		// For example: weights[0]=1, values[0]=15
		// mem[1] = max(mem[1], mem[0]+15) = 15
		// mem[2] = max(mem[2], mem[1]+15) = 30...
		// The item0 is repeat calculated
		for j := size; j >= weights[i]; j-- {
			mem[j] = max(mem[j], mem[j-weights[i]]+values[i])
		}
	}
	return mem[size]
}

package homework

func reverse(input []int64) (result []int64) {
	for a, b := 0, len(input)-1; a < b; a, b = a+1, b-1 {
		input[a], input[b] = input[b], input[a]
	}
	
	return input
}

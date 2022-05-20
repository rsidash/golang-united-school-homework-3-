package homework

func average(input [15]float32) (result float32) {
	var sum float32 = 0;

	for _, value := range input {
		sum += float32(value);
	}

	result = sum / float32(len(input));

	return result;
}

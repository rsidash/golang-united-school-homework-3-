package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	keys := []int{}

	for key := range input {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, value := range input {
		result = append(result, value)
	}

	return result
}

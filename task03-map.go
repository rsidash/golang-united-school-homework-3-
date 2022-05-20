package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	keys := make([]int, 0, len(input))

	for key := range input {
		keys = append(keys, key)
	}

	sort.Ints(keys) //sort keys alphabetically

	for _, value := range keys {
		result = append(result, input[value])
	}

	return result
}

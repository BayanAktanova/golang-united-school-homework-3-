package homework

import (
    "sort"
    
)

func sortMapValues(input map[int]string) (result []string) {
	result := make([]string, 0, len(input))

    for k := range input{
        result = append(result, k)
    }
    sort.Strings(result)
  	return result
}

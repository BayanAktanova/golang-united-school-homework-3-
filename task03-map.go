package homework

import (
    "fmt"
	"sort"
    
)

func sortMapValues(input map[int]string) (result []string) {
	keys := make([]int, 0, len(input))
  
    for k := range input{
        keys = append(keys, k)
    }
    sort.Ints(keys)
  
    return keys
}

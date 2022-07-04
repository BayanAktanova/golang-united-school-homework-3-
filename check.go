package main

import (
    "fmt"
	"sort"
    
)

func main() {
	basket := map[int]string{2: "a", 0: "b", 1: "c"}
	
	keys := make([]int, 0, len(basket))
  
    for k := range basket{
        keys = append(keys, k)
    }
    sort.Ints(keys)
  
    for _, k := range keys {
        fmt.Println(k, basket[k])
    }

}
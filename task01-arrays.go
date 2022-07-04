package homework

func average(input [15]float32) (result float32) {
	n := len(input)
     
    sum := 0
 
    for i := 0; i < n; i++ {
 
        sum += (array[i])
    }
     
    avg := (float32(sum)) / (float32(n))
     
    return avg
}

package homework

func average(input [15]float32) (result float32) {
	n, sum float32
	
	n := len(input)
     
    sum := 0
 
    for i := 0; i < n; i++ {
 
        sum += float32(input[i])
    }
     
    avg := sum/n
     
    return avg
}

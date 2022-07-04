package homework

func average(input [15]float32) (result float32) {
	n := len(input)
     
    sum := 0
 
    for i := 0; i < n; i++ {
 
        sum += (input[i])
    }
     
    avg := sum/n
     
    return avg
}

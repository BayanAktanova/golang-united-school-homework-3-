package homework

func average(input [15]float32) (result float32) {
	array := []int{1, 2, 3, 4}
     
    // size of the array
    n := len(input)
     
    // declaring a variable
    // to store the sum 
    sum := 0
 
    // traversing through the
    // array using for loop
    for i := 0; i < n; i++ {
 
        // adding the values of
        // array to the variable sum
        sum += (array[i])
    }
     
    // declaring a variable
    // avg to find the average
    avg := (float32(sum)) / (float32(n))
     
    // typecast all values to float
    // to get the correct result
    //fmt.Println("Sum = ", sum, "\nAverage = ", avg)

	return avg
}

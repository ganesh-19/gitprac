1) "PASSING ARRAY TO FUNCTION:"
	
	In Go language, you are allowed to pass an array as an argument in the function. 
	For passing an array as an argument in the function you have to first create a formal parameter using 
	the following syntax: 

		Syntax:  

		// For sized array
		func function_name(variable_name [size]type){
		// Code
		}

	a) as the way of declaring array is array_name   [size] type -- parameters are declared the same way 

	Using this syntax you can pass 1 or multiple dimensional arrays to the function.

2) "CODE":

	// here the size of the array is mentioned beforehand [...] does not work......
	// unable to return float as summ/size gives an integer .. how to get float
	func average_score(score [4]int, size int) int{
		
		var summ int = 0
		for i:=0; i<size; i++ {
			summ += score[i]
		}
		
		return summ/size
	}
			
	func main() {
		
		// ellipsis does not work for multi dimensional arrays
		runs_1 := [...] int {10,20,30,40}
		
		size := len(runs_1)
		
		average := average_score(runs_1, size)
			
		fmt.Println("average score is:", average)
	}



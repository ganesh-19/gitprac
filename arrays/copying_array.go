1) "COPYING ARRAYS":

	If we create a copy of an array by value and made some changes in the values of the original array, 
	then it will not reflect in the copy of that array. And if we create a copy of an array by reference 
	and made some changes in the values of the original array, then it will reflect in the copy of that array.

2) "SYNTAX":

		// creating a copy of an array by value
		arr := arr1

		// Creating a copy of an array by reference
		arr := &arr1

3) "CODE":

		func main() {
		
			// ellipsis does not work for multi dimensional arrays
			runs_1 := [...] int {10,20,30,40}
			
			//intialise runs_2 with runs_1 
			//creating a copy
			//copy by value
			runs_2 := runs_1
			fmt.Println(runs_1)
			fmt.Println(runs_2)	
			
			fmt.Println("copy by value and value changed")
			runs_2[1] = 56
			fmt.Println(runs_2)
			fmt.Println(runs_1)
			
			// copy by reference
			fmt.Println("copy be reference... runs_3 and runs_1")
			runs_3 := &runs_1
			runs_3[1] = 56
			fmt.Println(runs_3)
			fmt.Println(runs_1)
		
		}


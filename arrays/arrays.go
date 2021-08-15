1) ARRAY:

	In the program, sometimes we need to store a collection of data of the same type, like a list of student marks. 
	Such type of collection is stored in a program using an Array. An array is a fixed-length sequence that is used 
	to store homogeneous elements in the memory. Due to their fixed length array are not much popular like Slice 
	in Go language.

	In an array, you are allowed to store zero or more than zero elements in it. 
	The elements of the array are indexed by using the [] index operator with their zero-based position, 
	means the index of the first element is array[0] and the index of the last element is array[len(array)-1].


2) CREATING AND ACCESSING AN ARRAY:

	In Go language, arrays are created in two different ways:

3) Using var keyword: 
    	
	In Go language, an array is created using the var keyword of a particular type with name, size, and elements. 

	SYNTAX:

		var array_name[length]Type
		or
		var array_name[length]Type{item1, item2, item3, ...itemN}

	Important Points:

	    a) In Go language, arrays are mutable, so that you can use array[index] syntax to the left-hand side 
	       of the assignment to set the elements of the array at the given index.

	    	Var array_name[index] = element

	    b)  You can access the elements of the array by using the index value or by using for loop.
		c) 	In Go language, the array type is one-dimensional.
		d) 	The length of the array is fixed and unchangeable.
		e)	You are allowed to store duplicate elements in an array.


	[size] type is the actual description of an array 

	it is actually var array_name   [size]type

	not var array_name[size]    type

	array_name := [size] type 

4) CODE:

	package main

	import (
		"fmt"
	)

	func main() {
		var runs[10] int 
		
		// showing error
		// var runs_1 [10] int {10, 23, 34, 56, 67, 35, 100, 89, 77, 120}
		
		runs_1 := [10] int {10, 23, 34, 56, 67, 35, 100, 89, 77, 120}
		
		for i:=0; i<(len(runs)); i++{
			runs[i] = runs_1[i]
			fmt.Println(runs)
		}
	}


5) USING VAR KEYWORD:
	

		package main

		import (
			"fmt"
		)

		func main() {
			var runs[10] int 
			
			// showing error
			// var runs_1 [10] int {10, 23, 34, 56, 67, 35, 100, 89, 77, 120}
			
			runs_1 := [10] int {10, 23, 34, 56, 67, 35, 100, 89, 77, 120}
			
			for i:=0; i<(len(runs)); i++{
				runs[i] = runs_1[i]
				fmt.Println(runs)
			}
		}


6) MULTI DIMENSIONAL ARRAY:

	SYNTAX: 

		Array_name[Length1][Length2]..[LengthN]Type

	You can create a multidimensional array using Var keyword or using shorthand declaration


	Note: 
		In multi-dimension array, if a cell is not initialized with some value by the user, 
		then it will initialize with zero by the compiler automatically. 
		There is no uninitialized concept in the Golang.


		func main() {
			var runs[10] int 
			
			// arrays are always initialsed as zero
			fmt.Println(runs)
		}


	FINDING LENGTH OF MULTI DIMENSIONAL ARRAYS:

		func main() {
	
			// multi dimensional arrays
			runs_1 := [3][3] int {{10, 23, 34}, {56, 67, 35}, {100, 89, 77},}
			
			runs_2 := [3][4] int {{10, 23, 34, 52}, {56, 67, 35, 78}, {100, 89, 77, 92},}
			fmt.Println(len(runs_1))
			fmt.Println(len(runs_2[1]))
		}


	ACCESSING ELEMENTS:

		func main() {
	
			// multi dimensional arrays
			runs_1 := [3][3] int {{10, 23, 34}, {56, 67, 35}, {100, 89, 77},}
			
			runs_2 := [3][4] int {{10, 23, 34, 52}, {56, 67, 35, 78}, {100, 89, 77, 92},}
			fmt.Println(len(runs_1))
			fmt.Println(len(runs_2[1]))

			for i:=0; i<(len(runs_2)); i++{
			for j:=0; j<(len(runs_2[1])); j++{
				
				fmt.Println(runs_2[i][j])	
			}
			}
		}


7) IMPORTANT OBSERVATIONS ABOUT ARRAY:

	a) In an array, if an array does not initialized explicitly, then the default value of this array is 0

		func main() {
			var runs[10] int 
			
			// arrays are always initialsed as zero
			fmt.Println(runs)
		}

	b) In an array, you can find the length of the array using len() method as shown below

		func main() {
	
			// multi dimensional arrays
			runs_1 := [3][3] int {{10, 23, 34}, {56, 67, 35}, {100, 89, 77},}
			
			runs_2 := [3][4] int {{10, 23, 34, 52}, {56, 67, 35, 78}, {100, 89, 77, 92},}
			fmt.Println(len(runs_1))
			fmt.Println(len(runs_2[1]))

			for i:=0; i<(len(runs_2)); i++{
			for j:=0; j<(len(runs_2[1])); j++{
				
				fmt.Println(runs_2[i][j])	
			}
			}
		}

	c) In an array, if ellipsis ‘‘…’’ become visible at the place of length, then the length of the array is determined by the initialized elements


		func main() {
	
			// ellipsis does not work for multi dimensional arrays
			runs_1 := [...] int {10,20,30,40}
			
			runs_2 := [...] int {10, 20,20}
			fmt.Println(len(runs_1))
			fmt.Println(len(runs_2))	
		}

	d) In an array, you are allowed to iterate over the range of the elements of the array.

	e) In Go language, an array is of value type not of reference type. So when the array is assigned to a new variable, 
	   then the changes made in the new variable do not affect the original array.

	    func main() {
			
			// ellipsis does not work for multi dimensional arrays
			runs_1 := [...] int {10,20,30,40}
			
			//intialise runs_2 with runs_1 
			//creating a copy
			
			runs_2 := runs_1
			fmt.Println(runs_1)
			fmt.Println(runs_2)	
			
			runs_2[1] = 56
			fmt.Println(runs_2)
			fmt.Println(runs_1)
			
			runs_1[2] = 50
			fmt.Println(runs_2)
			fmt.Println(runs_1)
		}

	f) In an array, if the element type of the array is comparable, 
	   then the array type is also comparable. So we can directly compare two arrays using == operator. 


	   	func main() {
	
			// ellipsis does not work for multi dimensional arrays
			runs_1 := [...] int {10,20,30,40}
			runs_2 := [...] int {10,20,30,40}
			runs_3 := [...] int {10,50,50,40}
			
			//intialise runs_2 with runs_1 
			//creating a copy
			
			fmt.Println(runs_1 == runs_2)
			fmt.Println(runs_2 == runs_3)
		}

append funciton:
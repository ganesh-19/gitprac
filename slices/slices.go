1) "SLICES":

	Slice is more powerful, flexible, convenient than an array, and is a lightweight data structure. 
	Slice is a variable-length sequence which stores elements of a similar type, 
	you are not allowed to store different type of elements in the same slice. 
	It is just like an array having an index value and length, but the size of the slice is resized 
	they are not in fixed-size just like an array. Internally, slice and an array are connected with each other, 
	a slice is a reference to an underlying array. It is allowed to store duplicate elements in the slice. 
	The first index position in a slice is always 0 and the last one will be (length of slice – 1).

2) "DECLARAITON OF SLICE":

	A slice is declared just like an array, but it doesn’t contain the size of the slice. 
	So it can grow or shrink according to the requirement. 

	"SYNTAX":

		slice_name []type
		or 
		slice_name []type{}
		or 
		slice_name []type{value1, value2, value3, ...value n}

		example:
			var my_slice[]int  or var my_slice[]int{}

3) "COMPONENTS OF A SLICE:"
	
	A slice contains three components:

	    Pointer: The pointer is used to points to the first element of the array that is accessible through the slice. 
	    		  Here, it is not necessary that the pointed element is the first element of the array.
	    Length: The length is the total number of elements present in the array.
	    Capacity: The capacity represents the maximum size upto which it can expand.

	"EXAMPLE":

		func main() {
	
			runs_1 := [...] int {10,20,30,40, 50, 60}
			
			// intialising slice -- not neccessary if := is used directly on the slicing statement
			run_slice := []int{}
			
			//slicing statement
			run_slice = runs_1[2:4]
			
			// other way
			runs_2 := runs_1[1:3]
			
			fmt.Println(run_slice)
			fmt.Println(runs_2)
		}


4) "EXPLANATION":

	In the above example, we create a slice from the given array. 
	Here the pointer of the first slice pointed to index 2 because the lower bound of the slice is set to 2
	so it starts accessing elements from index 2 The length of the slice is 3, 
	which means the total number of elements present in the slice is 3 
	and the capacity of the slice 6 means it can store a maximum of 6 elements in it which is the size of runs_1.


5) "HOW TO CREATE AND INITIALISE A SLICE":

5a) "USING SLICE LITERAL":

	    You can create a slice using the slice literal. The creation of slice literal is just like an 
	    array literal, but with one difference you are not allowed to specify the size of the slice in the 
	    square braces[]. As shown in the below example, the right-hand side of this expression is the 
	    slice literal.

		var my_slice_1 = []string{"Geeks", "for", "Geeks"}

		Note: 
		  Always remember when you create a slice using a string literal, 
		  then it first creates an array and after that return a slice reference to it.


		"CODE":


			// short hand declaration
			runs_1 := []int{10,20,30,40, 50, 60}
			fmt.Println("slice is:", runs_1)
			
			//long hand declaration
			
			var runs_2 = [] int {10,20,30,40, 50, 60}
			fmt.Println("slice 2 is:", runs_2)	

			/* does not work
			runs_1[7] = 90
			fmt.Println("slice after new element is added:", runs_1)
			*/

	"WE CAN ADD ELEMENTS TILL THE CAPACITY IS REACHED:"

		var make_slice = make([]int, 4,7)
		fmt.Printf("slice is %v, length is %d, capacity is %d", make_slice, len(make_slice), cap(make_slice))
		
		// we can add elements till the capacity is reached
		make_slice = []int{1,2,3,4,5}
		fmt.Printf("\n slice is %v, length is %d, capacity is %d", make_slice, len(make_slice), cap(make_slice))
		fmt.Println("\n",make_slice)


5b) "USING AN ARRAY":

	    As we already know that the slice is the reference of the array so you can create a slice 
	    from the given array. For creating a slice from the given array first you need to specify 
	    the lower and upper bound, which means slice can take elements from the array starting from 
	    the lower bound to the upper bound. It does not include the elements above from the upper bound. 
	    As shown in the below example:

		'Syntax:'' 

			array_name[low:high]

		This syntax will return a new slice.

		"CODE":

			func main() {

				runs_1 := [...] int {10,20,30,40, 50, 60}
				
				// intialising slice -- not neccessary if := is used directly on the slicing statement
				run_slice := []int{}
				
				//slicing statement
				run_slice = runs_1[2:4]
				
				// other way
				runs_2 := runs_1[1:3]
				
				fmt.Println(run_slice)
				fmt.Println(runs_2)
			}

5c) "USING MAKE FUNCTION":

	    You can also create a slice using the make() function which is provided by the go library. 
	    This function takes three parameters, i.e, type, length, and capacity. 
	    Here, capacity value is optional. It assigns an underlying array with a size that is equal 
	    to the given capacity and returns a slice which refers to the underlying array. 
	    Generally, make() function is used to create an empty slice. 
	    Here, empty slices are those slices that contain an empty array reference.
		
		Syntax: 

			func make([]T, len, cap) []T

		"CODE":

			var make_slice = make([]int, 4,7)
			fmt.Printf("slice is %v, length is %d, capacity is %d", make_slice, len(make_slice), cap(make_slice))
			
			var make_slice_2 = make([]int, 7,7)
			fmt.Printf("slice is %v, length is %d, capacity is %d", make_slice_2, len(make_slice_2), cap(make_slice_2))

6) "HOW TO ITERATE OVER A SLICE":
	

6a) "USING FOR LOOP":
	
	It is the simplest way to iterate slice.

	"CODE:"

		// we can add elements till the capacity is reached
		make_slice := []int{1,2,3,4,5}
		
		// iterating over a slice:
		for i:=0; i<len(make_slice); i++ {
			
			fmt.Println(make_slice[i])
		}

6b) "USING RANGE IN FOR LOOP":
	
		// we can add elements till the capacity is reached
		make_slice := []int{1,2,3,4,5}
		// iterating over a slice:
		for index, element := range make_slice {
			fmt.Printf("\n index is: %d, element is: %d", index, element)
		}
		

6c) "USING A BLANK IDENTIFIER IN FOR LOOP:"
	
		In the range for loop, if you don’t want to get the index value of the elements then you can use 
		blank space(_) in place of index variable

			// we can add elements till the capacity is reached
			make_slice := []int{1,2,3,4,5}
			// iterating over a slice:
			for _, element := range make_slice {
				fmt.Printf("\nelement is: %d",element)
			}

7) "IMPORTANT POINTS ABOUT SLICE:"
	

7a) 'ZERO VALUE SLICE:'
	
	you are allowed to create a nil slice that does not contain any element in it. 
	So the capacity and the length of this slice is 0 
	Nil slice does not contain an array reference

	'CODE:'

		// we cannot initialise the zero value slice like these
		// make_slice := []int
		// var make_slice = []int 
		
		// correct way
		var make_slice []int
		fmt.Printf("length is %d capacity is %d", len(make_slice), cap(make_slice))

7b) 'MODIFYING SLICE:'
	
	As we already know that slice is a reference type it can refer an underlying array. 
	So if we change some elements in the slice, then the changes should also take place in the referenced array. 
	Or in other words, if you made any changes in the slice, then it will also reflect in the array.

	'CODE:'
		// declaring array
		runs_1 := [6]int{10,20,30,40, 50, 60}
		
		// making a slice
		runs_slice := runs_1[1:4]
		
		fmt.Printf("Original array: %v, Original slice: %v", runs_1, runs_slice)
		
		// modification
		runs_slice[1] = 50
		runs_slice[2] = 0
		fmt.Printf("\narray after slice modification: %v slice: %v", runs_1, runs_slice)

7c) 'COMPARISION OF SLICE:'
	
	In Slice, you can only use == operator to check the given slice is nill or not. 
	If you try to compare two slices with the help of == operator then it will give you an error.

	'CODE:'

		// declaring array
		runs_1 := [6]int{10,20,30,40, 50, 60}
		
		// making a slice
		slice1 := runs_1[0:3]
		slice2 := []int{10,20,30, 40}
		slice3 := []int{}
		
		/* invalid operation: slice1 == slice2 (slice can only be compared to nil)
		fmt.Println(slice1 == slice2)
		*/
		
		fmt.Println(slice1 == nil)
		fmt.Println(slice2 == nil)
		fmt.Println(slice3 == nil)

7d) 'MULTI DIMENSIONAL SLICE:'
	
	Multi-dimensional slice are just like the multidimensional array, except that slice does not contain the size. 

	Two ways of creating a multidimensional slice

		// Creating multi-dimensional slice
	    s1 := [][]int{{12, 34},
	        {56, 47},
	        {29, 40},
	        {46, 78},
	    }
	    // Accessing multi-dimensional slice
	    fmt.Println("Slice 1 : ", s1)

	    s2 := [][] int { []int{12, 34},
	        [] int {56, 47},
	        [] int {29, 40},
	        [] int {46, 78},
	    }
	    fmt.Println("Slice 2 : ", s2)

7e) 'SORTING OF SLICE:'

	you are allowed to sort the elements present in the slice. 
	The standard library of Go language provides the sort package which contains different types of sorting methods 
	for sorting the slice of ints, float64s, and strings. 
	These functions always sort the elements available is slice in ascending order.

		'CODE:'

			import (
			"fmt"
			"sort"
		)
				
		func main() {
			
		    // Creating multi-dimensional slice
		    s1 := [][]int{{12, 34},
		        {56, 47},
		        {29, 40},
		        {46, 78},
		    }
		    // Accessing multi-dimensional slice
		    fmt.Println("Slice 1 : ", s1)

		    slc1 := []string{"Python", "Java", "C#", "Go", "Ruby"}
		    slc2 := []int{45, 67, 23, 90, 33, 21, 56, 78, 89}

		    //sorting of slice 
			sort.Ints(slc2)
			sort.Strings(slc1)
			
			/* cannot use s1 (type [][]int) as type []int in argument to sort.Ints
			sort.Ints(s1)
			*/
			
			fmt.Println("Sorted slice:", slc1, slc2)


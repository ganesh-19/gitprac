link = https://www.geeksforgeeks.org/how-to-sort-a-slice-of-ints-in-golang/#main


1) "SORT:"
	
	Go language allows you to sort the elements of the slice according to its type. 
	So, an int type slice is sorted by using the following functions. 
	These functions are defined under the sort package so, you have to import sort package.

2) "Ints:"
	
	This function is used to only sorts a slice of ints and it sorts the elements of the slice in increasing order.

	Syntax:

		func Ints(slc []int)

	slc - represents the slice of ints 

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

3) 'Ints are sorted / Strings are sorted:'

	This function is used to check whether the given slice of ints is in the sorted form(in increasing order) or not. 
	This method returns true if the slice is in sorted form, or return false if the slice is not in the sorted form.
	
	Syntax:

		func IntsAreSorted(scl []int) 
		func StringsAreSorted(scl []string)
	
	'CODE:'

		func main() {
		
	    slc1 := []string{"Python", "Java", "C#", "Go", "Ruby"}
	    slc2 := []int{45, 67, 23, 90, 33, 21, 56, 78, 89}

	    //sorting of slice 
		sort.Ints(slc2)
		sort.Strings(slc1)
		
		/* cannot use s1 (type [][]int) as type []int in argument to sort.Ints
		sort.Ints(s1)
		*/
		res1 := sort.StringsAreSorted(slc1)
		res2 := sort.IntsAreSorted(slc2)
		
		fmt.Println("Sorted slice:", slc1, slc2)
		fmt.Println("Issorted:", res1, res2)
}










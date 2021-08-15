link = https://www.geeksforgeeks.org/how-to-trim-a-slice-of-bytes-in-golang/#main

1) "TRIM SLICE OF BYTES:"
	
	you are allowed to trim all the leading and trailing UTF-8-encoded code points from the given slice using Trim() 
	function. This function returns a subslice of the original slice by slicing off all leading and trailing 
	UTF-8-encoded code points which are specified in the given string. If the given slice of bytes does not contain 
	the specified string in it, then this function returns the original slice without any change. 
	It is defined under the bytes package so, you have to import bytes package in your program for accessing 
	Trim function.

2) Syntax:

	func Trim(ori_slice[]byte, cut_string string) []byte

	Here, ori_slice is the original slice of bytes and cut_string represent a string which you want to 
	trim in the given slice.

	'CODE:'

		package main

		import (
			"fmt"
			"bytes"
		)
				
		func main() {
			
		    slice_1 := []byte{'!', '!', 'G', 'e', 'e', 'k', 's', 'f',  
		                 'o', 'r', 'G', 'e', 'e', 'k', 's', '#', '#'}
		    slice_2 := []byte{'*', '*', 'A', 'p', 'p', 'l', 'e', '^', '^'}    
		    slice_3 := []byte{'%', 'g', 'e', 'e', 'k', 's', '%'} 
			
		    // start from ! to # --- innermost chars are taken
		    // start from % to ^
		    // take all
		    res1 := bytes.Trim(slice_1, "!#")
		    res2 := bytes.Trim(slice_2, "*^")
		    res3 := bytes.Trim(slice_3, "@") 

		    fmt.Printf("New Slice:\n")
		    fmt.Printf("\nSlice 1: %s", res1)
		    fmt.Printf("\nSlice 2: %s", res2)
		    fmt.Printf("\nSlice 3: %s", res3) 

		}
3) 'CODE:'

	
	import (
	"fmt"
	"bytes"
	)
			
	func main() {
	    
	// instead of slice_name directly slice is used here
	    res1 := bytes.Trim([]byte("****Welcome to GeeksforGeeks****"), "*")
	    res2 := bytes.Trim([]byte("!!!!Learning how to trim a slice of bytes@@@@"), "!@")
	    res3 := bytes.Trim([]byte("^^Geek&&"), "$")
	  
	    // Display the results
	    fmt.Printf("Final Slice:\n")
	    fmt.Printf("\nSlice 1: %s", res1)
	    fmt.Printf("\nSlice 2: %s", res2)
	    fmt.Printf("\nSlice 3: %s", res3) 

	}

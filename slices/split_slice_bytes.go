1) "SPILTTING SLICES OF BYTES:"

	Go slice of bytes, you are allowed to split the given slice using a Split() function. 
	This function splits a slice of bytes into all subslices separated by the given separator and 
	returns a slice which contains all these subslices. It is defined under the bytes package so, 
	you have to import bytes package in your program for accessing Split function.

2) "CODE:"
	
	slice_1 := []byte{'!', '!', 'G', 'e', 'e', 'k', 's', 
       'f', 'o', 'r', 'G', 'e', 'e', 'k', 's', '#', '#'}
      
    slice_2 := []byte{'A', 'p', 'p', 'l', 'e'}
      
    slice_3 := []byte{'%', 'g', '%', 'e', '%', 'e', 
                           '%', 'k', '%', 's', '%'}
  
    // Displaying slices
    fmt.Println("Original Slice:")
    fmt.Printf("Slice 1: %s", slice_1)
    fmt.Printf("\nSlice 2: %s", slice_2)
    fmt.Printf("\nSlice 3: %s", slice_3) 

    //splitting slice of bytes -- using Split function

    res1 := bytes.Split(slice_1, []byte("eek"))
    res2 := bytes.Split(slice_2, []byte(""))
    res3 := bytes.Split(slice_3, []byte("%"))
    
      // Display the results
    fmt.Printf("\n\nAfter splitting:")
    fmt.Printf("\nSlice 1: %s", res1)
    fmt.Printf("\nSlice 2: %s", res2)
    fmt.Printf("\nSlice 3: %s", res3) 

    



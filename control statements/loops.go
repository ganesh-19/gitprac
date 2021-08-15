1) LOOP:
	Go language contains only a single loop that is for-loop. 
	A for loop is a repetition control structure that allows us to write a loop that is executed a specific number of times. 
	In Go language, this for loop can be used in the different forms

2) SIMPLE FOR LOOP:

	It is similar that we use in other programming languages like C, C++, Java, C#, etc.

	Syntax:

		for initialization; condition; post{
		       // statements....
		}

	    The initialization statement is optional and executes before for loop starts. The initialization statement is always in a simple statement like variable declarations, increment or assignment statements, or function calls.
	    The condition statement holds a boolean expression, which is evaluated at the starting of each iteration of the loop. If the value of the conditional statement is true, then the loop executes.
	    The post statement is executed after the body of the for-loop. After the post statement, the condition statement evaluates again if the value of the conditional statement is false, then the loop ends.

	Code:

		var x = 2
	
		for i:=0; i<=4; i++{
			if (x == 2){
				fmt.Println("x is 2")
			} else {
				fmt.Println("x has other values")
			}
		}

3) FOR LOOP AS INFINITE LOOP:

	A for loop is also used as an infinite loop by removing all the three expressions from the for loop. 
	When the user did not write condition statement in for loop it means the condition statement is true and the loop goes into an infinite loop.

	Syntax:

	for{
	     // Statement...
	}

	CODE:

	var x = 2
	
	for {
		if (x == 2){
			fmt.Println("x is 2")
		} else {
			fmt.Println("x has other values")
		}
	}


4) FOR LOOP AS A WHILE LOOP:


	This loop is executed until the given condition is true. 
	When the value of the given condition is false the loop ends.

	SYNTAX:

		for condition{
		    // statement..
		}

	CODE:

		var x = 2
		i:=0
		for i<=4 {
			if (x == 2){
				fmt.Println("x is 2")
				i++
			} else {
				fmt.Println("x has other values")
				i++
			}
		}

5) SIMPLE RANGE IN FOR LOOP:

	Using simple range in for loop

	SYNTAX:

		for i, j:= range rvariable{
		   // statement..
		}

		here,

	    i and j are the variables in which the values of the iteration are assigned. They are also known as iteration variables.
	    The second variable, i.e, j is optional.
	    The range expression is evaluated once before the starting of the loop.

	CODE:

		range_array := []string{"goku", "gohan", "Vegeta"}
	
		// i and j stores the value of rvariable
	    	// i store index number of individual string and
	    	// j store individual string of the given array 
		
		for i, j:= range range_array {
			fmt.Println(i,j)
		}

6) USING FOR LOOP FOR STRINGS:

	A for loop can iterate over the Unicode code point for a string.

	SYNTAX:

		for index, chr:= range str{
		     // Statement..
		}

	CODE:

		for index, char:= range "dragon"{
			fmt.Printf("index is %d" + "character is %c \n", index,char)
		}

7) FOR MAPS:

	A for loop can iterate over the key and value pairs of the map.

	Syntax:

	for key, value := range map { 
	     // Statement.. 
	}

	
	CODE:

		mmap := map[int]string{
        22:"Geeks",
        33:"GFG",
        44:"GeeksforGeeks",
	    }
	    for key, value:= range mmap {
	       fmt.Println(key, value) 
	    } 

8) FOR CHANNEL:

	A for loop can iterate over the sequential values sent on the channel until it closed.


	SYNTAX:

		for item := range Chnl { 
		     // statements..
		}

	CODE:


		var chn1 = make(chan int)
	
		 fmt.Println(chn1)
		
		 go func(){
			chn1 <- 100
			chn1 <- 1000
			chn1 <- 10000
			chn1 <- 100000
		} ()
		
		for i:= range chn1 {
			fmt.Println(i)
		}

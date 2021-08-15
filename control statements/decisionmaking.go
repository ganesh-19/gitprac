1) The Decision making statements of Go programming are:

    if Statement
    if-else Statement
    Nested if Statement
    if-else-if Ladder

2) if statement

	Syntax:

	if(condition) {

	   // Statements to execute if
	   // condition is true
	}

	CODE:

		if(v < 1000) {
         
	      // print the following if 
	      // condition evaluates to true
	      fmt.Printf("v is less than 1000\n")
	   } 

3) IF .... ELSE 

	if (condition) {

	    // Executes this block if
	    // condition is true
	} else {

	    // Executes this block if
	    // condition is false
	}

4) NESTED IF STATEMENT:

	if (condition1) {

	   // Executes when condition1 is true
	   
	   if (condition2) {

	      // Executes when condition2 is true
	   }
	}


	CODE:

	func main() {
	var x,y,z = 2,3,4
		
		if (x!=y){
			fmt.Println(x!=y)
			
			if (y!=z){
			
				fmt.Println(x!=y, y==z)
			}
		}
	}


5) IF ELSE IF .... LADDER


	Important Points:

	    if statement can have zero or one else statements and it must come after any else if statements
	    if statement can have zero to many else if’s and it must come before the else.
	    None of the remaining else if’s or else’s will be tested if an else if succeeds


	SYNTAX:

		if(condition_1) {

		} else if(condition_2) {

		}
		.
		.
		. else {

		}

	CODE:

		if (x==y){
		fmt.Println(x!=y)
		
		if (y!=z){
		
			fmt.Println(x!=y, y==z)
			}
		} else if (x!=y && y==z) {
			fmt.Println("ELSEIF")
		} else {
			fmt.Println("ELSE")
		}






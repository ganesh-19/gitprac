1) VARIABLES

	 When a user enters a new value that will be used in the process of operation, can store temporarily in the Random Access Memory of the computer and these values in this part of memory vary throughout the execution and hence another term for this came which is known as Variables. 
	 So basically, a Variable is a placeholder of the information which can be changed at runtime. 
	 And variables allow to Retrieve and Manipulate the stored information.


2) Rules for Naming Variables:

    Variable names must begin with a letter or an underscore(_). And the names may contain the letters ‘a-z’ or ’A-Z’ or digits 0-9 as well as the character ‘_’.

    Geeks, geeks, _geeks23  // valid variable
    123Geeks, 23geeks      // invalid variable

    A variable name should not start with a digit.

    234geeks  // illegal variable 

    The name of the variable is case sensitive.

    geeks and Geeks are two different variables

    Keywords is not allowed to use as a variable name.
    There is no limit on the length of the name of the variable, but it is advisable to use an optimum length of 4 – 15 letters only.

3) Declaring a Variable

	In Go language variables are created in two different ways

	1) USING A VAR KEYWORD:

		Syntax:

			var variable_name type = expression

		a) If the type is removed, then the type of the variable is determined by the value-initialize in the expression.


		b) If the expression is removed, then the variable holds zero-value for the type like zero for the number, false for Booleans, “” for strings, and nil for interface and reference type. 
		   So, there is no such concept of an uninitialized variable in Go language.

			CODE:

			package main
   
			import "fmt"
			   
			func main() {
			  
			    // Variable declared and 
			    // initialized without expression
			    var myvariable1 int
			    var myvariable2 string
			    var myvariable3 float64
			  
			    // Display the zero-value of the variables
			    fmt.Printf("The value of myvariable1 is : %d\n",
			                                     myvariable1)
			  
			    fmt.Printf("The value of myvariable2 is : %s\n",
			                                     myvariable2)
			  
			    fmt.Printf("The value of myvariable3 is : %f",
			                                     myvariable3)
			} 

			prints default values


		c) multiple variables can be declared at once

		   if type not declared then default type for the value is taken

			package main

			import (
				"fmt"
			)

			func main() {
				var str1, str2= "goku", "gohan"
				check := str1 == str2
				
				fmt.Println(str1, str2, check)
			}

		d) You are allowed to initialize a set of variables by the calling function that returns multiple values. 

			i,j,... type = function(args, kwargs)


	2) Using short variable declaration: 

		The local variables which are declared and initialize in the functions are declared by using short variable declaration.

		Syntax:

			variable_name:= expression


		Important Points:

			a) the type of the variable is determined by the type of the expression. 

			b) Most of the local variables are declared and initialized by using short variable declarations due to their brevity and flexibility.
				The var declaration of variables are used for those local variables which need an explicit type that differs from the initializer expression, or for those variables whose values are assigned later and the initialized value is unimportant.
				
			c) Using short variable declaration you are allowed to declare multiple variables in the single declaration. 


				myvar1, myvar2, myvar3 := 800, 34, 56 


			d) n a short variable declaration, you are allowed to initialize a set of variables by the calling function that returns multiple values. 

				i,j ... := function(args, kwargs)


			e) A short variable declaration acts like an assignment only when for those variables that are already declared in the same lexical block. 
			   The variables that are declared in the outer block are ignored. 
			   And at least one variable is a new variable out of these two variables as shown in the below example. 


			   	// as an assignment for myvar2 variable
				// because same variable present in the same block
				// so the value of myvar2 is changed from 45 to 100
				myvar1, myvar2 := 39, 45 
				myvar3, myvar2 := 45, 100
				  
				// here compiler will gives error because
				// these variables are already defined
				myvar1, myvar2 := 43, 47
				myvar2:= 200 

				CODE:

				func main() {
				myvar1, myvar2 := 39, 45 
				myvar3, myvar2 := 45, 100
				fmt.Println(myvar1, myvar2, myvar3)
				
				// ERROR
				// ./prog.go:12:17: no new variables on left side of :=
				// ./prog.go:13:8: no new variables on left side of :=
				myvar1, myvar2 := 43, 47
				myvar2:= 200 

				// following works
				// as new variable is getting added everytime
				myvar1, myvar2, myvar4 := 43, 47, 23
				myvar2, myvar6 := 200 , 100
				
			f) Using short variable declaration you are allowed to declare multiple variables of different types in the single declaration. 
				The type of these variables are determined by the expression. 

				example:

				myvar1, myvar2, myvar3 := 800, "Geeks", 47.56 

				



		
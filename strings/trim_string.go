1. "HOW TO TRIM A STRING IN GOLANG:"
	In Go language, strings are different from other languages like Java, C++, Python, etc. 
	It is a sequence of variable-width characters where each and every character is represented by 
	one or more bytes using UTF-8 Encoding. You can trim a string in different ways using the 
	following list of functions. All these functions are defined under strings package, 
	so you have to import strings package in your program to access these functions.

2. 'TRIM':

	This function is used to trim the string 
	all the leading and trailing Unicode code points which are specified in this function.

	'SYNTAX:'
		func Trim(str string, cutstr string) string

		Here, str represent the current string and cutstr represents the elements which you want to 
		trim in the given string.

	'CODE:'
		package main

		import (
			"fmt"
			"strings"
		)
		
		var name = "bardock"
		
		func main() {
		
			// Creating and initializing string
			// Using shorthand declaration
			str1 := "!!Welcome to GeeksforGeeks !!"
			str2 := "@@This is the tutorial of Golang$$"
			str3 := "dragonballz"
		
			// Displaying strings
			fmt.Println("Strings before trimming:")
			fmt.Println("String 1: ", str1)
			fmt.Println("String 2:", str2)
			fmt.Println("String 3:", str3)
		
			// Trimming the given strings
			// Using Trim() function
			res1 := strings.Trim(str1, "!")
			res2 := strings.Trim(str2, "@$")
			res3 := strings.Trim(str3, "l")
			res4 := strings.Trim(str3, "a")
		
			// Displaying the results
			fmt.Println("\nStrings after trimming:")
			fmt.Println("Result 1: ", res1)
			fmt.Println("Result 2:", res2) 
			fmt.Println("Res3 and 4:", res3, res4)
		}

3. 'TRIMLEFT':
	This function is used to trim the left-hand side(specified in the function) 
	Unicode code points of the string.

	'Syntax':
		func TrimLeft(str string, cutstr string) string

		Here, str represent the current string and cutstr represents the left-hand side elements 
		which you want to trim in the given string. 
		
	'CODE':
		package main

		import (
			"fmt"
			"strings"
		)
		var name = "bardock"
		
		func main() {
		
			// Creating and initializing string
			// Using shorthand declaration
			str1 := "!!Welcome to GeeksforGeeks !!"
			str2 := "@@This is the tutorial of Golang$$"
			str3 := "dragonballz"
		
			// Displaying strings
			fmt.Println("Strings before trimming:")
			fmt.Println("String 1: ", str1)
			fmt.Println("String 2:", str2)
			fmt.Println("String 3:", str3)
		
			// Trimming the given strings
			// Using Trim() function
			res1 := strings.TrimLeft(str1, "!")
			res2 := strings.TrimLeft(str2, "@$")
			res3 := strings.TrimLeft(str3, "ll")
			res4 := strings.TrimLeft(str3, "an")
			
			res5 := strings.TrimRight(str1, "!")
			res6 := strings.TrimRight(str2, "@$")
			res7 := strings.TrimRight(str3, "ll")
			res8 := strings.TrimRight(str3, "an")
		
			// Displaying the results
			fmt.Println("\nStrings after trimming:")
			fmt.Println("Result 1: ", res1)
			fmt.Println("Result 2:", res2) 
			fmt.Println("Res3 and 4:", res3, res4)
		
			fmt.Println(res5, res6, res7, res8)

		}

4. "TRIM RIGHT:"

5. 'TRIM SPACE': 
	This function is used to trim all the leading and trailing white space from the specified string.

6. "TrimSuffix":
	This method is used to trim the trailing suffix string from the given string. 
	If the given string does not contain the specified suffix string, then this function returns 
	the original string without any change.

	'Syntax':
		
		func TrimSuffix(str, suffstr string) string
		Here, str represents the original string and suffstr represent the suffix string.
	
	'Example':

		// Go program to illustrate how to
		// trim a suffix string from the
		// given string
		package main
		
		import (
			"fmt"
			"strings"
		)
		
		// Main method
		func main() {
		
			// Creating and initializing string
			// Using shorthand declaration
			str1 := "Welcome, GeeksforGeeks"
			str2 := "This is the, tutorial of Golang"
		
			// Displaying strings
			fmt.Println("Strings before trimming:")
			fmt.Println("String 1: ", str1)
			fmt.Println("String 2:", str2)
		
			// Trimming suffix string from the given strings
			// Using TrimSuffix() function
			res1 := strings.TrimSuffix(str1, "GeeksforGeeks")
			res2 := strings.TrimSuffix(str2, "Hello")
		
			// Displaying the results
			fmt.Println("\nStrings after trimming:")
			fmt.Println("Result 1: ", res1)
			fmt.Println("Result 2:", res2)
		}
		
		Output:
		
		Strings before trimming:
		String 1:  Welcome, GeeksforGeeks
		String 2: This is the, tutorial of Golang
		
		Strings after trimming:
		Result 1:  Welcome, 
		Result 2: This is the, tutorial of Golang
	
	 
	
7. "TrimPrefix": 
	This method is used to trim the leading prefix string from the given string. 
	If the given string does not contain the specified prefix string, then this function returns 
	the original string without any change.
	
	'Syntax':
	
		func TrimPrefix(str, suffstr string) string
		Here, str represents the original string and suffstr represent the prefix string.
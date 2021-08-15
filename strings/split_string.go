1. "How to split a string in golang?":
	In Go language, strings are different from other languages like Java, C++, Python, etc. 
	It is a sequence of variable-width characters where each and every character is represented by 
	one or more bytes using UTF-8 Encoding. In Go strings, you are allowed to split a string into a 
	slice with the help of following functions. These functions are defined under strings package so, 
	you have to import strings package in your program for accessing these functions.

2.  "Split": 
	This function splits a string into all substrings separated by the given separator and returns 
	a slice which contains these substrings.

	'Syntax':

		func Split(str, sep string) []string

		Here, str is the string and sep is the separator. If str does not contain the given sep and 
		sep is non-empty, then it will return a slice of length 1 which contain only str. Or if the sep is empty, then it will split after each UTF-8 sequence. 
		Or if both str and sep are empty, then it will return an empty slice. 

	'code':
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
			res1 := strings.Split(str1, "!")
			res2 := strings.Split(str2, "@$")
			res3 := strings.Split(str3, "ll")
			res4 := strings.Split(str3, "an")
			

			// Displaying the results
			fmt.Println("\nStrings after trimming:")
			fmt.Println("Result 1: ", res1)
			fmt.Println("Result 2:", res2) 
			fmt.Println("Res3 and 4:", res3, res4)
		}

3. "SPLIT AFTER":

	'CODE':
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
	res1 := strings.SplitAfter(str1, "!")
	res2 := strings.SplitAfter(str2, "@$")
	res3 := strings.SplitAfter(str3, "ll")
	res4 := strings.SplitAfter(str3, "n")


	// Displaying the results
	fmt.Println("\nStrings after trimming:")
	fmt.Println("Result 1: ", res1)
	fmt.Println("Result 2:", res2) 
	fmt.Println("Res3 and 4:", res3, res4)

4. "Split After":
	This function splits a string into all substrings after each instance of the given separator 
	and returns a slice which contains these substrings.

	'Syntax':
		func SplitAfterN(str, sep string, m int) []string

	Here, str is the string, sep is the separator, and m is used to find the number of substrings 
	to return. Here, if m>0, then it returns at most m substrings and the last string substring will 
	not split. If m == 0, then it will return nil. If m<0, then it will return all substrings.

	'code':
		str1 := "Welcome, to the, online portal, of GeeksforGeeks"
		str2 := "My dog name is Dollar"
		str3 := "I like to play Ludo"
	
		// Displaying strings
		fmt.Println("String 1: ", str1)
		fmt.Println("String 2: ", str2)
		fmt.Println("String 3: ", str3)
	
		// Splitting the given strings
		// Using SplitAfterN() function
		res1 := strings.SplitAfterN(str1, ",", 2)
		res2 := strings.SplitAfterN(str2, "", 4)
		res3 := strings.SplitAfterN(str3, "!", 1)
		res4 := strings.SplitAfterN("", "GeeksforGeeks, geeks", 3)
	
		// Displaying the result
		fmt.Println("\nResult 1: ", res1)
		fmt.Println("Result 2: ", res2)
		fmt.Println("Result 3: ", res3)
		fmt.Println("Result 4: ", res4) 
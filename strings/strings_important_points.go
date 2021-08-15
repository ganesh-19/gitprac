1. "STRINGS ARE IMMUTABLE":
    strings are immutable once a string is created the value of the string cannot be changed. 
	Or in other words, strings are read-only. 
	If you try to change, then the compiler will throw an error. 

	'CODE':
		import "fmt"
	
		// Main function
		func main() {
		
			// Creating and initializing a string
			// using shorthand declaration
			mystr := "Welcome to GeeksforGeeks"
		
			fmt.Println("String:", mystr)
		
			/* if you trying to change
				the value of the string
				then the compiler will
				throw an error, i.e,
				cannot assign to mystr[1]
		
			mystr[1]= 'G'
			fmt.Println("String:", mystr)
			*/
		} 

2. "HOW TO ITERATE OVER A STRING":
	can iterate over string using for rang loop. 
	This loop can iterate over the Unicode code point for a string. 

	'CODE:'
	// see the difference between two 
	// %c used for character, same when %d is used it returns number
		import (
			"fmt"
		)
		var name = "bardock"
		func main() {
			for index, letter:= range name{
				fmt.Println("index and letter is:", index, letter)
			}
			for index, letter:= range name{
				fmt.Printf("index and letter is:%d, %c", index, letter)
			}
		}

3. "HOW TO ACCESS THE INDIVIDUAL BYTE OF THE STRING:"\

	The string is of a byte so, we can access each byte of the given string

	'CODE':
		var name = "bardock"

		for index, letter:= range name{
			fmt.Printf("\n index and letter, bytes is:%d, %c, %v", index, letter, letter)
		}
		
		for i:=0; i<len(name); i++{
			fmt.Printf("\n index, letter, bytes are: %d, %c, %v", i, name[i], name[i])
		}

4. "HOW TO CREATE A STRING FROM A SLICE:"
	In Go language, you are allowed to create a string from the slice of bytes. 

		myslice1 := []byte{0x47, 0x65, 0x65, 0x6b, 0x73} 
		myslice2 := []rune{0x0047, 0x0065, 0x0065,0x006b, 0x0073} 
		
		mystring1 := string(myslice1)
		mystring2 := string(myslice2)
		
		fmt.Println("the slices are:", myslice1, myslice2)
		fmt.Println("the strings are:", mystring1, mystring2

5. "HOW TO FIND THE LENGTH OF THE STRING:"

	In Golang string, you can find the length of the string using two functions one is len() 
	and another one is RuneCountInString(). The RuneCountInString() function is provided by 
	UTF-8 package, this function returns the total number of rune presents in the string. 
	And the len() function returns the number of bytes of the string. 

	'CODE':

		import (
			"fmt"
			"unicode/utf8"
		)
		
		var name = "bardock"
		
		func main() {
		
			myslice1 := []byte{0x47, 0x65, 0x65, 0x6b, 0x73} 
			myslice2 := []rune{0x0047, 0x0065, 0x0065,0x006b, 0x0073} 
			
			mystring1 := string(myslice1)
			mystring2 := string(myslice2)
			
			len1 := len(mystring2)
			len2 := utf8.RuneCountInString(mystring1)
			
			fmt.Println("the slices are:", myslice1, myslice2)
			fmt.Println("the strings are:", mystring1, mystring2)
			fmt.Println("the lengths are:", len1, len2)
		}


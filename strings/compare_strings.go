1) "compare":
	In Go language, the string is an immutable chain of arbitrary bytes encoded with UTF-8 encoding. 
	You are allowed to compare strings with each other using two different ways

2) 'Using comparison operators': 
	Go strings support comparison operators, i.e, ==, !=, >=, <=, <, >. 
	Here, the == and != operator are used to check if the given strings are equal or not, and 
	>=, <=, <, > operators are used to find the lexical order. The results of these operators 
	are of Boolean type, meaning if the condition is satisfied it will return true, otherwise, 
	return false.

	"Code":
		str1 := "abc"
		str2 := "abc"
		str3 := "def"
		
		fmt.Println(str1 == str2)
		fmt.Println(str3 > str1)
		fmt.Println(str1 != str2)

3) "Using Compare() method":
	You can also compare two strings using the built-in function Compare() provided by the strings 
	package. This function returns an integer value after comparing two strings lexicographically. 
	The return values are: 

	Return 0, if str1 == str2.
	Return 1, if str1 > str2.
	Return -1, if str1 < str2.

	'code':
		package main

		import (
			"fmt"	
			"strings"
		)
		
		func main() {
		
			str1 := "abc"
			str2 := "abc"
			str3 := "def"
			
			fmt.Println(strings.Compare(str1, str2))
			fmt.Println(strings.Compare(str2, str3))
			fmt.Println(strings.Compare(str1, str3))
		
		}

1) "STRINGS:"
	
	In Go language, strings are different from other languages like Java, C++, Python, etc. 
	it is a sequence of variable-width characters where each and every character is represented by one or more 
	bytes using UTF-8 Encoding. Or in other words, strings are the immutable chain of arbitrary bytes
	(including bytes with zero value) or string is a read-only slice of bytes and the bytes of the strings 
	can be represented in the Unicode text using UTF-8 encoding.

	Due to UTF-8 encoding Golang string can contain a text which is the mixture of any language present in the world, 
	without any confusion and limitation of the page. Generally, strings are enclosed in double-quotes ”.....”

	var name string = "........."
	name := "........"
	var name string 
	name = "........"

2) "String Literals":

	In Go language, string literals are created in two different ways:


2a) "USING DOUBLE QUOTES":
	
	Here, the string literals are created using double-quotes(“”). 
	This type of string support escape character as shown in the below table, 
	but does not span multiple lines. This type of string literals is widely used in Golang programs.


	Escape character 	Description
		\\ 				Backslash(\)
		\000 			Unicode character with the given 3-digit 8-bit octal code point
		\’ 				Single quote (‘). It is only allowed inside character literals
		\” 				Double quote (“). It is only allowed inside interpreted string literals
		\a 				ASCII bell (BEL)
		\b 				ASCII backspace (BS)
		\f 				ASCII formfeed (FF)
		\n 				ASCII linefeed (LF
		\r 				ASCII carriage return (CR)
		\t 				ASCII tab (TAB)
		\uhhhh 			Unicode character with the given 4-digit 16-bit hex code point.
						Unicode character with the given 8-digit 32-bit hex code point.
		\v 				ASCII vertical tab (VT)
		\xhh 			Unicode character with the given 2-digit 8-bit hex code point.


2b) "USING BACKTICKS:"

	Here, the string literals are created using backticks(“) and also known as raw literals. 
	Raw literals do not support escape characters, can span multiple lines, and may contain any character 
	except backtick. It is, generally, used for writing multiple line message, in the regular expressions, 
	and in HTML. 


3) "CODE:"
	
	func main() {
  
    // Creating and initializing a
    // variable with a string literal
    // Using double-quote
    My_value_1 := "Welcome to GeeksforGeeks"
  
    // Adding escape character
    My_value_2 := "Welcome!\nGeeksforGeeks"
  
    // Using backticks
    My_value_3 := `Hello!GeeksforGeeks`
  
    // Adding escape character
    // in raw literals
    My_value_4 := `Hello!\nGeeksforGeeks`
  
    // Displaying strings
    fmt.Println("String 1: ", My_value_1)
    fmt.Println("String 2: ", My_value_2)
    fmt.Println("String 3: ", My_value_3)
    fmt.Println("String 4: ", My_value_4)

	} 


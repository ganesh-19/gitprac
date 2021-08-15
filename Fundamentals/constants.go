1) CONSTANTS:

	once the value of constant is defined it cannot be modified further. 
	There can be any basic data types of constant like an integer constant, a floating constant, a character constant, 
	or a string literal. 

2) How to declare? 

	Constants are declared like variables but in using a const keyword as a prefix to declare constant with a specific type. 
	It cannot be declared using := syntax. 

	Example:

		const PI = 3.14
		const PI := 3.14 -- throws error

3) Untyped and Typed Numeric Constants: 
	Typed constants work like immutable variables can inter-operate only with the same type and 
	untyped constants work like literals can inter-operate with similar types. 
	Constants can be declared with or without a type in Go. 

	Example:

		const untypedInteger          	 = 123
		const untypedFloating typed   	 = 123.12

		const typedInteger int            = 123
		const typedFloatingPoint float64  = 123.12


4) Following is list of constant in Go Language: 

    Numeric Constant (Integer constant, Floating constant, Complex constant)
    String literals
    Boolean Constant

5) Numeric Constant: 

	Numeric constants are high-precision values. 
	As Go is a statically typed language that does not allow operations that mix numeric types. 
	You can’t add a float64 to an int, or even an int32 to an int. 
	Although, it is legal to write 1e6*time.Second or math.Exp(1) or even 1<<(‘\t’+2.0). 
	In Go, constants, unlike variables, behave like regular numbers.

	
	Numeric constant can be of 3 kinds i.e., integer, floating-point, complex 


	5 A) Integer Constant: 

		    A prefix specifies the base or radix: 0x or 0X for hexadecimal, 0 for octal, and nothing for decimal.
		    An integer literal can also have a suffix that is a combination of U(upper case) and L(upper case), 
		    for unsigned and long, respectively
		    It can be a decimal, octal, or hexadecimal constant.
		    An int can store at maximum a 64-bit integer, and sometimes less.

			Following are some examples of Integer Constant: 

			85         /* decimal */
			0213       /* octal */
			0x4b       /* hexadecimal */
			30         /* int */
			30u        /* unsigned int */
			30l        /* long */
			30ul       /* unsigned long */
			212         /* Legal */
			215u        /* Legal */
			0xFeeL      /* Legal */
			078         /* Illegal: 8 is not an octal digit */
			032UU       /* Illegal: cannot repeat a suffix */ 


	5 B) Complex constant: 

			Complex constants behave a lot like floating-point constants. 
			It is an ordered pair or real pair of integer constant( or parameter) 
			And the constant are separated by a comma, and the pair is enclosed in between parentheses. 
			The first constant is the real part, and the second is the imaginary part. 
			A complex constant, COMPLEX*8, uses 8 bytes of storage. 

			Example: 

			(0.0, 0.0) (-123.456E+30, 987.654E-29)

	5 C) Floating Type Constant: 


		    A floating type constant has an integer part, a decimal point, a fractional part, and an exponent part.
		    Can be represent floating constant either in decimal form or exponential form.
		    While representing using the decimal form, you must include the decimal point, the exponent, or both.
		    And while representing using the exponential form, must include the integer part, the fractional part, or both.

			Following are the examples of Floating type constant: 

			3.14159       /* Legal */
			314159E-5L    /* Legal */
			510E          /* Illegal: incomplete exponent */
			210f          /* Illegal: no decimal or exponent */
			.e55          /* Illegal: missing integer or fraction */

6) String Literals 

    Go supports two types of string literals i.e., the ” ” (double-quote style) and the ‘ ‘ (back-quote).
    Strings can be concatenated with + and += operators.
    A string contains characters that are similar to character literals: plain characters, escape sequences, and universal characters.And this is of untyped.
    The zero values of string types are blank strings, which can be represented with ” ” or ” in literal.
    String types are all comparable by using operators like ==, != and (for comparing of same types)

	Syntax 

	type _string struct {
	    elements *byte // underlying bytes
	    len      int   // number of bytes
	}

	Example to show string literals: 

	"hello, geeksforgeeks" 

	"hello, \ 

	geeksforgeeks" 

	"hello, " "geeks" "forgeeks" 

	Here, all of the above three statements are similar, i.e., they don’t have any particular type. 


	CODE:

	
package main
 
import "fmt"
 
func main(){
    const A = "GFG"
    var B = "GeeksforGeeks"
     
    // Concat strings.
    var helloWorld = A+ " " + B
    helloWorld += "!"
    fmt.Println(helloWorld)
     
    // Compare strings.
    fmt.Println(A == "GFG")  
    fmt.Println(B < A)
}

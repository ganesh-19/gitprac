1) Data types specify the type of data that a valid Go variable can hold. 
	In Go language, the type is divided into four categories which are as follows:

	    Basic type: Numbers, strings, and booleans come under this category.
	    Aggregate type: Array and structs come under this category.
	    Reference type: Pointers, slices, maps, functions, and channels come under this category.
	    Interface type

	Here, we will discuss Basic Data Types in the Go language. 
	The Basic Data Types are further categorized into three subcategories which are: 

	    Numbers
	    Booleans
	    Strings


2) Integers: 

		In Go language, both signed and unsigned integers are available in four different sizes as 
		shown in the below table. 
		The signed int is represented by int and the unsigned integer is represented by uint.


		Description


			int8 		8-bit signed integer
			int16 	16-bit signed integer
			int32 	32-bit signed integer
			int64 	64-bit signed integer
			uint8 	8-bit unsigned integer
			uint16 	16-bit unsigned integer
			uint32 	32-bit unsigned integer
			uint64 	64-bit unsigned integer
			int 	Both in and uint contain same size, either 32 or 64 bit.
			uint 	Both in and uint contain same size, either 32 or 64 bit.
			rune 	It is a synonym of int32 and also represent Unicode code points.
			byte 	It is a synonym of int8.
			uintptr It is an unsigned integer type. 
					Its width is not defined, but its can hold all the bits of a pointer value.



3) Floating-Point Numbers: 

	In Go language, floating-point numbers are divided into two categories as shown in the below table:

	Data Type 	

		Description
		float32 	32-bit IEEE 754 floating-point number
		float64 	64-bit IEEE 754 floating-point number


	CODE:

		func main(){
		var x float32 = 20.3
		y:=20.2	
		c := x -y 
		fmt.Printf("c is %f", c)
		fmt.Printf(" type of c is %T", c)
	}

		ERROR:

			when y:= 20.2 is used it is by default float64

			so c := x - y throws error 


4) Complex Numbers: 
	The complex numbers are divided into two parts are shown in the below table. 
	float32 and float64 are also part of these complex numbers. 
	The in-built function creates a complex number from its imaginary and real part and in-built imaginary and 
	real function extract those parts.

		complex64 	Complex numbers which contain float32 as a real and imaginary component.
		complex128 	Complex numbers which contain float64 as a real and imaginary component.


	CODE:

		package main

		import "fmt"

		func main(){
			var a complex128 = complex(2,3)
			var b complex64 = complex(3,4)
			
			fmt.Printf("type of a is %T and "+"type of b is %T", a,b)

			// complex number can be displayed using %f
			fmt.Printf("a is %f", a)
		}



5) Booleans

	The boolean data type represents only one bit of information either true or false. 
	The values of type boolean are not converted implicitly or explicitly to any other type. 


6) Strings

	The string data type represents a sequence of Unicode code points. 
	Or in other words, we can say a string is a sequence of immutable bytes, means once a string is created you cannot change that string. 
	A string may contain arbitrary data, including bytes with zero value in the human-readable form.
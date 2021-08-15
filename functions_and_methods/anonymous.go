1) ANONYMOUS FUNCTION:

	Go language provides a special feature known as an anonymous function. 
	An anonymous function is a function which doesn’t contain any name. 
	It is useful when you want to create an inline function. 
	In Go language, an anonymous function can form a closure. 
	An anonymous function is also known as function literal.


2) Syntax:

	func(parameter_list)(return_type){
	// code..

	// Use return statement if return_type are given
	// if return_type is not given, then do not 
	// use return statement
	return
	}()

	HERE:

		a) () --- calls the function 

	CODE: 

		func main() {
			func (){
			fmt.Println("Hello")
		}()
		}

	but 
		func main() {
			()
		}

		func (){
			fmt.Println("Hello")
		}

	does not work



3) you are allowed to assign an anonymous function to a variable. 
	When you assign a function to a variable, then the 
	type of the variable is of function type and you can call that variable like a function call 

	CODE:

		func main() {
			hello := func (){
			fmt.Println("Hello")
		}
			hello()
		}

4) You can also pass arguments in the anonymous function.

	CODE:

		func main() {
			hello := func (name string){
			fmt.Printf("the name is %s", name)
		}
			hello("goku")
		}


5) You can also return an anonymous function from another function. 

	CODE:
		// here name is the parameter for greetings -- of type string and the function returns string
		func greetings() func(name string) string {

			// hello variable has function anonymous and it takes name as parameter and returns string
			hello := func (name string) string{
			return "hello"+name
		}
			// returns the entire function to the variable which takes the call statement
			return hello
		}

		func main(){
			// initialising variable and refering function using variable greetgoku
			// greetsaiyan will take the anonymoius function -- referred by variable hello
			greetsaiyan := greetings()
			fmt.Println(greetsaiyan("goku"))
		}



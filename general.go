1) every time we intialise a variable something new should be there

2) two types of initialisation

3) no identifiers should be left unused -- variables, functions, packages 

4) always open curly braces like -- func main() {   , if (result && result1){

5) always write if...... else like this -- 

	if (result && result1){
		fmt.Println("Positive")
	}else{

6) else if not elif 

7) In printf ("..... %d" + "..... %s", <include variables at last>)
	
	Multiple quoted strings should not be used in first block -- use + instead

		fmt.Printf("index is %d" + "character is %c \n", index,char)
		fmt.Printf("index is %d", index,  "character is %c \n", char)

8) initialisation of array of strings --- var x = []string{"abc", "def", "fgh"}  (or) x := []string{"abc", "def", "fgh"}

9) MAP - dictionary in golang -- map[key_type] value_type {
									key1 : value1 , ......
									}
	mmap := map[int]string{
        "Geeks":22,
        "GFG":22,
        "GeeksforGeeks":22, ----------- gives error 

   mmap := map[int]string{
        22:"Geeks",
        33:"GFG",
        44:"GeeksforGeeks", ----------- correct 


10) making array ---- var x = []string{"abc", "def", "fgh"} --- var y = []int{1,2,3,4,5}


11) func function_name(Parameter-list)(Return_type){
		    // function body.....
		    return ...
		}

		> no need to mention return type if return statement not used -- else error





LIST OF ERRORS:

IN IF ELSE :
	
	- usually when keyword is wrongly mentioned
	- or braces are wrong

	./prog.go:13:4: syntax error: unexpected elif at end of statement
	./prog.go:15:4: syntax error: unexpected else after top level declaration

IN RETURN DURING CALL BY REFERENCE:

	./prog.go:19:2: cannot use a (type *int) as type int in return argument

	-- while swapping values while returning values which are in args we get this error 
	-- when we return the locally declared variable we don't get this
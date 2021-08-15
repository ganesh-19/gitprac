1) SWITCH:

	A switch statement is a multiway branch statement. 
	It provides an efficient way to transfer the execution to different parts of a code based on the value(also called case) of the expression. 
	Go language supports two types of switch statements:

	    A) Expression Switch
	    B) Type Switch

2) EXPRESSION SWITCH:

	Expression switch is similar to switch statement in C, C++, Java language. 
    It provides an easy way to dispatch execution to different parts of code based on the value of the expression. 

    SYNTAX:

    	Syntax:

		switch optstatement; optexpression{
		case expression1: Statement..
		case expression2: Statement..
		...
		default: Statement..
		}

	Important Points:

	    Both optstatement and optexpression in the expression switch are optional statements.
	    If both optstatementand optexpression are present, then a semi-colon(;) is required in between them.
	    If the switch does not contain any expression, then the compiler assume that the expression is true.
	    The optional statement, i.e, optstatement contains simple statements like variable declarations, increment or assignment statements, or function calls, etc.
	    If a variable present in the optional statement, then the scope of the variable is limited to that switch statement.
	    In switch statement, the case and default statement does not contain any break statement. But you are allowed to use break and fallthrough statement if your program required.
	    The default statement is optional in switch statement.
	    If a case can contain multiple values and these values are separated by comma(,).
	    If a case does not contain any expression, then the compiler assume that te expression is true.

	CODE: 

		switch num:=3; num-1{
			case 1:
			fmt.Println("One")
			
			case 2:
			fmt.Println("Two")
			
			case 3:
			fmt.Println("Three")
			
			case 4:
			fmt.Println("Four")
		}


		OTHER WAY:

		  var num int = 3
		  switch num-1{
			case 1:
			fmt.Println("One")
			
			case 2:
			fmt.Println("Two")
			
			case 3:
			fmt.Println("Three")
			
			case 4:
			fmt.Println("Four")
		}

		USING STRING:

		var num string = "Five"
		  switch num{
			case "One":
			fmt.Println("One")
			
			case "Two":
			fmt.Println("Two")
			
			case "Three", "Four", "Five":
			fmt.Println("Five")
		}

		WITHOUT EXPRESSION:

		var num int = 2
		  switch {
			case num ==1:
			fmt.Println("One")
			
			case num ==2:
			fmt.Println("Two")
			
		}


3) TYPE SWITCH:

	Type switch is used when you want to compare types. 
	In this switch, the case contains the type which is going to compare with the type present in the switch expression.

	SYNTAX:

		switch optstatement; typeswitchexpression{
		case typelist 1: Statement..
		case typelist 2: Statement..
		...
		default: Statement..
		}

	IMPORTANT POINTS:

	    The optional statement, i.e., optstatement is similar as in the switch expression.
	    If a case can contain multiple values and these values are separated by comma(,).
	    In type switch statement, the case and default statement do not contain any break statement. But you are allowed to use break and fallthrough statement if your program required.
	    The default statement is optional in type switch statement.
	    The typeswitchexpression is an expression whose result is a type.
	    If an expression is assigned in typeswitchexpression using := operator, then the type of that variable depends upon the type present in case clause. If the case clause contains two or more types, then the type of the variable is the type in which it is created in typeswitchexpression.

	CODE: 

		var value interface{}
		  switch value.(type){
			case int:
			fmt.Println("One")
			
			case bool:
			fmt.Println("Two")
			
			default:
			fmt.Println("default")
			
		}

		ERROR: var num = 3 is not working 
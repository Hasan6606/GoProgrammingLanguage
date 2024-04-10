package main

import (
<<<<<<< HEAD
	"firstmodule/recursive"
	"fmt"
	//"firstmodule/mathexp"
=======
	//"firstmodule/mathexp" //Define structure of module which is calle "mathexp" 
	"fmt"     //Define fmt module because of fact that printed or taking variable by user.
	"firstmodule/recursive" //Define structure of module which is calle "recursive" 
>>>>>>> 28a32a1e407036adf198521d420eee72c25cfc0c
)

//I am Hasan who is the 4th grade students at Gazi University.
//I have referanced coursed by MR. Demirog at ITA Academy.
//I hope it will be helpful effect it.
//If you have any question. you can send a mail or contacted communication with the Linkedln.


func main() {

	//First App like fundamentals of learning programming language
	// if you print any material such as variable,text . First of all, you can import line 3.
	//After that fmt.Println(variable,"printing text variable")
	//fmt.Println("You can write text or your variable up here :)")

	//Example-1 Finding Root
	//if I have a equation which is called "(a*x^2)+(b*x)+c" and wanted root(s) this equation.
	//First of all, parameters of a,b,c attend to discriminant equation disc=(b*b)-(4*a*c)

	//if disc is biger than zero. This equation have different two roots.
	//if disc is equal to zero. This equation have same two roots.
	//if disc is less than zero that mean . This equation have different two roots.

	//I have taken variables which are a,b,c by users.
	//After that,I have attend these on the FirstExample() function.
	/**
	var firstVariable int16
	var secondVariable int16
	var thirdVariable int16
	fmt.Println("a*x^2+b*x+c please enter variables which are a,b,c as ordered:")
	fmt.Scanln(&firstVariable)
	fmt.Println("Enter b value:")
	fmt.Scanln(&secondVariable)
	fmt.Println("Enter c value:")
	fmt.Scanln(&thirdVariable)

	mathexp.FirstExample(firstVariable, secondVariable, thirdVariable)
	**/
	/**
	//Exampe-2:Calculating Factorial Number
	// Factorial number formula  4!=4*3*2*1 when if you enter 0 . The result is equal to 1.
	//First of all, A number is taking by user.
	//After that, Factorial number was calculated by Factorial number as Recursive Function Method.
	//Define variable as taking a number by users.
	var facNumber int32
	//Printed a message on the display.
	fmt.Println("Hello,please enter number for calculating factorial variable:")
	//Defined number on the facNumber variable.
	fmt.Scanln(&facNumber)
	//Defined the result in the "result" variable
	result := recursive.Factorial(int32(facNumber))
	//Printed the Result on the display
	fmt.Println(result)
 	**/

	var lenOfArray int
	var paramArray []int 
	fmt.Println("Enter the len of array:")
	fmt.Scanln(&lenOfArray)
	for i := 1; i <lenOfArray; i++ {
		fmt.Println("Enter value of ",i," elements:")
		fmt.Scanln(&paramArray[i-1])
	}
	fmt.Printlm("Enter researched number on array:")
	var resNum int
	fmt.Scanln(&resNum)
	result:=loop.FindBruteForceAlgorithm(paramArray[], resNum)
	fmt.Println(result)
}

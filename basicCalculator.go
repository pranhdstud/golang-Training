package main
import (
	"fmt"
	
)

func main(){
	var numOne, numTwo int
	fmt.Printf("Enter number One : ")
	fmt.Scanf("%d\n",&numOne)
	fmt.Printf("Enter number Two : ")
	fmt.Scanf("%d\n",&numTwo)
	result := calculator(numOne,numTwo)
	fmt.Printf("The result of operation is %d",result)
}

func calculator(numOne int, numTwo int) int{
	var answer int
	var operation rune
	fmt.Printf("1.Press + for addition\n2.Press - for substraction\n3.Press * for multiplication\n4.Press / for division\n")
	fmt.Scanf("%c",&operation)
	fmt.Printf("Operation is %c\n",operation)	
    if numOne >= numTwo {
    switch operation{
    case '+' :
    	answer = numOne + numTwo
    	break
    case '-' :
    	answer = numOne - numTwo
    	break
    case '*' :
    	answer = numOne * numTwo
        break
    case '/' :
    	answer = numOne / numTwo
    	break	
    }
    
    }
    if numTwo >= numOne {
    	switch operation{
    case '+' :
    	answer = numTwo + numOne
    	break
    case '-' :
    	answer = numTwo - numOne
    	break
    case '*' :
    	answer = numTwo * numOne
        break
    case '/' :
    	answer = numTwo / numOne
    	break	
    }
    }    
return answer;


}
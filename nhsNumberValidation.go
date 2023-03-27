package main

import (
	"fmt"
	"strconv"
)

func main(){
	fmt.Println("Enter your 10 digit NHS number: ")
    var nhsNumber int

    fmt.Scanln(&nhsNumber) 

	var nhsNumLen= len(strconv.Itoa(nhsNumber));
	if(nhsNumLen!=10){
		fmt.Println("Error: NHS number must be a 10 digit number")
		return
	}

	var valueCheckerSum int
	checkDigit:=nhsNumber%10 //gets last digit of the number as check digit
	nhsNumber=nhsNumber/10 //removes the last digit which is the check digit

	for i := 2; i <=nhsNumLen; i++ {
		digit:=nhsNumber%10 //gets last digit of the number
		valueCheckerSum=valueCheckerSum+(digit*i) //sum of multiplying first nine digits by the weighing factor(i)
		nhsNumber=nhsNumber/10 //removes the last digit which is added into valuecheckerSum
	}
	
	reminderChecker:=valueCheckerSum%11

	if(reminderChecker==0){ // if reminder is 0, NHS num is valid
		fmt.Println("Given NHS Number is valid")
	}else{
		valueChecker:=11-reminderChecker
	 	if(valueChecker==checkDigit){ //conditions to check digit validity
			fmt.Println("Given NHS Number is valid")
		}else{
			fmt.Println("Error: Given NHS Number is valid is invalid")
		}
	}	
}
package main

import (
	"fmt"
	"os"
	"strconv"
)
var rules string = "One operation: --add,--sub,--mul--divi with two numbers //Eg: --add 12 13"
func main(){
	args := os.Args
	if len(args)!=4{
		fmt.Println(rules)
		// fmt.Print(len(args))
		return
	}
	operation := args[1]
	firstnumber := args[2]
	secondnumber := args[3]

	if operation != "--add" && operation != "--sub" &&
		operation != "--mul" && operation != "--div" {
		fmt.Println(rules)
		return
	}
	num1,err1  := strconv.Atoi(firstnumber)
	num2,err2  := strconv.Atoi(secondnumber)
	if err1 !=nil || err2!=nil{
		fmt.Println("Error: please enter valid numbers")
		return
	}


	switch operation{
	case "--add":
		fmt.Println("****************")
		fmt.Println("")
		fmt.Printf("Your Calculated Value is %v \n",num1+num2)
		fmt.Println("")
		fmt.Println("****************")
	case "--sub":
		fmt.Println("****************")
		fmt.Println("")
		fmt.Printf("Your Calculated Value is %v \n",num1-num2)
		fmt.Println("")
		fmt.Println("****************")
	case "--mul":
		fmt.Println("****************")
		fmt.Println("")
		fmt.Printf("Your Calculated Value is %v \n",num1*num2)
		fmt.Println("")
		fmt.Println("****************")
	case "--div":
		if num2 == 0{
			fmt.Println("Division by zero error")
			
			return
		}
		fmt.Println("****************")
		fmt.Println("")
		fmt.Printf("Your Calculated Value is %v \n",num1/num2)
		fmt.Println("")
		fmt.Println("****************")

	}
	fmt.Println("Program done")
	
	
	
}

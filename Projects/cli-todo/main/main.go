package main

import (
	"bufio"
	"cli-todo/todos"
	"fmt"
	"os"
	"strings"
)
func main(){

fmt.Println("Welcome to Todo list Project")
var usertodo string

//reading a full line

reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your todo: ")
	usertodo, _ = reader.ReadString('\n')

	usertodo = strings.TrimSpace(usertodo)
	todos.AddTodo(usertodo)
	fmt.Println("")
	todos.MarkDown(1)
	todos.Displaytodo()
	



}

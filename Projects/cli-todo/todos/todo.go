package todos

import "fmt"


type Todo struct{
	ID int
	Title string
	Done bool
}
var todo = []Todo{}
func AddTodo(title string){
	
	addtodo := Todo{
		ID:len(todo)+1,
		Title: title,
		Done:false,
	}
	todo = append(todo,addtodo )
}
func  Displaytodo()  {
	for _,todo := range todo{
		fmt.Println("Title: ",todo.Title)
		fmt.Println("Status: ",todo.Done)
	}
	
}
func MarkDown(id int){
	for i :=range todo{
		if todo[i].ID == id{
			todo[i].Done = true
			return
		}

	}
}
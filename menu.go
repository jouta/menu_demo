package main
 
import "fmt"
 
func main() {
	for {
		var input string
		fmt.Print("Please choose a page(print 0 or 1) :")
      	fmt.Scan(&input)
      	if input == "quit" {
         	break
      	}
      	switch input {
      		case "0":
          		fmt.Println("the 1st page")
      		case "1":
          		fmt.Println("the 2nd page")
      		default:
          		fmt.Println("err")
      	}
  	}
}
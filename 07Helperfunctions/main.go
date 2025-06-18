package main

import "fmt"

func greet(name string) string {
	return "Hello," + name + "!"
}

func main() {
	// Call the greet function and save its result
	message := greet("jordan")
	// Print the returned message
	fmt.Println(message)
}


package main

func greet (name string) string{
	return "Call," + name + "!"
}

func main() {
	message := greet("Mike")
	fmt.Println(message)
}
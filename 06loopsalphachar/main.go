package main
import "fmt"

//func main () {
//char := 'A'
//fmt.Printf(Character: %c \n, char)
//fmt.Printf("ASCII/Unicode : %d and type: %T\n, char, char")
//}

func main() {
	word := "Hello Coders"
	for i, r := range word {
		fmt.Printf("character %d: %c = %d\n", i, r, r)
	}
	fmt.Println(word)
}
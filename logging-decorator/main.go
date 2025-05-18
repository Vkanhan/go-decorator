package main 

import(
	"fmt"
)

func WithLogging(f func()) func() {
	return func() {
		fmt.Println("Before function call")
		f()
		fmt.Println("After function call")
	}
}

func sayHello() {
	fmt.Println("Hello")
}

func main() {
	loggedSayHello := WithLogging(sayHello)
	loggedSayHello()
}
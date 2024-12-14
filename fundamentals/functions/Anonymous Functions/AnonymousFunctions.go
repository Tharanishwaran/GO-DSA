package main

import "fmt"

func main() {
    // Anonymous function assigned to a variable
    greet := func(name string) string {
        return "Hello, " + name
    }
    
    // Immediate invocation
    message := greet("Alice") // message = "Hello, Alice"
	fmt.Println(message)
    
    // Closure capturing external variable
    multiplier := func(factor int) func(int) int {
        return func(x int) int {
            return x * factor
        }
    }
    
    double := multiplier(2)						
    result := double(5) // result = 10

	fmt.Println(result)
}

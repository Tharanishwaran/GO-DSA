package main

import "fmt"

// import (
//     "fmt"
//     "log"
// )

// // Function that might cause a panic
// func divideNumbers(a, b int) int {
//     if b == 0 {
//         // Explicitly cause a panic
//         panic("cannot divide by zero")
//     }
//     return a / b
// }

// // Safe division function using recover
// func safeDivide(a, b int) (result int, err error) {
//     // Deferred function to handle potential panic
//     defer func() {
//         // Recover must be called inside a deferred function
//         if r := recover(); r != nil {
//             // Log the recovered panic
//             log.Printf("Recovered from panic: %v", r)
            
//             // Convert panic to error
//             err = fmt.Errorf("panic occurred: %v", r)
//             result = 0
//         }
//     }()

//     // Potentially panicking operation
//     return divideNumbers(a, b), nil
// }

// func main() {
//     // Scenario 1: Normal division
//     result1, err1 := safeDivide(10, 2)
//     fmt.Printf("Normal division - Result: %d, Error: %v\n", result1, err1)

//     // Scenario 2: Division by zero
//     result2, err2 := safeDivide(10, 0)
//     fmt.Printf("Division by zero - Result: %d, Error: %v\n", result2, err2)
// }

func divideNumbers(a, b int) int {
    if b == 0 {
        panic("division by zero")
    }

	fmt.Println(a/b)
    return a / b
}

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered:", r)
        }
    }()
    
    divideNumbers(10, 0)
}
package main

import "fmt"

func main()  { 
    num := 1
    // p := &num
var p int = &num

    fmt.Println(&num)    
    fmt.Println(*p)
  
}

func revalue(n *int) int {
    *n = 5
    return *n
}

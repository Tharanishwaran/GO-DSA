package main

import (

	"fmt"

)


func P(v ...interface{}) {
    fmt.Println(v...)
}

func PF(format string, v ...interface{}) {
    fmt.Printf(format, v...)
	fmt.Println()
	
}

func main(){

  
  symbols := map [string]uint{

    "A":4,
    "B":7,
    "C":12,
    "D":20,

  }

  symbolArr := generateSymbolsArray(symbols) 

  P(symbolArr)

  var balance uint = 21


  GetName()


  for balance > 0 {

  bet :=  GetBet(balance)

  if bet == 0{
	break
  }  

  balance -= bet
}

 PF("you left with, $%d.", balance)

}




func generateSymbolsArray(symbols map[string]uint) []string {
  
 symbolArr := []string{}

  for symbol, count := range symbols{

    for i := uint(0); i < count; i++ {

      symbolArr = append(symbolArr, symbol)
    }

  }

  return symbolArr

}

// func getSpin(reel []string) [][]string {

//   result := [][]

//   {{},{},{}}
  
// }
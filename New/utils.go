package main

import "fmt"

func GetName() string{

	name  := "" 
 
   fmt.Print("Enter Your name : ")
 
   _, err := fmt.Scanln(&name)
 
   if err != nil{
		
	  return ""
 
   }
 
   fmt.Printf("Welcomr %s, lets play! \n",name)
 
 
   return name
 
 }
 
 func GetBet(balance uint) uint {
 
   var bet uint
 
   for true {
	
	 fmt.Printf("Enter your bet or 0 to quit (balance =$5%d): ",balance)
	 fmt.Scan(&bet)
 
	 if bet > balance {
 
		 P("Bet cannot be larger than balance.")    
	 } else {
		 break
	 }
 
   }
  return bet
 
 }
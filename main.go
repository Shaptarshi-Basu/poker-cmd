package main

import (
	"fmt"
	"log"
	"os"
	"poker/operations"
)

func main() {
	args := os.Args[1:] // Get all command-line arguments except for the program name

	if len(args) < 2 {
		fmt.Println("Please provide two string arguments as the two hands of the poker.")
		return
	}
	fh := args[0]
	sh := args[1]
	player1 := operations.Player{Hand_Number: "Hand 1", Hand: fh}
	player2 := operations.Player{Hand_Number: "Hand 2", Hand: sh}
	result, err := operations.GetWinner(player1, player2)
	if err != nil {
		log.Fatalln(err.Error())
	}
	println(result)

}

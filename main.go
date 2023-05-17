package main

import (
	"fmt"
	"log"
	"poker/operations"
)

func main() {
	var fh, sh string

	fmt.Println("Enter the first Hand:")
	fmt.Scan(&fh)

	fmt.Println("Enter the second Hand:")
	fmt.Scan(&sh)

	player1 := operations.Player{Hand_Number: "Hand 1", Hand: fh}
	player2 := operations.Player{Hand_Number: "Hand 2", Hand: sh}
	result, err := operations.GetWinner(player1, player2)
	if err != nil {
		log.Fatalln(err.Error())
	}
	println(result)

}

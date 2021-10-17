package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ticTacToe/src/tictactoeGame"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Reader intialized start writing.....")
	var input []string
	for true {
		text, _ := reader.ReadString('\n')
		text = strings.Trim(text, "\n")
		if text == "exit" {
			break
		}
		input = append(input, text)

	}
	game := &tictactoeGame.TicTacToeGame{}
	if err := game.Intialize(3, input[:2], "square", ""); err != nil {
		fmt.Printf("failure in intializing board %+v \n", err)
		return
	}
	game.Play(input[2:])

}

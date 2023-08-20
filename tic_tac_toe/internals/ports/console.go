package ports

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/tic_tac_toe/internals/domains"

	"github.com/tic_tac_toe/internals/domains/enums"
)

type Console struct {
	service *domains.GameService
	scanner *bufio.Scanner
	players int
}

func NewConsole(service *domains.GameService, scanner *bufio.Scanner, players int) *Console {
	return &Console{
		scanner: scanner,
		players: players,
		service: service,
	}
}

func (console *Console) StartGame() {
	players := players(console.scanner, console.players)
	console.service.AddPlayers(players, 0)

	moves := moves(console.scanner)
	fmt.Println(" ***************** Starting game play ***************")
	for _, curMove := range moves {
		gameResp := console.service.PlayMove(curMove)
		if gameResp.Continue == true {
			if gameResp.Error == nil {
				fmt.Println(fmt.Sprintf(console.service.PrintLayout()))
			} else {
				fmt.Println(gameResp.Msg)
			}
			continue
		}
		fmt.Println(fmt.Sprintf(console.service.PrintLayout()))
		fmt.Println(gameResp.Msg)
		break

	}
}

func players(scanner *bufio.Scanner, player int) []*domains.Player {
	players := make([]*domains.Player, 0)
	for index := 0; index < player; index++ {
		fmt.Println("Enter player name for :- ", enums.Layout(enums.TokenByIndex(index+1)))
		scanner.Scan()
		players = append(players, domains.NewPlayer(scanner.Text(), enums.TokenByIndex(index+1)))

	}
	return players
}

func moves(scanner *bufio.Scanner) []domains.Move {
	var (
		moves      = make([]domains.Move, 0)
		input      string
		inputMoves []string
	)
	for true {
		scanner.Scan()
		input = scanner.Text()
		if inputMoves = strings.Split(input, " "); len(inputMoves) != 2 {
			break
		}

		inputMoveOne, _ := strconv.Atoi(inputMoves[0])
		inputMoveTwo, _ := strconv.Atoi(inputMoves[1])
		moves = append(moves, *domains.NewMove(inputMoveOne, inputMoveTwo))

	}
	return moves
}

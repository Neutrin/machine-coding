package tictactoeGame

import "testing"

func Test_Initalize(t *testing.T) {
	type args struct {
		size         int
		playerList   []string
		boardType    string
		strategyType string
	}
	tests := []struct {
		name    string
		wantErr bool
		args    args
	}{
		{
			name:    "Positive with size 3",
			wantErr: false,
			args: args{
				size:         3,
				playerList:   []string{"x nitin", "o ujjwal"},
				boardType:    "square",
				strategyType: "hashmap",
			},
		},
	}
	for _, curTest := range tests {
		t.Run(curTest.name, func(t *testing.T) {
			game := &TicTacToeGame{}
			err := game.Intialize(curTest.args.size, curTest.args.playerList, curTest.args.boardType, curTest.args.strategyType)
			if (curTest.wantErr == true && err == nil) || (err != nil && !curTest.wantErr) {
				t.Errorf("Wanted error = %v, got = %v", curTest.wantErr, err)
			}
		})
	}
}

package winningStrategySquare

import (
	"reflect"
	"testing"
)

func Test_Intialize(t *testing.T) {

	type args struct {
		size      int
		userCount int
	}
	tests := []struct {
		name     string
		args     args
		expected *WinningStrategySquareHashMapImpl
		wantErr  bool
	}{
		{
			name: "intialise with 3 value",
			args: args{
				size:      3,
				userCount: 2,
			},
			expected: &WinningStrategySquareHashMapImpl{
				size:             3,
				rowUserMap:       make(map[int]map[rune]int),
				colUserMap:       make(map[int]map[rune]int),
				diagnalCount:     make(map[rune]int),
				antiDiagnalCount: make(map[rune]int),
			},
			wantErr: false,
		},
	}
	for _, curTests := range tests {
		t.Run(curTests.name, func(t *testing.T) {
			strategy := &WinningStrategySquareHashMapImpl{}

			err := strategy.Intialize(curTests.args.size, curTests.args.userCount)
			if (err == nil && curTests.wantErr == true) || (err != nil && curTests.wantErr == false) {
				t.Errorf("Wanted = %v and got = %v", curTests.wantErr, err)
			}
			if err == nil && !reflect.DeepEqual(
				*curTests.expected, *strategy) {
				t.Errorf("Wanted = %+v , got = %+v", *curTests.expected, *strategy)
			}
		})
	}
}

func Test_MakeMove(t *testing.T) {
	type args struct {
		peice    rune
		row, col int
	}
	tests := []struct {
		name     string
		args     args
		object   *WinningStrategySquareHashMapImpl
		expected bool
	}{
		{
			name: "Row winning",
			object: &WinningStrategySquareHashMapImpl{
				size: 3,
				rowUserMap: map[int]map[rune]int{
					1: map[rune]int{
						rune('x'): 2,
						rune('o'): 2,
					},
					2: make(map[rune]int),
				},
				colUserMap:       make(map[int]map[rune]int),
				diagnalCount:     make(map[rune]int),
				antiDiagnalCount: make(map[rune]int),
			},
			expected: true,
			args: args{
				peice: rune('x'),
				row:   1, col: 1,
			},
		},
		{
			name: "col winning",
			object: &WinningStrategySquareHashMapImpl{
				size: 3,
				colUserMap: map[int]map[rune]int{
					1: map[rune]int{
						rune('x'): 2,
						rune('o'): 2,
					},
					2: make(map[rune]int),
				},
				rowUserMap:       make(map[int]map[rune]int),
				diagnalCount:     make(map[rune]int),
				antiDiagnalCount: make(map[rune]int),
			},
			expected: true,
			args: args{
				peice: rune('x'),
				row:   1, col: 1,
			},
		},
		{
			name: "Make diagnal win",
			object: &WinningStrategySquareHashMapImpl{
				size: 3,
				rowUserMap: map[int]map[rune]int{
					1: map[rune]int{
						rune('x'): 1,
					},
					2: map[rune]int{
						rune('x'): 1,
						rune('o'): 1,
					},
				},
				colUserMap: make(map[int]map[rune]int),
				diagnalCount: map[rune]int{
					rune('x'): 2,
					rune('o'): 1,
				},
				antiDiagnalCount: map[rune]int{
					rune('x'): 1,
					rune('o'): 1,
				},
			},
			args: args{
				peice: rune('x'),
				row:   3, col: 3,
			},
			expected: true,
		},
		{
			name: "Anti dignal win",
			object: &WinningStrategySquareHashMapImpl{
				size:         3,
				rowUserMap:   make(map[int]map[rune]int),
				colUserMap:   make(map[int]map[rune]int),
				diagnalCount: make(map[rune]int),
				antiDiagnalCount: map[rune]int{
					rune('x'): 2,
				},
			},
			args: args{
				peice: rune('x'),
				row:   1, col: 3,
			},
			expected: true,
		},
	}
	for _, curTest := range tests {
		t.Run(curTest.name, func(t *testing.T) {
			expected := curTest.object.MakeMove(curTest.args.peice, curTest.args.row, curTest.args.col)
			if expected != curTest.expected {
				t.Errorf("Respone expected was = %v and got = %v", curTest.expected, expected)
			}
		})
	}
}

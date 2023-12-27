package domain

import (
	"fmt"
	"testing"
)

func TestPosition(t *testing.T) {
	type args struct {
		input string
	}
	pos, err := NewPosition("e2")
	if err != nil {
		t.Errorf(" failed got error ")
	}
	if pos.RowNo != 5 && pos.ColNo != 2 {
		t.Errorf(" got = %+v", *pos)
	}

}

func TestPrint(t *testing.T) {
	b := NewDefBoard(8)
	b.Print()
}

func TestIsValidMove(t *testing.T) {
	b := NewDefBoard(8)
	k := NewKnight(White)
	intialPos, _ := NewPosition("g1")
	finalPos, _ := NewPosition("f3")
	fmt.Println(" intial ", *intialPos)
	fmt.Println(" final ", *finalPos)
	err := k.IsValidMove(intialPos, finalPos, b)
	if err != nil {
		t.Errorf(" came error ")
	}
	finalPos, _ = NewPosition("e2")
	err = k.IsValidMove(intialPos, finalPos, b)
	if err == nil {
		t.Errorf(" error should not be nil ")
	}

}

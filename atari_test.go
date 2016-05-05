package main

import (
	"testing"
	"reflect"
	"fmt"
)

func TestBuildingEmptyBoard(t *testing.T) {
	cases := []struct {
		inSize int
		want [][]int
	}{
		{3, [][]int{[]int{0, 0, 0}, []int{0, 0, 0}, []int{0, 0, 0}}},
	}
	for _, c := range cases {
		got := buildEmptyBoard(c.inSize)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\nWant: %v\nGot:  %v", c.want, got)
		} else {
			//fmt.Println("\nTest buildEmptyBoard(boardSize) function:")
			//fmt.Printf("size: %v\nresult: %v\n", c.inSize, got)
		}
	}
}

func TestPuttingStoneOnBoard(t *testing.T) {
	cases := []struct {
		stone Stone
		inBoard [][]int
		want [][]int
	}{
		{Stone{1, 1, 1}, [][]int{[]int{0, 0, 0}, []int{0, 0, 0}, []int{0, 0, 0}}, [][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{0, 0, 0}}},
		{Stone{2, 1, 2}, [][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{0, 0, 0}}, [][]int{[]int{0, 0, 0}, []int{0, 1, 2}, []int{0, 0, 0}}},
	}
	for _, c := range cases {
		got := c.stone.putOnBoard(c.inBoard)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\nWant: %v\nGot:  %v\n", c.want, got)
		} else {
			//fmt.Println("\nTest putStoneOnBoard(emptyBoard) function:")
			//fmt.Printf("stone: %v\nresult: %v\n", c.stone, c.want)
		}
	}
}

func TestPuttingStoneOnOtherStone(t *testing.T) {
	cases := []struct {
		stone Stone
		inBoard [][]int
		want bool
	}{
		{Stone{2, 1, 1}, [][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{0, 0, 0}}, false},
		{Stone{2, 1, 2}, [][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{0, 0, 0}}, true},
		{Stone{1, 1, 1}, [][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{0, 0, 0}}, false},
	}
	for _, c := range cases {
		got := c.stone.assertMovePossible(c.inBoard)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\nWant: %v\nGot:  %v\n", c.want, got)
		} else {
			//fmt.Println("\nTest assertMovePossible(emptyBoard) function:")
			//fmt.Printf("board: %v\nstone: %v\nresult: %v\n", c.inBoard, c.stone, c.want)
		}
	}
}

func TestPlayingGame(t *testing.T) {
	cases := []struct {
		stones []Stone
		inBoard [][]int
		want [][]int
	}{
		{[]Stone{{1, 1, 1}, {2, 1, 2}}, [][]int{[]int{0, 0, 0}, []int{0, 0, 0}, []int{0, 0, 0}}, [][]int{[]int{0, 0, 0}, []int{0, 1, 2}, []int{0, 0, 0}}},
		{[]Stone{{2, 1, 1}}, [][]int{[]int{0, 0, 0}, []int{0, 0, 0}, []int{0, 0, 0}}, [][]int{[]int{0, 0, 0}, []int{0, 2, 0}, []int{0, 0, 0}}},
	}
	for _, c := range cases {
		got := playGame(c.stones, c.inBoard)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\nWant: %v\nGot:  %v\n", c.want, got)
		} else {
			//fmt.Println("\nTest playGame(stones, emptyBoard) function:")
			//fmt.Printf("stones: %v\nresult: %v\n", c.stones, c.want)
		}
	}
}

func TestPlayingGameWithIllegalMove(t *testing.T) {
	cases := []struct {
		stones []Stone
		inBoard [][]int
		want [][]int
	}{
		{[]Stone{{1, 1, 1}, {2, 1, 2}, {1, 1, 2}}, [][]int{[]int{0, 0, 0}, []int{0, 0, 0}, []int{0, 0, 0}}, [][]int{[]int{0, 0, 0}, []int{0, 1, 2}, []int{0, 0, 0}}},
		{[]Stone{{1, 1, 1}, {2, 1, 2}, {1, 1, 1}, {2, 2, 2}}, [][]int{[]int{0, 0, 0}, []int{0, 0, 0}, []int{0, 0, 0}}, [][]int{[]int{0, 0, 0}, []int{0, 1, 2}, []int{0, 0, 0}}},
	}
	for _, c := range cases {
		got := playGame(c.stones, c.inBoard)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\nWant: %v\nGot:  %v\n", c.want, got)
		} else {
			//fmt.Println("\nTest playGameWithIllegalMove(stones, emptyBoard) function:")
			//fmt.Printf("stones: %v\nresult: %v\n", c.stones, c.want)
		}
	}
}

func TestFindNeighboors(t *testing.T) {
	cases := []struct {
		stone Stone
		inBoard [][]int
		want []Stone
	}{
		{Stone{1, 1, 1}, [][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{0, 0, 0}}, []Stone{{0, 1, 0}, {0, 1, 2}, {0, 0, 1} ,{0, 2, 1}}},
		{Stone{1, 0, 0}, [][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{0, 0, 0}}, []Stone{{0, 0, 1}, {0, 1, 0}}},
		{Stone{1, 2, 2}, [][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{0, 0, 0}}, []Stone{{0, 2, 1}, {0, 1, 2}}},
		{Stone{1, 0, 2}, [][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{0, 0, 0}}, []Stone{{0, 0, 1}, {0, 1, 2}}},
		{Stone{1, 2, 0}, [][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{0, 0, 0}}, []Stone{{0, 2, 1}, {0, 1, 0}}},
		{Stone{1, 1, 1}, [][]int{[]int{0, 2, 0}, []int{2, 1, 2}, []int{0, 2, 0}}, []Stone{{2, 1, 0}, {2, 1, 2}, {2, 0, 1}, {2, 2, 1}}},
	}
	for _, c := range cases {
		got := c.stone.getNeighboors(c.inBoard)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\nWant: %v\nGot:  %v\n", c.want, got)
		} else {
			//fmt.Println("\nTest getNeighboors(stone, emptyBoard) function:")
			//fmt.Printf("stone: %v\nresult: %v\n", c.stone, c.want)
		}
	}
}

func TestFindingOpponentForStone(t *testing.T) {
	cases := []struct {
		stone Stone
		inBoard [][]int
		want []Stone
	}{
		{Stone{1, 1, 1}, [][]int{[]int{0, 2, 0}, []int{0, 1, 0}, []int{0, 0, 0}}, []Stone{{2, 0, 1}}},
		{Stone{2, 1, 1}, [][]int{[]int{0, 1, 0}, []int{1, 2, 1}, []int{1, 1, 0}}, []Stone{{1, 1, 0}, {1, 1, 2}, {1, 0, 1}, {1, 2, 1}}},
	}
	for _, c := range cases {
		got := c.stone.getOpponents(c.inBoard)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\nWant: %v\nGot:  %v\n", c.want, got)
		} else {
			//fmt.Println("\nTest getOpponents(stone, emptyBoard) function:")
			//fmt.Printf("stones: %v\nresult: %v\n", c.stone, c.want)
		}
	}
}

func TestFindingFriendsForStone(t *testing.T) {
	cases := []struct {
		stone Stone
		inBoard [][]int
		want []Stone
	}{
		{Stone{1, 1, 1}, [][]int{[]int{0, 2, 0}, []int{1, 1, 1}, []int{0, 0, 0}}, []Stone{{1, 1, 0}, {1, 1, 2}}},
		{Stone{2, 1, 1}, [][]int{[]int{0, 2, 0}, []int{1, 2, 1}, []int{0, 1, 0}}, []Stone{{2, 0, 1}}},
	}
	for _, c := range cases {
		got := c.stone.getFriends(c.inBoard)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\nWant: %v\nGot:  %v\n", c.want, got)
		} else {
			//fmt.Println("\nTest getFriends(stone, emptyBoard) function:")
			//fmt.Printf("stones: %v\nresult: %v\n", c.stone, c.want)
		}
	}
}

func TestStoneInGroup(t *testing.T) {
	cases := []struct {
		stone Stone
		group []Stone
		want bool
	}{
		{Stone{1, 1, 1}, []Stone{{1, 1, 1}, {1, 1, 0}}, true},
		{Stone{1, 1, 1}, []Stone{{1, 1, 0}, {1, 1, 1}}, true},
		{Stone{2, 1, 1}, []Stone{{1, 1, 1}, {2, 0, 1}}, false},
	}
	for _, c := range cases {
		got := isStoneInGroup(c.stone, c.group)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\nWant: %v\nGot:  %v\n", c.want, got)
		} else {
			//			fmt.Println("\nTest isStoneInGroup(stone, emptyBoard) function:")
			//			fmt.Printf("stones: %v\nresult: %v\n", c.stone, c.want)
		}
	}
}

func TestMakingGroupForStone(t *testing.T) {
	cases := []struct {
		stone Stone
		inBoard [][]int
		want []Stone
	}{
		{Stone{1, 1, 1}, [][]int{[]int{0, 2, 0}, []int{1, 1, 0}, []int{0, 0, 0}}, []Stone{{1, 1, 1}, {1, 1, 0}}},
		{Stone{2, 1, 1}, [][]int{[]int{0, 2, 0}, []int{1, 2, 1}, []int{0, 1, 0}}, []Stone{{2, 1, 1}, {2, 0, 1}}},
		{Stone{2, 1, 1}, [][]int{[]int{0, 2, 0}, []int{1, 2, 1}, []int{0, 2, 0}}, []Stone{{2, 1, 1}, {2, 0, 1}, {2, 2, 1}}},
		{Stone{2, 2, 1}, [][]int{[]int{0, 2, 0}, []int{1, 2, 1}, []int{0, 2, 0}}, []Stone{{2, 1, 1}, {2, 0, 1}, {2, 2, 1}}},
		{Stone{2, 2, 1}, [][]int{[]int{2, 2, 2}, []int{2, 2, 2}, []int{0, 2, 0}}, []Stone{{2, 1, 1}, {2, 0, 1}, {2, 2, 1}, {2, 0, 2}, {2, 0, 0}, {2, 1, 2}, {2, 1, 0}}},
	}
	for _, c := range cases {
		got := makeGroupForStone(c.stone, c.inBoard)
		if !assertStonesSlicesEqual(got, c.want) {
			t.Errorf("\nWant: %v\nGot:  %v\n", c.want, got)
		} else {
			//			fmt.Println("\nTest makeGroupForStone(stone, emptyBoard) function:")
			//			fmt.Printf("stones: %v\nresult: %v\n", c.stone, c.want)
		}
	}
}

func TestCountingStoneLiberties(t *testing.T) {
	cases := []struct {
		stone Stone
		board [][]int
		want int
	}{
		{Stone{1, 1, 1}, [][]int{[]int{0, 2, 0}, []int{2, 1, 0}, []int{0, 0, 0}}, 2},
		{Stone{1, 1, 1}, [][]int{[]int{0, 2, 0}, []int{2, 1, 0}, []int{0, 2, 0}}, 1},
		{Stone{1, 1, 1}, [][]int{[]int{0, 2, 0}, []int{2, 1, 0}, []int{0, 1, 0}}, 1},
	}
	for _, c := range cases {
		got := countStoneLiberties(c.stone, c.board)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\nWant: %v\nGot:  %v\n", c.want, got)
		} else {
			//			fmt.Println("\nTest countStoneLiberties(stone, board):")
			//			fmt.Printf("result: %v\n", c.want)
		}
	}
}

func TestCountingGroupLiberties(t *testing.T) {
	cases := []struct {
		group []Stone
		board [][]int
		want int
	}{
		{[]Stone{{1, 1, 1}}, [][]int{[]int{0, 2, 0}, []int{2, 1, 0}, []int{0, 2, 0}}, 1},
		{[]Stone{{1, 1, 1}, {1, 1, 0}}, [][]int{[]int{0, 1, 0}, []int{2, 1, 0}, []int{0, 2, 0}}, 3},
		{[]Stone{{1, 1, 1}, {1, 1, 0}}, [][]int{[]int{0, 1, 0}, []int{0, 1, 0}, []int{0, 0, 0}}, 5},
		{[]Stone{{1, 0, 0}, {1, 0, 1}, {1, 0, 2}, {1, 1, 0}, {1, 1, 2}, {1, 2, 0}, {1, 2, 1}, {1, 2, 2}}, [][]int{[]int{1, 1, 1}, []int{1, 0, 1}, []int{1, 1, 1}}, 1},
	}
	for _, c := range cases {
		got := countGroupLiberties(c.group, c.board)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\nWant: %v\nGot:  %v\n", c.want, got)
		} else {
			//			fmt.Println("\nTest countGroupLiberties(group, board):")
			//			fmt.Printf("result: %v\n", c.want)
		}
	}
}

func TestGroupIsAlive(t *testing.T) {
	cases := []struct {
		group []Stone
		board [][]int
		want bool
	}{
		{[]Stone{{1, 1, 1}}, [][]int{[]int{0, 2, 0}, []int{2, 1, 2}, []int{0, 2, 0}}, false},
		{[]Stone{{1, 1, 1}}, [][]int{[]int{0, 0, 0}, []int{2, 1, 2}, []int{0, 2, 0}}, true},
		{[]Stone{{1, 1, 1}, {1, 0, 1}}, [][]int{[]int{2, 1, 2}, []int{2, 1, 2}, []int{0, 0, 0}}, true},
		{[]Stone{{1, 1, 1}, {1, 0, 1}}, [][]int{[]int{2, 1, 2}, []int{2, 1, 2}, []int{0, 2, 0}}, false},
		{[]Stone{{1, 0, 0}, {1, 0, 1}, {1, 0, 2}, {1, 1, 0}, {1, 1, 2}, {1, 2, 0}, {1, 2, 1}, {1, 2, 2}}, [][]int{[]int{1, 1, 1}, []int{1, 0, 1}, []int{1, 1, 1}}, true},
		{[]Stone{{1, 0, 0}, {1, 0, 1}, {1, 0, 2}, {1, 1, 0}, {1, 1, 2}, {1, 2, 0}, {1, 2, 1}, {1, 2, 2}}, [][]int{[]int{1, 1, 1}, []int{1, 2, 1}, []int{1, 1, 1}}, false},
	}
	for _, c := range cases {
		got := isGroupAlive(c.group, c.board)
		if !reflect.DeepEqual(got, c.want) {
			fmt.Printf("\n%v => %v\n", countGroupLiberties(c.group, c.board), findLibertiesForGroup(c.group, c.board))
			printBoardToConsole(c.board)

			t.Errorf("\nWant: %v\nGot:  %v\n", c.want, got)
		} else {
			//			fmt.Println("\nTest isGroupAlive(group, board):")
			//			fmt.Printf("result: %v\n", c.want)
		}
	}
}

func TestGettingMovesPartOfGameRecordAsArray(t *testing.T) {
	cases := []struct {
		gameRecord string
		want []string
	}{
		{"(;SZ[3]PB[Lee]PW[Alpha]RE[B+1];B[bb];W[ab];B[aa];W[ba])", []string{"B[bb]", "W[ab]", "B[aa]", "W[ba]"}},
	}
	for _, c := range cases {
		got := getMovesPartOfGameRecordAsArray(c.gameRecord)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\nWant: %v\nGot:  %v\n", c.want, got)
		} else {
			//			fmt.Printf("Game record: %v\nMoves: %v\n", c.gameRecord, c.want)
		}
	}
}

func TestTransformingRecordedMoveToStoneStruct(t *testing.T) {
	cases := []struct {
		recordedMove string
		stone Stone
	}{
		{"B[bb]", Stone{1, 1, 1}},
		{"W[ab]", Stone{2, 0, 1}},
		{"B[aa]", Stone{1, 0, 0}},
		{"W[ba]", Stone{2, 1, 0}},
		{"B[cc]", Stone{1, 2, 2}},
	}
	for _, c := range cases {
		got := transformMoveToStone(c.recordedMove)
		if !reflect.DeepEqual(got, c.stone) {
			t.Errorf("\nWant: %v\nGot:  %v\n", c.stone, got)
		} else {
			//			fmt.Printf("Transform %v => %v\n", c.recordedMove, c.stone)
		}
	}
}

func TestTransformingMovesToStones(t *testing.T) {
	cases := []struct {
		moves []string
		stones []Stone
	}{
		{[]string{"B[bb]", "W[ab]", "B[aa]", "W[ba]"}, []Stone{{1, 1, 1}, {2, 0, 1}, {1, 0, 0}, {2, 1, 0}}},
	}
	for _, c := range cases {
		got := transformMovesToStones(c.moves)
		if !reflect.DeepEqual(got, c.stones) {
			t.Errorf("\nWant: %v\nGot:  %v\n", c.stones, got)
		}
	}
}

func TestGettingMovesFromGameRecord(t *testing.T) {
	cases := []struct {
		gameRecord string
		want []Stone
	}{
		{"(;SZ[3]PB[Lee]PW[Alpha]RE[B+1];B[bb];W[ab];B[aa];W[ba])", []Stone{{1, 1, 1}, {2, 0, 1}, {1, 0, 0}, {2, 1, 0}}},
	}
	for _, c := range cases {
		got := getMovesFromGameRecord(c.gameRecord)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\nWant: %v\nGot:  %v\n", c.want, got)
		}
	}
}

func TestWritingGameOnBoard(t *testing.T) {
	cases := []struct {
		gameRecord string
		stonesOnBoard [][]int
	}{
		{"(;SZ[3]PB[Lee]PW[Alpha]RE[B+1];B[bb];W[ab];B[aa];W[ba])", [][]int{{1, 2, 0}, {2, 1, 0}, {0, 0, 0}}},
	}
	for _, c := range cases {
		got := writeGameOnBoard(c.gameRecord)
		if !reflect.DeepEqual(got, c.stonesOnBoard) {
			t.Errorf("\nWant: %v\nGot:  %v\n", c.stonesOnBoard, got)
			printBoardToConsole(got)
		} else {
			printBoardToConsole(got)
		}
	}
}

func assertStonesSlicesEqual(first, second []Stone) bool {
	var result bool
	if len(first) != len(second) {
		return false
	}

	for len(first) > 0 {
		firstElement := first[:1]
		var isInArray bool = false
		for _, secondElement := range second {
			if firstElement[0] == secondElement {
				isInArray = true
				result = true
				continue
			}
		}
		if isInArray == false {
			result = false
		}
		first = first[1:]
	}
	return result
}

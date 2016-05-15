package stone

import (
	"testing"
	"reflect"
)

func TestNew(t *testing.T) {
	cases := []struct {
		color, x, y int
		want Stone
	}{
		{1, 2, 3, Stone{1, 2, 3}},
	}
	for _, c := range cases {
		got := New(c.color, c.x, c.y)

		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\nWant: %v\nGot:  %v", c.want, got)
		}
	}
}

func TestPutOnBoard(t *testing.T) {
	cases := []struct {
		stone Stone
		board, want [][]int
	}{
		{Stone{1, 1, 1}, [][]int{[]int{0, 0, 0}, []int{0, 0, 0}, []int{0, 0, 0}}, [][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{0, 0, 0}}},
		{Stone{2, 1, 2}, [][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{0, 0, 0}}, [][]int{[]int{0, 0, 0}, []int{0, 1, 2}, []int{0, 0, 0}}},
	}
	for _, c := range cases {
		got := c.stone.PutOnBoard(c.board)

		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\nWant: %v\nGot:  %v", c.want, got)
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
		got := c.stone.FindNeighboors(c.inBoard)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\nWant: %v\nGot:  %v\n", c.want, got)
		} else {
			//fmt.Println("\nTest findNeighboors(stone, emptyBoard) function:")
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
		got := c.stone.FindOpponents(c.inBoard)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\nWant: %v\nGot:  %v\n", c.want, got)
		} else {
			//fmt.Println("\nTest findOpponentForStone(stone, emptyBoard) function:")
			//fmt.Printf("stones: %v\nresult: %v\n", c.stone, c.want)
		}
	}
}

func TestFindLibertiesForStone(t *testing.T) {
	cases := []struct {
		stone Stone
		inBoard [][]int
		want []Stone
	}{
		{Stone{1, 1, 1}, [][]int{[]int{0, 2, 0}, []int{0, 1, 0}, []int{0, 0, 0}}, []Stone{{0, 1, 0}, {0, 1, 2}, {0, 2, 1}}},
		{Stone{2, 1, 1}, [][]int{[]int{0, 1, 0}, []int{1, 2, 1}, []int{1, 1, 0}}, []Stone{}},
	}
	for _, c := range cases {
		got := c.stone.FindLiberties(c.inBoard)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\nWant: %v\nGot:  %v\n", c.want, got)
		} else {
			//fmt.Println("\nTest findOpponentForStone(stone, emptyBoard) function:")
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
		got := c.stone.FindFriends(c.inBoard)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\nWant: %v\nGot:  %v\n", c.want, got)
		} else {
			//fmt.Println("\nTest findFriendsForStone(stone, emptyBoard) function:")
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
		got := c.stone.IsInGroup(c.group)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\nWant: %v\nGot:  %v\n", c.want, got)
		} else {
			//fmt.Println("\nTest isStoneInGroup(stone, emptyBoard) function:")
			//fmt.Printf("stones: %v\nresult: %v\n", c.stone, c.want)
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
		got := c.stone.makeGroup(c.inBoard)
		if !assertStonesSlicesEqual(got, c.want) {
			t.Errorf("\nWant: %v\nGot:  %v\n", c.want, got)
		} else {
			//fmt.Println("\nTest makeGroupForStone(stone, emptyBoard) function:")
			//fmt.Printf("stones: %v\nresult: %v\n", c.stone, c.want)
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
		got := c.stone.countLiberties(c.board)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\nWant: %v\nGot:  %v\n", c.want, got)
		} else {
			//fmt.Println("\nTest countStoneLiberties(stone, board):")
			//fmt.Printf("result: %v\n", c.want)
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
		got := TransformMoveToStone(c.recordedMove)
		if !reflect.DeepEqual(got, c.stone) {
			t.Errorf("\nWant: %v\nGot:  %v\n", c.stone, got)
		} else {
			//fmt.Printf("Transform %v => %v\n", c.recordedMove, c.stone)
		}
	}
}

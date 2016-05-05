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
		got := putStoneOnBoard(c.stone, c.inBoard)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\nWant: %v\nGot:  %v\n", c.want, got)
		} else {
			//fmt.Println("\nTest putStoneOnBoard(stone, emptyBoard) function:")
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
		got := assertIllegalMove(c.stone, c.inBoard)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\nWant: %v\nGot:  %v\n", c.want, got)
		} else {
			//fmt.Println("\nTest putStoneOnBoard(stone, emptyBoard) function:")
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
		got := findNeighboors(c.stone, c.inBoard)
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
		got := findOpponentForStone(c.stone, c.inBoard)
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
		got := findFriendsForStone(c.stone, c.inBoard)
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
		got := isStoneInGroup(c.stone, c.group)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\nWant: %v\nGot:  %v\n", c.want, got)
		} else {
			fmt.Println("\nTest makeGroupForStone(stone, emptyBoard) function:")
			fmt.Printf("stones: %v\nresult: %v\n", c.stone, c.want)
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
	}
	for _, c := range cases {
		got := makeGroupForStone(c.stone, c.inBoard)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\nWant: %v\nGot:  %v\n", c.want, got)
		} else {
			fmt.Println("\nTest makeGroupForStone(stone, emptyBoard) function:")
			fmt.Printf("stones: %v\nresult: %v\n", c.stone, c.want)
		}
	}
}

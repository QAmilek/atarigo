package stone

type Stone struct {
	Color int
	X int
	Y int
}

var stoneColors = map[string]int{
	"B": 1,
	"W": 2,
}

var coordinates = map[string]int{
	"a": 0,
	"b": 1,
	"c": 2,
	"d": 3,
	"e": 4,
	"f": 5,
	"g": 6,
	"h": 7,
}

type Stones []Stone

func New(color, x, y int) Stone {
	return Stone{Color: color, X: x, Y: y}
}

func (stone *Stone) PutOnBoard(board [][]int) [][]int {
	board[stone.X][stone.Y] = stone.Color

	return board
}

func (stone *Stone) IsMovePossible(board [][]int) bool {
	placeOnBoard := board[stone.X][stone.Y]

	if placeOnBoard != 0 {
		return false
	}

	return true
}

func (stone *Stone) FindNeighboors(board [][]int) []Stone {
	liberties := []Stone{}

	if stone.Y > 0 {
		liberties = append(liberties, Stone{board[stone.X][stone.Y-1], stone.X, stone.Y - 1})
	}
	if stone.Y < len(board) - 1 {
		liberties = append(liberties, Stone{board[stone.X][stone.Y+1], stone.X, stone.Y + 1})
	}
	if stone.X > 0 {
		liberties = append(liberties, Stone{board[stone.X-1][stone.Y], stone.X - 1, stone.Y})
	}
	if stone.X < len(board) - 1 {
		liberties = append(liberties, Stone{board[stone.X+1][stone.Y], stone.X + 1, stone.Y})
	}

	return liberties
}

func (stone *Stone) FindOpponents(board [][]int) []Stone {
	neighboors := stone.FindNeighboors(board)
	opponents := []Stone{}

	for _, liberty := range neighboors {
		if liberty.Color != 0 && stone.Color != liberty.Color {
			opponents = append(opponents, liberty)
		}
	}

	return opponents
}

func (stone *Stone) FindFriends(board [][]int) []Stone {
	neighboors := stone.FindNeighboors(board)
	friends := []Stone{}

	for _, liberty := range neighboors {
		if stone.Color == liberty.Color {
			friends = append(friends, liberty)
		}
	}

	return friends
}

func (stone *Stone) IsInGroup(group []Stone) bool {
	for _, element := range group {
		if stone.equalTo(element) {
			return true
		}
	}

	return false
}

func (stone *Stone) equalTo(otherStone Stone) bool {
	if stone.X == otherStone.X && stone.Y == otherStone.Y && stone.Color == otherStone.Color {
		return true
	}

	return false
}

func (stone *Stone) MakeGroup(board [][]int) []Stone {
	group := []Stone{}
	toCheck := []Stone{}
	friends := []Stone{}
	group = append(group, Stone{stone.Color, stone.X, stone.Y})
	toCheck = append(toCheck, Stone{stone.Color, stone.X, stone.Y})

	for len(toCheck) > 0 {
		firstToCheck := toCheck[:1]

		friends = firstToCheck[0].FindFriends(board)
		for _, friend := range friends {
			if !friend.IsInGroup(group) {
				group = append(group, friend)
				toCheck = append(toCheck, friend)
			}
		}
		toCheck = toCheck[1:]
	}

	return group
}

func (stone *Stone) FindLiberties(board [][]int) []Stone {
	neighboors := stone.FindNeighboors(board)
	liberties := []Stone{}

	for _, liberty := range neighboors {
		if liberty.Color == 0 {
			liberties = append(liberties, liberty)
		}
	}

	return liberties
}

func (stone *Stone) countLiberties(board [][]int) int {
	liberties := stone.FindLiberties(board)

	return len(liberties)
}

func TransformMoveToStone(move string) Stone {
	color := move[0:1]
	x := move[2:3]
	y := move[3:4]

	s := New(
		stoneColors[color],
		coordinates[x],
		coordinates[y],
	)

	return s
}

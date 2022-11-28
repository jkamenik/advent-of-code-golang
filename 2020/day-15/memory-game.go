package main

type MemoryGame struct {
	// index is the number spoken, values an array of all the times it was spoken
	board      map[int][]int
	turn       int
	lastSpoken int
}

func NewMemoryGame(seed []int) *MemoryGame {
	x := MemoryGame{board: map[int][]int{}, turn: 0}

	for _, value := range seed {
		x.turn++
		x.speak(value)
	}

	return &x
}

func (game *MemoryGame) speak(value int) int {
	turns, ok := game.board[value]
	if !ok {
		turns = []int{game.turn}
	} else {
		turns = append(turns, game.turn)
	}
	game.board[value] = turns
	game.lastSpoken = value

	return value
}

// Turn takes a turn and returns the value for that turn
func (game *MemoryGame) Turn() int {
	game.turn++

	turns, ok := game.board[game.lastSpoken]
	if !ok || len(turns) == 1 {
		return game.speak(0)
	}

	// fmt.Printf("Turns: %v\n", turns)

	prev := turns[len(turns)-2]
	next := turns[len(turns)-1]

	return game.speak(next - prev)
}
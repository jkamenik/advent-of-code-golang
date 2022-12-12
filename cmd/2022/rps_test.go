package twentyTwentyTwo

import "testing"

func TestRPSP1Score(t *testing.T) {
	table := []struct {
		name  string
		round RPS
		score int
	}{
		// Various Draws
		{"draw - rock", RPS{"A", "X"}, 1 + 3},
		{"draw - paper", RPS{"B", "Y"}, 2 + 3},
		{"draw - scissors", RPS{"C", "Z"}, 3 + 3},

		// P1 Wins
		{"p1 wins - rock", RPS{"A", "Z"}, 1 + 6},
		{"p1 wins - paper", RPS{"B", "X"}, 2 + 6},
		{"p1 wins - scissors", RPS{"C", "Y"}, 3 + 6},

		// P1 Loses
		{"p1 loses - rock", RPS{"A", "Y"}, 1},
		{"p1 loses - paper", RPS{"B", "Z"}, 2},
		{"p1 loses - scissors", RPS{"C", "X"}, 3},
	}

	for _, game := range table {
		score := game.round.P1Score()
		if game.score != score {
			t.Errorf("Game %v expected to have a score of %v but had a score of %v", game.name, game.score, score)
		}
	}
}

func TestRPSP2Score(t *testing.T) {
	table := []struct {
		name  string
		round RPS
		score int
	}{
		// Various Draws
		{"draw - rock", RPS{"A", "X"}, 1 + 3},
		{"draw - paper", RPS{"B", "Y"}, 2 + 3},
		{"draw - scissors", RPS{"C", "Z"}, 3 + 3},

		// P1 Wins
		{"p1 wins - rock", RPS{"A", "Z"}, 1},
		{"p1 wins - paper", RPS{"B", "X"}, 2},
		{"p1 wins - scissors", RPS{"C", "Y"}, 3},

		// P1 Loses
		{"p1 loses - rock", RPS{"A", "Y"}, 1 + 6},
		{"p1 loses - paper", RPS{"B", "Z"}, 2 + 6},
		{"p1 loses - scissors", RPS{"C", "X"}, 3 + 6},
	}

	for _, game := range table {
		score := game.round.P2Score()
		if game.score != score {
			t.Errorf("Game %v expected to have a score of %v but had a score of %v", game.name, game.score, score)
		}
	}
}

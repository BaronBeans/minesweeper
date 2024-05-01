package game

import (
	"fmt"
	"math/rand"
)

type Bomb struct {
	X int
	Y int
}

func GenerateBombs(width, height, count int) ([]Bomb, error) {
	bombs := make([]Bomb, 0, count)
BOMB:
	for {
		if len(bombs) == count {
			break
		}
		x := rand.Intn(width)
		y := rand.Intn(height)
		// check if bomb already exists
		for _, bomb := range bombs {
			if bomb.X == x && bomb.Y == y {
				continue BOMB
			}
		}
		bombs = append(bombs, Bomb{x, y})
	}

	if len(bombs) != count {
		return nil, fmt.Errorf("could not generate %d bombs", count)
	}

	return bombs, nil
}

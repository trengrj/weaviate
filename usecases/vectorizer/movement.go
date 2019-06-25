package vectorizer

import "fmt"

// MoveTo moves one vector toward another
func (v *Vectorizer) MoveTo(source []float32, target []float32, weight float32,
) ([]float32, error) {
	return move(source, target, weight, towards)
}

// MoveAwayFrom moves one vector away from another
func (v *Vectorizer) MoveAwayFrom(source []float32, target []float32, weight float32,
) ([]float32, error) {
	return move(source, target, weight, awayFrom)
}

type direction float32

const (
	towards  = direction(1)
	awayFrom = direction(-1)
)

func move(source []float32, target []float32, weight float32, direction direction) ([]float32, error) {
	multiplier := float32(0.5)

	if len(source) != len(target) {
		return nil, fmt.Errorf("movement: vector lengths don't match: got %d and %d", len(source), len(target))
	}

	if weight < 0 || weight > 1 {
		return nil, fmt.Errorf("movement: force must be between 0 and 1: got %f", weight)
	}

	res := make([]float32, len(source), len(source))
	for i, sourceItem := range source {
		res[i] = sourceItem*(1-weight*multiplier) + float32(direction)*target[i]*(weight*multiplier)
	}

	return res, nil

}

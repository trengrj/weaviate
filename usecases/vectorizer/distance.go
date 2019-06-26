package vectorizer

import (
	"fmt"
	"math"
)

// NormalizedDistance between two arbitrary vectors, errors if dimensions don't
// match, will return results between 0 (no distance) and 1 (maximum distance)
func (v *Vectorizer) NormalizedDistance(a, b []float32) (float32, error) {
	sim, err := cosineSim(a, b)
	if err != nil {
		return 1, fmt.Errorf("normalized distance: %v", err)
	}

	return (1 - sim) / 2, nil
}

func cosineSim(a, b []float32) (float32, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("vectors have different dimensions")
	}

	var (
		sumProduct float64
		sumASquare float64
		sumBSquare float64
	)

	for i := range a {
		sumProduct += float64(a[i] * b[i])
		sumASquare += float64(a[i] * a[i])
		sumBSquare += float64(b[i] * b[i])
	}

	return float32(sumProduct / (math.Sqrt(sumASquare) * math.Sqrt(sumBSquare))), nil
}

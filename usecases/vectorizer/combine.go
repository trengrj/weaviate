//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package vectorizer

// CombineVectors combines all of the vector into sum of their parts
func CombineVectors(vectors [][]float32) []float32 {
	maxVectorLength := 0
	for i := range vectors {
		if len(vectors[i]) > maxVectorLength {
			maxVectorLength = len(vectors[i])
		}
	}
	sums := make([]float32, maxVectorLength)
	dividers := make([]float32, maxVectorLength)
	for _, vector := range vectors {
		for i := 0; i < len(vector); i++ {
			sums[i] += vector[i]
			dividers[i]++
		}
	}
	combinedVector := make([]float32, len(sums))
	for i := 0; i < len(sums); i++ {
		combinedVector[i] = sums[i] / dividers[i]
	}

	return combinedVector
}

package weighted_random

import "math/rand"

type num struct {
	val string
	weight int
}

func weighted_random(input []num) string {
	for _, v := range input {
		sum += v.weight
	}

	return rand.Int31n(sum)
}

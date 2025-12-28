package main

import (
	"math"
)

func sigmoid(v []float64) []float64 {
	res := make([]float64, len(v))
	for i, x := range v {
		res[i] = 1 / (1 + math.Exp(-x))
	}

	return res
}

func softmax(v []float64) []float64 {
	maxValue := v[0]
	for _, x := range v {
		if x > maxValue {
			maxValue = x
		}
	}

	res := make([]float64, len(v))
	var sum float64

	for i, x := range v {
		res[i] = math.Exp(x - maxValue)
		sum += res[i]
	}
	for i := range res {
		res[i] /= sum
	}

	return res
}

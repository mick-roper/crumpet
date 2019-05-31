package core

import "math"

func getMax(x []float64) float64 {
	var m float64

	for i := 0; i < len(x); i++ {
		if x[i] > m {
			m = x[i]
		}
	}

	return m
}

func getMin(x []float64) float64 {
	var m float64 = 10000000000 // arbitrarily large number

	for i := 0; i < len(x); i++ {
		if x[i] < m {
			m = x[i]
		}
	}

	return m
}

func getAvg(x []float64) float64 {
	l := (float64)(len(x))
	var sum float64

	for i := 0; i < len(x); i++ {
		sum += (float64)(x[i])
	}

	return sum / l
}

func getStdDev(x []float64) float64 {
	avg := getAvg(x)

	results := make([]float64, len(x))

	for i := 0; i < len(x); i++ {
		a := (x[i]) - avg
		n := a * a
		results[i] = n
	}

	avg = getAvg(results)

	return math.Sqrt(avg)
}
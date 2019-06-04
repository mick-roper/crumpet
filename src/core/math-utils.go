package core

import (
	"math"
	"sort"
)

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
	var n float64

	l := len(x)

	for i := 0; i < l; i++ {
		n = x[i]
		if n < m {
			m = n
		}
	}

	return m
}

func getAvg(x []float64) float64 {
	count := (float64)(len(x))
	var sum float64

	for i := 0; i < len(x); i++ {
		sum += (x[i])
	}

	return sum / count
}

func getStdDev(x []float64) float64 {
	avg := getAvg(x)

	results := make([]float64, len(x))

	for i := 0; i < len(x); i++ {
		n := (x[i]) - avg
		results[i] = n * n
	}

	avg = getAvg(results)

	return math.Sqrt(avg)
}

func getPercentileAverage(x []float64, pc float64) float64 {
	n := make([]float64, len(x))
	copy(n, x)
	sort.Float64s(n)

	floatBoundary := float64(len(x)) * pc
	intBoundary := int(floatBoundary)

	avg := getAvg(n[:intBoundary])

	return avg
}

func getMedian(x []float64) float64 {
	var m float64
	l := len(x)

	if (l > 0) {
		n := make([]float64, l)
		copy(n, x)
		sort.Float64s(n)

		if (l % 2 == 0) {
			m = n[l / 2 -1]
		} else {
			m = n[l / 2]
		}
	}

	return m
}
// Package entropy provides functions for calculating entropy from
// probabilities and sequences.
package entropy

import "math"

func Entropy(p []float64) float64 {
	// probably should assert sum(p) near 1.0
	h := 0.0
	for i := 0; i < len(p); i++ {
		h -= p[i] * math.Log2(p[i])
	}
	return h
}

func SeqEntropy(seq string) float64 {
	// get nucleotide counts
	var a, c, g, t int = 0, 0, 0, 0
	for i := 0; i < len(seq); i++ {
		switch seq[i] {
		case 'A':
			a++
		case 'C':
			c++
		case 'G':
			g++
		case 'T':
			t++
		}
	}

	// convert to slice of probabilities
	var prob []float64
	if a > 0 {
		prob = append(prob, float64(a)/float64(len(seq)))
	}
	if c > 0 {
		prob = append(prob, float64(c)/float64(len(seq)))
	}
	if g > 0 {
		prob = append(prob, float64(g)/float64(len(seq)))
	}
	if t > 0 {
		prob = append(prob, float64(t)/float64(len(seq)))
	}

	return Entropy(prob)
}

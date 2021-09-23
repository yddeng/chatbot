package chatbot

import (
	"errors"
	"math"
)

func Cosine(a []float64, b []float64) (cosine float64, err error) {
	count := 0
	length_a := len(a)
	length_b := len(b)
	if length_a > length_b {
		count = length_a
	} else {
		count = length_b
	}
	sumA := 0.0
	s1 := 0.0
	s2 := 0.0
	for k := 0; k < count; k++ {
		if k >= length_a {
			s2 += math.Pow(b[k], 2)
			continue
		}
		if k >= length_b {
			s1 += math.Pow(a[k], 2)
			continue
		}
		sumA += a[k] * b[k]
		s1 += math.Pow(a[k], 2)
		s2 += math.Pow(b[k], 2)
	}
	if s1 == 0 || s2 == 0 {
		return 0.0, errors.New("Vectors should not be null (all zeros). ")
	}
	return sumA / (math.Sqrt(s1) * math.Sqrt(s2)), nil
}

func Cosine_sim(a, b map[string]int) float64 {
	allworlds := make([]string, 0, len(a)+len(b))
	for v := range a {
		allworlds = append(allworlds, v)
	}
	for v := range b {
		allworlds = append(allworlds, v)
	}

	ca := make([]float64, len(allworlds))
	cb := make([]float64, len(allworlds))
	for i, v := range allworlds {
		ca[i] = float64(a[v])
		cb[i] = float64(b[v])
	}

	if ret, err := Cosine(ca, cb); err == nil {
		return ret
	}
	return 0
}

package main

import (
	"math"
	"sort"
)

func maxFrequency(buckets [][]float64) float64 {
	var max float64
	for _, b := range buckets {
		n := float64(len(b))
		if n > max {
			max = n
		}
	}
	return max
}

func maxInt(values []int) int {
	max := math.MinInt64
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}

func maxFloat64(values []float64) float64 {
	max := -math.MaxFloat64
	for _, v := range values {
		max = math.Max(max, v)
	}
	return max
}

func minFloat64(values []float64) float64 {
	min := math.MaxFloat64
	for _, v := range values {
		min = math.Min(min, v)
	}
	return min
}

func avgFloat64(fs []float64) float64 {
	var sum float64
	for _, f := range fs {
		sum += float64(f)
	}
	return sum / float64(len(fs))
}

func avgInt(fs []int) float64 {
	var sum float64
	for _, f := range fs {
		sum += float64(f)
	}
	return sum / float64(len(fs))
}

func categories(buckets []Bucket) []string {
	var keys []string
	for _, b := range buckets {
		keys = append(keys, b.Category)
	}
	return keys
}

func frequencies(buckets []Bucket) []int {
	var values []int
	for _, b := range buckets {
		values = append(values, b.Frequency)
	}
	return values
}

func frequenciesFloat64(buckets []Bucket) []float64 {
	var values []float64
	for _, b := range buckets {
		values = append(values, float64(b.Frequency))
	}
	return values
}

func percentile(p float64, values []float64) float64 {
	samples := make([]float64, len(values))
	for i, v := range values {
		samples[i] = v
	}
	sort.Float64s(samples)

	val := p * float64(len(samples))
	index := int(math.Ceil(val))

	return samples[index-1]
}

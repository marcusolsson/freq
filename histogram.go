package main

func hist(values []float64, n int) ([][]float64, []float64) {
	max := maxFloat64(values)
	min := minFloat64(values)

	bucketSize := (max - min) / float64(n)

	var edges []float64
	for i := 0; i < n; i++ {
		edges = append(edges, min+float64(i)*bucketSize)
	}

	edges = append(edges, max)

	buckets := make([][]float64, n)

	// Initialize buckets.
	for i := range buckets {
		buckets[i] = make([]float64, 0)
	}

	// Fill buckets.
	for _, val := range values {
		index := int((val - min) / bucketSize)

		if index == n {
			index--
		}

		buckets[index] = append(buckets[index], val)
	}

	return buckets, edges
}

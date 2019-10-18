package main

import "testing"

func TestPercentile(t *testing.T) {
	buckets := []float64{
		30, 30, 31, 31, 32,
		35, 35, 36, 37, 38, 39, 39,
		40, 40, 41, 42, 44, 44, 44, 44,
		45, 45, 46, 47, 48,
		51, 52,
		55, 55, 56,
		61, 63, 64,
		73,
	}

	if got, want := percentile(0.50, buckets), 44.0; got != want {
		t.Errorf("got = %v; want = %v", got, want)
	}
	if got, want := percentile(0.90, buckets), 61.0; got != want {
		t.Errorf("got = %v; want = %v", got, want)
	}
	if got, want := percentile(0.95, buckets), 64.0; got != want {
		t.Errorf("got = %v; want = %v", got, want)
	}
	if got, want := percentile(0.99, buckets), 73.0; got != want {
		t.Errorf("got = %v; want = %v", got, want)
	}
}

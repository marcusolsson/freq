package main

import "testing"

func TestBuildColumn(t *testing.T) {
	for _, tt := range []struct {
		In  float64
		Out string
	}{
		{In: -1.0, Out: " "},
		{In: 0.0, Out: " "},
		{In: 0.1, Out: " "},
		{In: 0.2, Out: "▏"},
		{In: 0.3, Out: "▎"},
		{In: 0.4, Out: "▍"},
		{In: 0.5, Out: "▌"},
		{In: 0.6, Out: "▋"},
		{In: 0.7, Out: "▊"},
		{In: 0.8, Out: "▉"},
		{In: 0.9, Out: "█"},
		{In: 1.0, Out: "█"},
		{In: 2.0, Out: "██"},
		{In: 2.5, Out: "██▌"},
		{In: 5.3, Out: "█████▎"},
	} {
		if got := column(tt.In); got != tt.Out {
			t.Errorf("(%v) got = %q; want = %q", tt.In, got, tt.Out)
		}
	}
}

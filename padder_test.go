package main

import (
	"testing"
)

func TestPadding_Fill(t *testing.T) {
	for _, tt := range []struct {
		In    string
		Width int
		Out   string
	}{
		{
			In:    "foobar",
			Width: 10,
			Out:   "foobar    ",
		},
		{
			In:    "–",
			Width: 5,
			Out:   "–    ",
		},
		{
			In:    "世界",
			Width: 6,
			Out:   "世界  ",
		},
	} {
		t.Run("", func(t *testing.T) {
			got := fill(tt.In, tt.Width)

			if got != tt.Out {
				t.Fatalf("got = %q; want = %q", got, tt.Out)
			}
		})
	}
}

func TestPadding_Justify(t *testing.T) {
	for _, tt := range []struct {
		In    string
		Width int
		Out   string
	}{
		{
			In:    "foobar",
			Width: 10,
			Out:   "    foobar",
		},
		{
			In:    "–",
			Width: 5,
			Out:   "    –",
		},
		{
			In:    "世界",
			Width: 6,
			Out:   "  世界",
		},
	} {
		t.Run("", func(t *testing.T) {
			got := just(tt.In, tt.Width)

			if got != tt.Out {
				t.Fatalf("got = %q; want = %q", got, tt.Out)
			}
		})
	}
}

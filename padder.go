package main

import (
	"strings"

	runewidth "github.com/mattn/go-runewidth"
)

func fill(s string, w int) string {
	return s + strings.Repeat(" ", w-runewidth.StringWidth(s))
}

func just(s string, w int) string {
	return strings.Repeat(" ", w-runewidth.StringWidth(s)) + s
}

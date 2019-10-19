package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestPrintBarChart(t *testing.T) {
	h := []bucket{
		{Category: "foobar", Frequency: 3},
		{Category: "世界", Frequency: 2},
		{Category: "–", Frequency: 1},
	}

	var buf bytes.Buffer

	printBarChart(&buf, h, 3, false)

	want := `
foobar ███▏ 3
世界   ██▏ 2
–      █▏ 1
`

	w := strings.TrimLeft(want, "\n")

	if buf.String() != w {
		t.Errorf("unexpected output:\ngot:\n%s\n\nwant:\n%s", buf.String(), w)
	}
}

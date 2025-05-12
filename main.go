package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"

	runewidth "github.com/mattn/go-runewidth"
	flag "github.com/spf13/pflag"
)

func main() {
	usage := `Usage: freq [options] [-]

Options:
  -h, --help            Show this help message and exit.
  --histogram           Prints a histogram. Requires numeric-only input.
  --summary             Prints a summary.
  --sort-by=<order>     Specify sort order: one of count, label (default is count).
  --desc                Sort buckets in descending order (default is false).
  --justify             Aligns categories and ranges to the right.
  --column-width=<num>  Width of the largest bin (default is 30).
  --buckets=<num>       Number of histogram buckets (default is 10).
`

	var (
		histogram   = flag.BoolP("histogram", "", false, "")
		summary     = flag.BoolP("summary", "", false, "")
		sortBy      = flag.StringP("sort-by", "", "count", "")
		columnWidth = flag.Float64P("column-width", "", 30, "")
		desc        = flag.BoolP("desc", "", false, "")
		justify     = flag.BoolP("justify", "", false, "")
		buckets     = flag.IntP("buckets", "", 10, "")
	)

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
	}

	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)

	var vals []string
	for scanner.Scan() {
		vals = append(vals, strings.TrimSpace(scanner.Text()))
	}

	if len(vals) == 0 {
		os.Exit(0)
	}

	if *histogram {
		// Histograms requires numerical data, so we need to make sure every
		// line is a number.
		var samples []float64

		for _, v := range vals {
			sample, err := strconv.ParseFloat(v, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, "found non-numerical input:", v)
				os.Exit(1)
			}

			samples = append(samples, sample)
		}

		buckets, edges := hist(samples, *buckets)

		printHistogram(os.Stdout, buckets, edges, *columnWidth, *justify)

		if *summary {
			printSummary(os.Stdout, samples)
		}
	} else {
		buckets := categoricalBuckets(vals)

		// In contrast to histograms, bar charts allow you to sort the bars.
		sortBuckets(buckets, *sortBy, *desc)

		printBarChart(os.Stdout, buckets, *columnWidth, *justify)

		if *summary {
			printSummary(os.Stdout, frequenciesFloat64(buckets))
		}
	}
}

// printHistogram displays a histogram. The bar width determines the width of
// the widest bar. Labels can optionally be right justified.
func printHistogram(out io.Writer, buckets [][]float64, edges []float64, barWidth float64, justify bool) {
	var labels []string
	for i := 0; i < len(edges)-1; i++ {
		labels = append(labels, fmt.Sprintf("%.6g-%.6g", edges[i], edges[i+1]))
	}

	var (
		maxFreq    = maxFrequency(buckets)
		labelWidth = maxStringWidth(labels)
	)

	for idx, bucket := range buckets {
		var (
			normalizedWidth = float64(len(bucket)) / maxFreq
			width           = normalizedWidth * barWidth
			prefix          = paddedString(labels[idx], labelWidth, justify)
		)

		fmt.Fprintf(out, "%s %s %d\n", prefix, column(width), len(bucket))
	}
}

// printBarChart displays a bar chart. The bar width determines the width of
// the widest bar. Categories can optionally be right justified.
func printBarChart(out io.Writer, buckets []bucket, barWidth float64, justify bool) {
	var (
		maxFreq       = maxInt(frequencies(buckets))
		categoryWidth = maxStringWidth(categories(buckets))
	)

	for _, bucket := range buckets {
		var (
			normalizedWidth = float64(bucket.Frequency) / float64(maxFreq)
			width           = normalizedWidth * barWidth
			prefix          = paddedString(bucket.Category, categoryWidth, justify)
		)

		fmt.Fprintf(out, "%s %s %d\n", prefix, column(width), bucket.Frequency)
	}
}

// paddedString returns the string justified in a string of given width.
func paddedString(str string, width int, justify bool) string {
	if justify {
		return just(str, width)
	}

	return fill(str, width)
}

// printSummary displays additional statistics about the dataset.
func printSummary(out io.Writer, numbers []float64) {
	stats := []string{
		fmt.Sprintf("%s=%g", "p50", percentile(0.5, numbers)),
		fmt.Sprintf("%s=%g", "p90", percentile(0.9, numbers)),
		fmt.Sprintf("%s=%g", "p95", percentile(0.95, numbers)),
		fmt.Sprintf("%s=%g", "p99", percentile(0.99, numbers)),
		fmt.Sprintf("%s=%g", "min", minFloat64(numbers)),
		fmt.Sprintf("%s=%g", "max", maxFloat64(numbers)),
		fmt.Sprintf("%s=%g", "avg", avgFloat64(numbers)),
	}

	fmt.Fprintln(out)
	fmt.Fprintln(out, "summary:")
	fmt.Fprintln(out, " "+strings.Join(stats, ", "))
}

var boxes = []string{"▏", "▎", "▍", "▌", "▋", "▊", "▉", "█"}

// columns returns a horizontal bar of a given size.
func column(size float64) string {
	fraction := size - math.Floor(size)
	index := int(fraction * float64(len(boxes)))

	return strings.Repeat(boxes[len(boxes)-1], int(size)) + boxes[index]
}

// maxStringWidth returns the width of the widest string in a string slice. It
// supports CJK through the go-runewidth package.
func maxStringWidth(strs []string) int {
	var max int

	for _, str := range strs {
		w := runewidth.StringWidth(str)
		if w > max {
			max = w
		}
	}

	return max
}

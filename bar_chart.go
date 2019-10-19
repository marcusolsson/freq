package main

import "sort"

type bucket struct {
	Category  string
	Frequency int
}

func categoricalBuckets(vals []string) []bucket {
	buckets := make(map[string]int)

	for _, v := range vals {
		if _, ok := buckets[v]; !ok {
			buckets[v] = 0
		}
		buckets[v]++
	}

	res := make([]bucket, 0, len(buckets))
	for k, v := range buckets {
		res = append(res, bucket{Category: k, Frequency: v})
	}

	return res
}

func sortBuckets(buckets []bucket, orderBy string, desc bool) {
	var sorter sort.Interface

	switch orderBy {
	case "category":
		sorter = byCategory(buckets)
	default:
		sorter = byFrequency(buckets)
	}

	if desc {
		sorter = sort.Reverse(sorter)
	}

	sort.Sort(sorter)
}

type byCategory []bucket

func (b byCategory) Len() int           { return len(b) }
func (b byCategory) Less(i, j int) bool { return b[i].Category < b[j].Category }
func (b byCategory) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

type byFrequency []bucket

func (b byFrequency) Len() int           { return len(b) }
func (b byFrequency) Less(i, j int) bool { return b[i].Frequency < b[j].Frequency }
func (b byFrequency) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

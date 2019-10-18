package main

import "sort"

type Bucket struct {
	Category  string
	Frequency int
}

func bucketizeCategories(vals []string) []Bucket {
	buckets := make(map[string]int)
	for _, v := range vals {
		if _, ok := buckets[v]; !ok {
			buckets[v] = 0
		}
		buckets[v] = buckets[v] + 1
	}

	var res []Bucket
	for k, v := range buckets {
		res = append(res, Bucket{Category: k, Frequency: v})
	}
	return res
}

func sortBuckets(buckets []Bucket, orderBy string, desc bool) {
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

type byCategory []Bucket

func (a byCategory) Len() int           { return len(a) }
func (a byCategory) Less(i, j int) bool { return a[i].Category < a[j].Category }
func (a byCategory) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type byFrequency []Bucket

func (a byFrequency) Len() int           { return len(a) }
func (a byFrequency) Less(i, j int) bool { return a[i].Frequency < a[j].Frequency }
func (a byFrequency) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

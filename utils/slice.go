package utils

import "sort"

type OrderDirection = byte

const (
	Desc OrderDirection = 0
	Asc  OrderDirection = 1
)

func SortStringSlice(list []string, dir OrderDirection) {
	if dir == Asc {
		sort.Strings(list)
	} else {
		sort.Sort(sort.Reverse(sort.StringSlice(list)))
	}
}

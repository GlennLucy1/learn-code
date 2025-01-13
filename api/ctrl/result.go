package ctrl

import (
	"strconv"
)

type KeyVal struct {
	Key string
	Val int
}

type OrderMap struct {
	slice []KeyVal
}

func NewOrderMap() *OrderMap {
	return &OrderMap{
		slice: make([]KeyVal, 0),
	}
}

func (o *OrderMap) Set(k string, v int) {
	if _, idx, ok := o.Check(k); ok {
		o.slice[idx].Val = v
	} else {
		o.slice = append(o.slice, KeyVal{
			Key: k,
			Val: v,
		})
	}
}

func (o *OrderMap) Get(k string) int {
	for _, v := range o.slice {
		if v.Key == k {
			return v.Val
		}
	}
	return 0
}

func (o *OrderMap) Check(k string) (int, int, bool) {
	for i, v := range o.slice {
		if v.Key == k {
			return v.Val, i, true
		}
	}
	return 0, 0, false
}

func Speculate(detect []string, right []string) []string {
	final := make([]string, 0)
	rightMap := NewOrderMap()

	for _, v := range detect {
		rightMap.Set(v, 1)
	}

	that := ""
	for _, v := range right {
		if val, _, ok := rightMap.Check(v); ok {
			val += 1
			rightMap.Set(v, val)
		} else {
			that = v
		}
	}

	for _, v := range rightMap.slice {
		if v.Val == 2 {
			final = append(final, v.Key)
		}
		if v.Val == 1 {
			final = append(final, that)
		}
	}

	return final
}

func CleanResult(detect []string) []string {
	if len(detect) >= 4 {
		return detect[:4]
	}
	return nil
}

func CalcRate(detect []string, right []string) int {
	rightNum := 0
	for _, v := range right {
		for _, vv := range detect {
			if v == vv {
				rightNum++
			}
		}
	}
	return rightNum
}

func GetOrder(detect []string, right []string) string {
	final := ""
	rightMap := make(map[string]int)
	for i, v := range detect {
		rightMap[v] = i + 1
	}
	for _, v := range right {
		final += strconv.Itoa(rightMap[v])
	}
	return final
}

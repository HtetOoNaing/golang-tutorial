package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip:   0,
	}
	return b
}

func (b bill) format() string {
	fs := "Bill breakdown: \n"
	total := 0.0
	for k, v := range b.items {
		fs += fmt.Sprintf("%-10v ...$%v \n", k+" :", v)
		total += v
	}
	fs += fmt.Sprintf("%-10v ...$%v \n", "total :", total)
	return fs
}

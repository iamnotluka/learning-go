package main

import "fmt"

type bill struct {
	name string
	items map[string]float64
	tip float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name: name,
		items: map[string]float64{"pie": 6, "cake": 5},
		tip: 0,
	}

	return b
}

// receiver function demo (format the bill)
func (b bill) format() string {
	s := "bill breakdown: \n"
	var total float64 = 0

	for k, v := range b.items {
		s += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
		total += v
	}

	s += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)

	return s
}
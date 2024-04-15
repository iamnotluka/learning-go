package main

import (
	"fmt"
	"os"
)

type bill struct {
	name string
	items map[string]float64
	tip float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name: name,
		items: map[string]float64{},
		tip: 0,
	}

	return b
}

// receiver function demo (format the bill)
func (b bill) format() string {
	s := "bill breakdown: \n"
	var total float64 = b.tip

	s += fmt.Sprintf("%-25v ...$%0.2f\n", "tip:", b.tip)

	for k, v := range b.items {
		s += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
		total += v
	}

	s += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)

	return s
}

// update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("./bills/"+b.name+".txt", data, 0644)

	if (err != nil) {
		panic(err)
	}

	fmt.Println("Bill was saved to file")
}
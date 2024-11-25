package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// NewBill creates a new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// Format the bill
func (b *bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	// List items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
		total += v
	}

	// Add tip
	fs += fmt.Sprintf("%-25v ...$%v\n", "tip:", b.tip)

	// Total
	fs += fmt.Sprintf("%-25v ...$%v\n", "total:", total+b.tip)

	return fs
}

// UpdateTip updates the tip of the bill
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// AddItem adds an item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// RemoveItem removes an item from the bill
func (b *bill) removeItem(name string) {
	delete(b.items, name)
}

// SaveBill saves the bill to a file
func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Bill was saved to file")
}
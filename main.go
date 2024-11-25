package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	myBill := createBill()

	promptOptions(myBill)
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the bill name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	b := newBill(name)
	fmt.Println("Created the bill -", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)
	fmt.Println(opt)

	switch opt {
		case "a":
			name, _ := getInput("Item name: ", reader)
			price, _ := getInput("price: ", reader)

			p, err := strconv.ParseFloat(price, 64)
			if err != nil {
				fmt.Println("The price must be a number")
				promptOptions(b)
			}
			b.addItem(name, p)

			fmt.Println(name, price)
			promptOptions(b)
		case "t": 
			tip, _ := getInput("Enter tip amount ($): ", reader)

			p, err := strconv.ParseFloat(tip, 64)
			if err != nil {
				fmt.Println("The price must be a number")
				promptOptions(b)
			}
			b.updateTip(p)

			fmt.Println("You tipped: ", tip)
			promptOptions(b)
		case "s":
			b.save()
			fmt.Println("Bill saved to", b.name, ".txt")
		default: 
			fmt.Println("That is not a valid option")
			promptOptions(b)
	}
}
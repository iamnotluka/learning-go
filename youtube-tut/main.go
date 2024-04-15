package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var someName string = "someName"

func main() {
	// STRINGS
	var nameOne string = "luka"
	var nameTwo = "relja"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "lukaUpdated"
	nameThree = "tamara"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "ethan"

	fmt.Println(nameFour)

	// INTEGERS
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	var signed16bitInt int16 = 252
	var signed8bitInt int8 = -128
	var unsigned8bitInt uint8 = 255

	fmt.Println(signed16bitInt, signed8bitInt, unsigned8bitInt)

	// FLOATS
	var signed32bitFloat = 25.98
	var signed64bitFloat = 84838483848384.88978798 // mostly used

	unspecifiedFloatDefaultsTo64bit := 1.5

	fmt.Println(signed32bitFloat, signed64bitFloat, unspecifiedFloatDefaultsTo64bit)

	// Prints without newline char (use \n for a new line
	fmt.Print("Prints without the '/n' character at the end. ")
	fmt.Print("Line added to the previous one.\n")
	fmt.Print("This line in a new line.")

	// Printing variables
	fmt.Println("my age is", ageOne, "and the name is", nameFour)

	// Formatted strings
	fmt.Printf("myage is %q and my name is %q\n", ageThree, nameThree)

	// %T type
	fmt.Printf("ageThree is of type %T \n", ageThree)
	// %f float, %0.Xf (X dec points)
	fmt.Printf("ageTwo is %0.2f", signed32bitFloat)

	// Sprintf (save formatted strings)
	var savedString = fmt.Sprint("my age is", ageOne, "and the name is ", nameFour)

	fmt.Println(savedString)

	// ARRAYS (fixed length)
	var ages [3]int = [3]int{1, 2, 3}

	var shorterAges = [3]int{1, 2, 3}

	otherDeclaration := [3]int{1, 2, 3}

	names := [4]string{"Joshy", "Mario", "Peach", "Bowser"}
	names[1] = "Luka"
	fmt.Println(ages, len(ages))
	fmt.Println(shorterAges, len(shorterAges))
	fmt.Println(otherDeclaration, len(otherDeclaration))
	fmt.Println(names, len(names))

	// SLICES (using array under the hood)
	var scores = []int{100, 50, 60}
	scores[1] = 1
	scores = append(scores, 58)

	fmt.Println(scores, len(scores))

	// SLICE RANGES
	rangeOne := names[1:3]
	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)
	rangeTwo := names[2:]
	fmt.Println(rangeTwo)
	rangeThree := names[:3]
	fmt.Println(rangeThree)

	greeting := "hi friends!"

	fmt.Println(strings.Contains(greeting, "hi"))
	fmt.Println(strings.Contains(greeting, "Hi"))
	fmt.Println(strings.ReplaceAll(greeting, " ", "_"))
	fmt.Println((strings.ToUpper(greeting)))
	fmt.Println(strings.Index(greeting, "fr"))
	fmt.Println(strings.Split(greeting, " "))

	more_ages := []int{12, 15, 16, 2, 50, 20, 60}

	sort.Ints(more_ages)
	fmt.Println(more_ages)

	more_names := []string{"hello", "hi"}

	fmt.Println(sort.SearchStrings(more_names, "hello"))

	// LOOPS
	x := 0
	for x < 5 {
		fmt.Println("value of x is:", x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("value of i is:", i)
	}

	for index, value := range names {
		fmt.Printf("the value %v is at index %v\n", index, value)
		value = "new" // Doesn't update, obtains the value in the new variable
	}

	fmt.Println(names)

	// BOOLEANS <= >= == !=
	condition1 := 30 < 40
	condition2 := 30 == 40

	if (condition1) {
		fmt.Println("condition1 is true")
	} else if !condition2 {
		fmt.Println("condition2 is true")
	} else {
		fmt.Println("This runs if condition 1 and 2 are false.")
	}

	sayGreeting("Luka")
	sayBye("Luka")
	cycleNames([]string{"Luka", "Relja", "Tamara"}, sayGreeting)
	cycleNames([]string{"Luka", "Relja", "Tamara"}, sayBye)

	sayHello("Mario")

	for _, v := range points {
		fmt.Println(v)
	}

	showScore()

	// MAPS
	menu := map[string]float64{
		"soup": 4.99,
		"pie": 7.99,
		"salad": 6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["soup"])

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// Integers as key type
	phonebook := map[int]string{
		1: "luka",
		2: "relja",
		3: "tamara",
	}
	fmt.Println(phonebook)
	fmt.Println(phonebook[1])

	phonebook[1] = "luka_changed"
	fmt.Println(phonebook)

	// 1 group variables: Strings, ints, floats, booleans, arrays and structs (non-pointer values)
	luka_name := "luka"

	luka_name = updateName(luka_name)

	fmt.Println(luka_name)

	// 2 group variables: Slices, maps, functions (pointer values)
	fmt.Println(menu)
	updateMenu(menu)
	fmt.Println(menu)

	// POINTERS
	pointer_demo := "pointer_demo"
	fmt.Println(&pointer_demo)
	fmt.Println(pointer_demo)

	pointer_demo_mem := &pointer_demo

	updateNameMem(pointer_demo_mem)

	fmt.Println(pointer_demo)
	
	// STRUCT & CUSTOM TYPE
	mybill := newBill("mario's bill")

	fmt.Println(mybill)

	// Receiver functions
	formatterBill := mybill.format()
	fmt.Println(formatterBill)

	mybill.updateTip(1)
	mybill.addItem("item", 3)
	formatterBill = mybill.format()
	fmt.Println(formatterBill)

	// User input
	users_bill := createBill()
	promptOptions(&users_bill)
	fmt.Println(users_bill)
}

func updateNameMem(x *string) {
	*x = "updated_name_in_memory"
}

func updateMenu(x map[string]float64) {
	x["soup"] = 55
}
func updateName(n string) string {
	n = "n_updated"
	return n
}

func sayGreeting(n string) {
	fmt.Printf("Good morning %v\n", n)
}

func sayBye(n string) {
	fmt.Printf("Bye %v\n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func promptOptions(b  *bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)
	
	switch opt {
	case "a":
		fmt.Println("Adding item...")
		item, _ := getInput("What's the item name: ", reader)
		price, _ := getInput("What's the item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)

		if (err != nil) {
			fmt.Println("The price must be a number.")
			promptOptions(b)
			break
		}

		b.addItem(item, p)

		fmt.Println("added tiem", item, "with price", price)
		promptOptions(b)
	case "t":
		tip, _ := getInput("What's is the tip: ", reader)

		t, err := strconv.ParseFloat(tip, 64)

		if (err != nil) {
			fmt.Println("The tip must be a number.")
			promptOptions(b)
			break
		}

		b.updateTip(t)

		fmt.Println("added tip", tip)
		promptOptions(b)
	case "s":
		fmt.Println("saving bill")
	default:
		fmt.Println("Not a valid option")
		promptOptions(b)
	}
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created bill -", b.name)

	return b
}
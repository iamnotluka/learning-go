package main

import "fmt"

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

}

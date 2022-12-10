package chapter3

import (
	"errors"
	"fmt"
	"log"
	"math"
	"reflect"
)

var metersPerLiter float64

// %f
// %d
// %s
// %t
// %v
// %#v
// %T
// %%
func Printf() {
	resultString := fmt.Sprintf("About on-third: %0.2f", 1.0/3.0)
	fmt.Println(resultString)

	fmt.Printf("A float: %f\n", 3.1415)
	fmt.Printf("An integer: %d\n", 14)
	fmt.Printf("A String: %s\n", "hello")
	fmt.Printf("A boolean: %t\n", false)
	fmt.Printf("Value: %v %v %v\n", 1.2, "\t", true)
	fmt.Printf("Value: %#v %#v %#v\n", 1.2, "\t", true)
	fmt.Printf("Type: %T %T %T\n", 1.2, "\t", true)
	fmt.Printf("Percent sign: %%\n")
	fmt.Printf("%v %v %v\n", "", "\t", "\n")
	fmt.Printf("%#v %#v %#v\n", "", "\t", "\n")

	fmt.Printf("%12s | %s\n", "Product", "Cost in Cents")
	fmt.Println("----------------------------------")
	fmt.Printf("%12s | %2d\n", "Stamp", 50)
	fmt.Printf("%12s | %2d\n", "Paper Clips", 5)
	fmt.Printf("%12s | %2d\n", "Tape", 99)
}

func DisplayArea() {
	metersPerLiter = 10.0
	ret, err := PaintNeeded(4.2, 2.0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f liters needed\n", ret)

	ret, err = PaintNeeded(5.2, 3.5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f liters needed\n", ret)

	err = errors.New("height can't be negative")
	fmt.Println(err)

	err = fmt.Errorf("a height of %0.2f is invalid", -2.3333)
	fmt.Println(err)

	val1, val2 := floatParts(23.4)
	fmt.Println("val1 is ", val1, "val2 is ", val2)

	root, err := squareRoot(3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%0.3f\n", root)
	}

	quotient, err := divide(5.6, 0.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%0.2f\n", quotient)
	}

	amount := 6
	doubleByValue(amount)
	fmt.Println("double amount is = ", amount)
	// point address
	fmt.Println("the amount address is = ", &amount)
	var myInt int = 4
	var myIntPointer *int = &myInt
	fmt.Println("myInt address is = ", myIntPointer)
	fmt.Println("myInt type is = ", reflect.TypeOf(myIntPointer))
	fmt.Println("myInt value is = ", *myIntPointer)
	var myFloat float64 = 3.2
	var myFloatPointer *float64 = &myFloat
	fmt.Println("myFloat address is = ", myFloatPointer)
	fmt.Println("myFloat type is = ", reflect.TypeOf(myFloatPointer))
	fmt.Println("myFloat value is = ", *myFloatPointer)
	var myBool bool = true
	var myBoolPointer *bool = &myBool
	fmt.Println("myBool address is = ", myBoolPointer)
	fmt.Println("myBool type is = ", reflect.TypeOf(myBoolPointer))
	fmt.Println("myBool value is = ", *myBoolPointer)
	*myBoolPointer = false
	fmt.Println("myBool value is = ", myBool)
	printPointer(myBoolPointer)

	doubleByRef(&amount)
	fmt.Println("amount double by ref is ", amount)

	truth := true
	negate(&truth)
	fmt.Println("truth negate is = ", truth)
	lies := false
	negate(&lies)
	fmt.Println("lies negate is = ", lies)
}

func createPointer() *float64 {
	var myFloat = 9.4
	return &myFloat
}

func printPointer(myBoolPointer *bool) {
	fmt.Println(*myBoolPointer)
}

func negate(myBoolean *bool) {
	*myBoolean = !*myBoolean
}
func SayHi() {
	fmt.Println("Hi!")
}

func RepeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(line)
	}
}

func PaintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}

	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}

	area := width * height
	return area / metersPerLiter, nil
}

func double(number float64) float64 {
	return number * 2
}

func status(grade float64) string {
	if grade < 60.0 {
		return "failing"
	}
	return "passing"
}

func floatParts(number float64) (integerPart int, fractionPart float64) {
	wholeNumber := math.Floor(number)
	return int(wholeNumber), number - wholeNumber
}

func squareRoot(number float64) (float64, error) {
	if number < 0 {
		return 0, fmt.Errorf("can't get square root of negative number")
	}
	return math.Sqrt(number), nil
}

func divide(dividend float64, divisor float64) (float64, error) {
	if divisor == 0.0 {
		return 0, errors.New("can't divide by 0")
	}
	return dividend / divisor, nil
}

func doubleByValue(number int) {
	number *= 2
}

func doubleByRef(number *int) {
	*number *= 2
}

package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"strings"
	"time"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool
var i, j int = 1, 2

var ToBe bool = false
var MaxInt uint64 = 1<<64 - 1
var z complex128 = cmplx.Sqrt(-5 + 12i)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	// fmt.Println("hello world")
	// rand.Seed(time.Now().UnixNano())
	fmt.Println("My favorite a number is ", rand.Float64())
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println("Pi = ", math.Pi)
	fmt.Println("add(1, 2) = ", add(1, 2))
	a, b := swap("hello", "world")
	fmt.Println("swap(hello, world) = ", a, b)
	e, d := split(18)
	fmt.Println("split(17) = ", e, d)

	var i int
	fmt.Println(i, c, python, java)

	// var c, python, java = true, false, "no!"
	c, python, java := true, false, "yes!"
	k := 3
	fmt.Println(c, python, java, i, j, k)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var z int
	var q float64
	var f bool
	var s string
	fmt.Printf("%v %v %v %q\n", z, q, f, s)

	v := 42
	w := float64(v)
	u := uint64(w)
	fmt.Println(v, w, u)

	var x1, y1 int = 3, 4
	var f1 float64 = math.Sqrt(float64(x1*x1 + y1*y1))
	var z1 uint = uint(f1)
	fmt.Println(x1, y1, f1, z1)

	v1 := 3.14
	fmt.Printf("v is of type %T\n", v1)

	const Pi = 3.14
	const World = "world"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(1 << 2))
	fmt.Println(needFloat(1 << 2))
	fmt.Println(needFloat(1 << 100))

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum1 := 1
	for sum1 < 1000 {
		sum1 += sum1
	}
	fmt.Println(sum1)

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))

	fmt.Println("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("Into two days.")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() > 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	// defer fmt.Println("world")

	// fmt.Println("hello")
	// fmt.Println("world2")

	// fmt.Println("counting")

	// for i := 0; i < 10; i += 1 {
	// 	defer fmt.Println(i)
	// }

	// fmt.Println("Done")

	i1, j1 := 42, 2701

	p := &i1
	fmt.Println(*p)
	*p = 27
	fmt.Println("i1 = ", *p)

	p = &j1
	*p = *p / 37
	fmt.Println("j1 = ", j1)

	v2 := Vertex{1, 2}
	v2.X = 4
	p1 := &v2
	p1.X = 1e9

	fmt.Printf("Type %T Value %v\n", Vertex{1, 2}, Vertex{1, 2})
	fmt.Printf("Type %T Value %v\n", v2, v2)

	vec1 := Vertex{1, 2}
	vec2 := Vertex{X: 1}
	vec3 := Vertex{}
	vecp := &Vertex{1, 2}
	fmt.Println(vec1, vec2, vec3, vecp)

	// var a1 [2]string
	// a1[0] = "Hello"
	// a1[1] = "World"
	// fmt.Println(a[0], a[1])
	// fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s1 []int = primes[1:4]
	fmt.Println(s1)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a1 := names[0:2]
	b1 := names[1:3]
	fmt.Println(a1, b1)

	b1[0] = "XXX"
	fmt.Println(a1, b1)
	fmt.Println(names)

	q1 := []int{2, 3, 5, 6, 11, 13}
	fmt.Println(q1)

	r1 := []bool{true, false, true, true, false, true}
	fmt.Println(r1)

	s2 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(s2)

	s3 := []int{2, 3, 5, 7, 11, 13}

	s4 := s3[1:4]
	fmt.Println(s4)

	s5 := s3[:2]
	fmt.Println(s5)

	s6 := s3[1:]
	fmt.Println(s6)

	s7 := []int{1, 2, 3, 4, 5, 6, 7}
	printSlice(s7)

	s7 = s7[:0]
	printSlice(s7)

	s7 = s7[:4]
	printSlice(s7)

	s7 = s7[2:]
	printSlice(s7)

	s7 = s7[:5]
	printSlice(s7)

	var s8 []int
	fmt.Println(s8, len(s8), cap(s8))

	if s8 == nil {
		fmt.Println("nill!")
	}

	a2 := make([]int, 5)
	printSliceWithName("a2", a2)

	b2 := make([]int, 0, 5)
	printSliceWithName("b2", b2)

	c2 := b2[:2]
	printSliceWithName("c2", c2)

	d2 := c2[2:5]
	printSliceWithName("d2", d2)

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i += 1 {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	var s9 []int
	printSlice(s9)

	s9 = append(s9, 0)
	printSlice(s9)

	s9 = append(s9, 1)
	printSlice(s9)

	s9 = append(s9, 2, 3, 4, 5)
	printSlice(s9)

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow1 := make([]int, 10)
	for i := range pow1 {
		pow1[i] = 1 << uint(i)
	}

	for _, value := range pow1 {
		fmt.Printf("%d\n", value)
	}

	type Vertex1 struct {
		Lat, Long float64
	}

	var m = map[string]Vertex1{
		"Bell Labs": {
			290.99, -923.23,
		},
		"Google": {
			23.34, -122.34,
		},
	}
	fmt.Println(m)

	m1 := make(map[string]int)

	m1["Answer"] = 32
	fmt.Println(m1)
	m1["Answer"] = 34
	fmt.Println(m1)

	delete(m1, "Answer")
	fmt.Println(m1)

	v3, ok3 := m1["Answer"]
	fmt.Println("The value: ", v3, "Present?", ok3)

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i += 1 {
		fmt.Println(pos(i), neg(-2*i))
	}

	v4 := Vertex{3, 4}
	fmt.Println(v4.Abs())

}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func printSlice(s []int) {
	fmt.Printf("len = %d cap = %d %v\n", len(s), cap(s), s)
}

func printSliceWithName(s string, x []int) {
	fmt.Printf("%s len = %d cap = %d %v\n", s, len(x), cap(x), x)
}

type Vertex struct {
	X float64
	Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

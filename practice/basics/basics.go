package main

import (
	"fmt"
	"math"
	"reflect"
	"time"
)

func basic_values() {
	fmt.Println("Hello " + "World")

	fmt.Println("1+1 =", 1+1, "2+2", 2+2)
	fmt.Println("3.0/7.0=", 3.0/7.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

func basic_variables() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}

func basic_const(s string) {
	fmt.Println(s)
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))

	//const l int = 10
	var l int = 10
	//l += 1
	fmt.Println(l)
}

func basic_for_loop() {
	i := 10
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 5; n < 12; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

func basic_if_else() {
	if 2 > 1 {
		fmt.Println("2 > 1")
	}

	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			fmt.Println("i is even")
		} else if i == 3 {
			fmt.Println("i is 3")
		} else {
			fmt.Println("i is odd")
		}
	}
}

func basic_switch() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	whatAmI := func(l interface{}) {
		switch t := l.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

func basic_array() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	// c := [3][3]int This wont work, cant dynamically set multi-dimensional array
	var c [3][3]int
	fmt.Println(c)
	c[0][1] = 10
	fmt.Println(c)
}

func basic_slices() {

	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	g := []int{1}
	// g := []int{1, "a"} Cant have multiple types
	// g := []int Cant dynamically create a slice without a value
	fmt.Println(reflect.TypeOf(g))
	fmt.Println(reflect.TypeOf(twoD))
	fmt.Println(reflect.ValueOf(twoD).Kind())
	fmt.Println(reflect.ValueOf(g).Kind())
}

func basic_maps() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k1"]
	fmt.Println("prs:", prs)

	x, _ := m["k1"]
	fmt.Println("x", x)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	map_type := map[string]map[string]int{"a": {"b": 1}}
	fmt.Println("map_type:", map_type)
	fmt.Println("map_type['a']['b']:", map_type["a"]["b"])
	a, b := map_type["a"]
	fmt.Println(a)
	fmt.Println(b)

}

func basic_range() {

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

func basic_function_add(a int, b int) int {
	return a + b
}

func basic_function_add_add(a, b, c int, d string) int {

	fmt.Println(d)
	return a + b + c
}

func basic_func_mult_return() (int, int) {
	return 1, 5
}

func basic_func_mult_return_type() (int, string) {
	return 1, "Hi"
}

func basic_variadic(nums ...int) int {
	sum := 0
	fmt.Print(nums, " ") // no new line with Print
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	return sum
}

func basic_variadic_type(a string, nums ...int) int {
	fmt.Println(a)
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum
}

func basic_closure() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func basic_factorial(x int) int {
	fmt.Println("Start x: ", x)
	if x == 0 {
		fmt.Println("x == 0 ", x)
		return 1
	}
	y := x * basic_factorial(x-1)
	fmt.Println("y: ", y)
	return y
}

func basic_fib(past int, cur int, breaker int) int {
	for cur <= breaker {
		return basic_fib(cur, (past + cur), breaker)
	}

	return cur
}

func basic_next_fib(next int) int {
	some := basic_fib(0, 1, next)
	return some
}

func main() {
	// basic_values()
	//basic_variables()
	//const s string = "constant"
	//basic_const(s)
	//basic_for_loop()
	//basic_if_else()
	//basic_switch()
	//basic_array()
	//basic_slices()
	//basic_maps()
	//basic_range()
	//fmt.Println(basic_function_add(1, 2))
	//basic_function_add_add(1, 2, 3, "hi")
	//a, b := basic_func_mult_return()
	//fmt.Println(a)
	//fmt.Println(b)

	//_, c := basic_func_mult_return()
	//fmt.Println(c)
	//d, e := basic_func_mult_return_type()
	//fmt.Println(d)
	//fmt.Println(e)

	//basic_variadic(1, 2, 3, 4)

	//c := make([]int, 0)
	//c = append(c, 1, 2, 3)
	//basic_variadic(c...)
	//basic_variadic_type("'String'", 3, 4, 5)

	////////////////////////////////
	// Closures
	////////////////////////////////
	//pos, neg := basic_closure(), basic_closure()

	//for i := 0; i < 10; i++ {
	//	fmt.Println(
	//		pos(i),
	//		neg(i*-1))
	//}

	//next_int := basic_closure()
	////The closure keeps track of the state for each initialized object
	//fmt.Println(next_int(1))
	//fmt.Println(next_int(2))
	//fmt.Println(next_int(3))

	//next_ints := basic_closure()
	//fmt.Println(next_ints(4))
	//fmt.Println(next_int(4))

	////////////////////////////////
	// Recursion
	////////////////////////////////
	//fmt.Println(basic_factorial(7))
	//fmt.Println(basic_fib(0, 1, 100))
	//fmt.Println(basic_next_fib(5))

	//var fib func(n int) int
	//fib = func(n int) int {
	//	if n < 2 {
	//		return n
	//	}
	//	return fib(n-1) + fib(n-2)

	//}
	//fmt.Println(fib(7))
}

package main

import (
	"fmt"
	"time"
)

func displayDate(format string, prefix string) {
	fmt.Println(prefix, time.Now().Format(format))
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func addNum(num1, num2 int) int {
	return num1 + num2
}

func countOddEven(s string) (int, int) {
	odds, evens := 0, 0
	for _, c := range s {
		if int(c)%2 == 0 {
			evens++
		} else {
			odds++
		}
	}
	return odds, evens
}

func addNums(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func fib() func() int {
	f1 := 0
	f2 := 1
	return func() int {
		f1, f2 = f2, f1+f2
		return f1
	}
}

func filter(arr []int, cond func(int) bool) []int {
	result := []int{}
	for _, v := range arr {
		if cond(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	displayDate("Mon 2006-01-02 15:04:05", "Current date and time: ")
	x := 5
	y := 10
	swap(&x, &y)
	fmt.Println(x, y)

	s := addNum(5, 6)
	fmt.Println(s)
	o, e := countOddEven("12345")
	fmt.Println(o, e)

	fmt.Println(addNums(1, 2, 3, 4, 5))
	fmt.Println(addNums(1, 2, 3))

	var i func() int
	i = func() int {
		return 5
	}
	fmt.Println(i())

	gen := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(gen())
	}

	a := []int{1, 2, 3, 4, 5}
	evens := filter(a,
		func(val int) bool {
			return val < 4
		})
	fmt.Println(evens)
}

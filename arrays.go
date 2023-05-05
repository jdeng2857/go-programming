package main

import (
	"errors"
	"fmt"
	"strconv"
)

func arrays() {
	nums := [5]int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println(len(nums))

	var table [5][6]string
	for row := 0; row < 5; row++ {
		for column := 0; column < 6; column++ {
			table[row][column] =
				strconv.Itoa(row) + "," + strconv.Itoa(column)
		}
	}
	fmt.Println(table)

	var cube [4][3][3]string
	for row := 0; row < 4; row++ {
		for column := 0; column < 3; column++ {
			for depth := 0; depth < 3; depth++ {
				cube[row][column][depth] =
					strconv.Itoa(row) + strconv.Itoa(column) + strconv.Itoa(depth)
			}
		}
	}
	fmt.Println(cube)
}

func slices() {
	s := make([]int, 5)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	r := make([]int, 2, 5)
	fmt.Println(len(r))
	fmt.Println(cap(r))

	t := []int{1, 2, 3, 4, 5}
	fmt.Println(len(t))
	fmt.Println(cap(t))
	t = append(t, 6, 7, 8)
	fmt.Println(t)
	fmt.Println(len(t))
	fmt.Println(cap(t))
	t = append(t, 9, 10)
	fmt.Println(len(t))
	fmt.Println(cap(t))
	u := t
	fmt.Println(u)
	fmt.Println(t)
	u[9] = 100
	fmt.Println(u)
	fmt.Println(t)
	t = append(t, 11)
	fmt.Println(u)
	fmt.Println(t)
	fmt.Println(len(u))
	fmt.Println(cap(u))
	fmt.Println(len(t))
	fmt.Println(cap(t))
}

func sliceOperations() {
	var c [3]string
	c[0] = "iOS"
	c[1] = "Android"
	c[2] = "Windows"
	fmt.Println(c[:2])
	fmt.Println(c[1:])
	fmt.Println(c[:])

	t := []int{1, 2, 3, 4, 5}
	fmt.Println(len(t))
	fmt.Println(cap(t))
	t = t[2:4]
	fmt.Println(t)
	fmt.Println(len(t))
	fmt.Println(cap(t))

	for i, v := range t {
		fmt.Println(i, v)
	}

	nums1 := [5]int{1, 2, 3, 4, 5}
	nums2 := nums1
	fmt.Println(nums1)
	fmt.Println(nums2)
	nums2[0] = 0
	fmt.Println(nums1)
	fmt.Println(nums2)
	t = []int{1, 2, 3, 4, 5}
	v := make([]int, 10)
	copy(v, t)
	fmt.Println(v)
	fmt.Println(t)
}

func insert(orig []int, index int, value int) ([]int, error) {
	if index < 0 {
		return nil, errors.New("Index cannot be less than 0")
	}

	if index >= len(orig) {
		return append(orig, value), nil
	}

	orig = append(orig[:index+1], orig[index:]...)
	orig[index] = value
	return orig, nil
}

func delete(orig []int, index int) ([]int, error) {
	if index < 0 || index >= len(orig) {
		return nil, errors.New("Index out of range")
	}
	orig = append(orig[:index], orig[index+1:]...)
	return orig, nil
}

func main() {
	arrays()
	slices()
	sliceOperations()

	t := []int{1, 2, 3, 4, 5}
	t, err := insert(t, 2, 9)
	if err == nil {
		fmt.Println(t)
	} else {
		fmt.Println(err)
	}
	t, _ = delete(t, 2)
	if err == nil {
		fmt.Println(t)
	} else {
		fmt.Println(err)
	}

}

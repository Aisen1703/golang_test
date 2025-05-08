package main

import (
	"fmt"
)

func uniqueElements(arr []int) []int {
	i := make(map[int]bool)
	for _, num := range arr {
		i[num] = true
	}

	result := make([]int, 0, len(i))
	for s := range i {
		result = append(result, s)
	}
	return result
}

func uniqueUnion(a, b []int) []int {
	w := make(map[int]bool)
	result := []int{}

	for _, moon := range a {
		if !w[moon] {
			w[moon] = true
			result = append(result, moon)
		}
	}
	for _, moon := range b {
		if !w[moon] {
			w[moon] = true
			result = append(result, moon)
		}
	}

	return result
}
func uniqueConst(a, b []int) []int {
	l := make(map[int]bool)
	var inter []int

	for _, v := range a {
		l[v] = true

	}

	for _, v := range b {
		if l[v] && !contains(inter, v) {
			inter = append(inter, v)
		}
	}

	return inter
}
func contains(x []int, i int) bool {
	for _, v := range x {
		if v == i {
			return true
		}
	}
	return false
}

func main() {
	var n int
	fmt.Print("Введите размер массива 1:")
	fmt.Scan(&n)
	a := make([]int, n)
	for i := range a {
		fmt.Printf("Введите элемент #%d:", i+1)
		fmt.Scan(&a[i])
	}
	var o int
	fmt.Print("Введите размер массива 2:")
	fmt.Scan(&o)
	b := make([]int, o)
	for g := range b {
		fmt.Printf("Введите элемент #%d:", g+1)
		fmt.Scan(&b[g])
	}

	m := uniqueElements(a)
	c := uniqueElements(b)
	v := uniqueUnion(a, b)
	x := uniqueConst(a, b)

	fmt.Printf("a=%d;", a)
	fmt.Printf(" b=%d\n", b)
	fmt.Printf("%d %d\n", m, c)
	fmt.Printf("%d\n", x)
	fmt.Printf("%d\n", v)

}

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

	fmt.Printf("a=%d;", a)
	fmt.Printf(" b=%d\n", b)
	fmt.Printf("%d %d\n", m, c)

}

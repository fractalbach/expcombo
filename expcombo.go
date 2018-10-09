/*
======================================================================

	 ~.,_,.~-^-~.,_ Expression Generation _,.~-^-~.,_,.~
			     Permutations
		 ___________________________________

			   Chris Achenbach
======================================================================
*/

package main

import (
	"fmt"
)

const line = "\n=================================================="

const (
	_ = iota
	add
	sub
	mul
	div
)

func makeNumPool() []int {
	arr := []int{}
	for i := 0; i < 5; i++ {
		arr = append(arr, i+1)
	}
	return arr
}

func makeOperPool() []int {
	return []int{add, sub, mul, div}
}

var (
	numpool     = makeNumPool()
	operpool    = makeOperPool()
	NumChannel  = make(chan []int)
	OperChannel = make(chan []int)
)

func main() {
	fmt.Printf("NumberPool:\n%v\n", numpool)
	fmt.Printf("IntPool:\n%v\n", operpool)
	fmt.Println("Results:")

	n := 5
	nums := make([]int, n)
	ops := make([]int, n-1)

	numCombos := [][]int{}
	operCombos := [][]int{}

	go StartGenerator(NumChannel, nums, numpool)
	for n := range NumChannel {
		numCombos = append(numCombos, n)
	}

	go StartGenerator(OperChannel, ops, operpool)
	for o := range OperChannel {
		operCombos = append(operCombos, o)
	}

	// fmt.Println(numCombos)
	// fmt.Println(operCombos)

	for _, n := range numCombos {
		for _, o := range operCombos {
			display(n, o)
		}
	}
}

func StartGenerator(results chan []int, arr, pool []int) {
	spot := 0
	generate(results, arr, pool, spot)
	close(results)
}

func generate(results chan []int, arr, pool []int, spot int) {
	if spot >= len(arr) {
		x := make([]int, len(arr))
		copy(x, arr)
		results <- x
		return
	}
	for i := range pool {
		remaining := make([]int, len(pool)-1)
		for j := range remaining {
			index := (j + i + 1) % len(pool)
			remaining[j] = pool[index]
		}
		arr[spot] = pool[i]
		generate(results, arr, remaining, spot+1)
	}
}

func display(nums []int, ops []int) {
	s := ""
	for i := 0; i < len(nums)-2; i++ {
		s += fmt.Sprintf("%v%v", nums[i], ShowOper(ops[i]))
	}
	s += fmt.Sprint(nums[len(nums)-1])
	fmt.Println(s)
}

func ShowOper(i int) string {
	switch i {
	case add:
		return "+"
	case sub:
		return "-"
	case mul:
		return "*"
	case div:
		return "%"
	}
	return "_"
}

func calc(nums []int, ops []int) {
	// TODO

	// NOTE:
	// use a stack to make sure you have operator order
}

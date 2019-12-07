package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	fmt.Printf("Q1: %v\n", Q1([]int{0, 1, 2, 3, 4}, q1))
	fmt.Printf("Q2: %v\n", Q2([]int{5, 6, 7, 8, 9}, q1))
}
func Q1(settings, arr []int) int {
	var high int
	for _, perm := range permutations(settings) {
		o := 0
		for _, s := range perm {
			o = tape(arr, []int{s, o})[0]
		}
		if o > high {
			high = o
		}
	}
	return high
}

func Q2(settings, arr []int) int {
	var high int
	for _, perm := range permutations(settings) {
		pipes := make([]chan int, len(perm))
		for i := range pipes {
			pipes[i] = make(chan int, 1)
		}
		var wg sync.WaitGroup
		for i, s := range perm {
			wg.Add(1)
			go tape2(arr, pipes[i], pipes[(i+1)%5], &wg)
			pipes[i] <- s
		}
		pipes[0] <- 0
		wg.Wait()
		out := <-pipes[0]
		if out > high {
			high = out
		}
	}
	return high
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func tape2(sarr []int, in, out chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	arr := make([]int, len(sarr)+5)
	copy(arr, sarr)
	p := 0
	for arr[p]%100 != 99 {
		op := arr[p]
		m1, m2, _ := (op/100)%10, (op/1000)%10, (op/10000)%10
		// immediate mode, dereference as needed
		p1, p2, p3 := arr[p+1], arr[p+2], arr[p+3]

		switch op % 100 {
		case 1, 2, 5, 6, 7, 8:
			if m2 == 0 {
				p2 = arr[p2]
			}
			fallthrough
		case 4:
			if m1 == 0 {
				p1 = arr[p1]
			}

		}

		switch op % 100 {
		case 1:
			// p1 + p2
			arr[p3] = p1 + p2
			p += 4
		case 2:
			// p1 * p2
			arr[p3] = p1 * p2
			p += 4
		case 3:
			// read from input
			arr[p1] = <-in
			p += 2
		case 4:
			// output
			out <- p1
			p += 2
		case 5:
			// jump if not equal
			if p1 != 0 {
				p = p2
			} else {
				p += 3
			}
		case 6:
			// jump if equal
			if p1 == 0 {
				p = p2
			} else {
				p += 3
			}
		case 7:
			// p1 < p2n
			if p1 < p2 {
				arr[p3] = 1
			} else {
				arr[p3] = 0
			}
			p += 4
		case 8:
			// p1 == p2
			if p1 == p2 {
				arr[p3] = 1
			} else {
				arr[p3] = 0
			}
			p += 4

		case 99:
			// breakout
			return
		default:
			log.Fatalf("got opcode %d at position %d", arr[p], p)
		}
	}
	return
}

func tape(sarr, in []int) []int {
	arr := make([]int, len(sarr)+5)
	copy(arr, sarr)
	p := 0
	var out []int
	for arr[p]%100 != 99 {
		op := arr[p]
		m1, m2, _ := (op/100)%10, (op/1000)%10, (op/10000)%10
		// immediate mode, dereference as needed
		p1, p2, p3 := arr[p+1], arr[p+2], arr[p+3]

		switch op % 100 {
		case 1, 2, 5, 6, 7, 8:
			if m2 == 0 {
				p2 = arr[p2]
			}
			fallthrough
		case 4:
			if m1 == 0 {
				p1 = arr[p1]
			}

		}

		switch op % 100 {
		case 1:
			// p1 + p2
			arr[p3] = p1 + p2
			p += 4
		case 2:
			// p1 * p2
			arr[p3] = p1 * p2
			p += 4
		case 3:
			// read from input
			arr[p1] = in[0]
			in = in[1:]
			p += 2
		case 4:
			// output
			out = append(out, p1)
			p += 2
		case 5:
			// jump if not equal
			if p1 != 0 {
				p = p2
			} else {
				p += 3
			}
		case 6:
			// jump if equal
			if p1 == 0 {
				p = p2
			} else {
				p += 3
			}
		case 7:
			// p1 < p2n
			if p1 < p2 {
				arr[p3] = 1
			} else {
				arr[p3] = 0
			}
			p += 4
		case 8:
			// p1 == p2
			if p1 == p2 {
				arr[p3] = 1
			} else {
				arr[p3] = 0
			}
			p += 4

		case 99:
			// breakout
			return out
		default:
			log.Fatalf("got opcode %d at position %d", arr[p], p)
		}
	}
	return out
}

var (
	q1 = []int{
		3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 38, 47, 72, 97, 122, 203, 284, 365, 446, 99999, 3, 9, 1001, 9, 3, 9, 1002, 9, 5, 9, 1001, 9, 4, 9, 4, 9, 99, 3, 9, 102, 3, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 2, 9, 102, 5, 9, 9, 101, 3, 9, 9, 1002, 9, 5, 9, 101, 4, 9, 9, 4, 9, 99, 3, 9, 101, 5, 9, 9, 1002, 9, 3, 9, 101, 2, 9, 9, 102, 3, 9, 9, 1001, 9, 2, 9, 4, 9, 99, 3, 9, 101, 3, 9, 9, 102, 2, 9, 9, 1001, 9, 4, 9, 1002, 9, 2, 9, 101, 2, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 99, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 99, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 99,
	}
)

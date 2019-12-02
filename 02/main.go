package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	arr := make([]int, len(q1))
	copy(arr, q1)
	arr[1], arr[2] = 12, 2
	p := 0
loop:
	for {
		switch arr[p] {
		case 1:
			arr[arr[p+3]] = arr[arr[p+1]] + arr[arr[p+2]]
		case 2:
			arr[arr[p+3]] = arr[arr[p+1]] * arr[arr[p+2]]
		case 99:
			fmt.Printf("final out: %d\n", arr[0])
			break loop
		default:
			log.Fatalf("got opcode %d at position %d", arr[p], p)
		}
		p += 4
	}

	var wg sync.WaitGroup
	for a := 0; a < 100; a++ {
		for b := 0; b < 100; b++ {
			wg.Add(1)
			go func(a, b int) {
				defer wg.Done()
				arr := make([]int, len(q1))
				copy(arr, q1)
				arr[1], arr[2] = a, b
				p := 0
			loop2:
				for {
					switch arr[p] {
					case 1:
						arr[arr[p+3]] = arr[arr[p+1]] + arr[arr[p+2]]
					case 2:
						arr[arr[p+3]] = arr[arr[p+1]] * arr[arr[p+2]]
					case 99:
						if arr[0] == 19690720 {
							fmt.Printf("found result with %d and %d, ans %d", a, b, 100*a+b)
						}
						break loop2
					default:
						log.Fatalf("got opcode %d at position %d", arr[p], p)
					}
					p += 4
				}

			}(a, b)
		}

	}
	wg.Wait()

}

var q1 = []int{
	1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 13, 19, 2, 9, 19, 23, 1, 23, 6, 27, 1, 13, 27, 31, 1, 31, 10, 35, 1, 9, 35, 39, 1, 39, 9, 43, 2, 6, 43, 47, 1, 47, 5, 51, 2, 10, 51, 55, 1, 6, 55, 59, 2, 13, 59, 63, 2, 13, 63, 67, 1, 6, 67, 71, 1, 71, 5, 75, 2, 75, 6, 79, 1, 5, 79, 83, 1, 83, 6, 87, 2, 10, 87, 91, 1, 9, 91, 95, 1, 6, 95, 99, 1, 99, 6, 103, 2, 103, 9, 107, 2, 107, 10, 111, 1, 5, 111, 115, 1, 115, 6, 119, 2, 6, 119, 123, 1, 10, 123, 127, 1, 127, 5, 131, 1, 131, 2, 135, 1, 135, 5, 0, 99, 2, 0, 14, 0,
}

var t1 = []int{
	1, 9, 10, 3,
	2, 3, 11, 0,
	99,
	30, 40, 50,
}

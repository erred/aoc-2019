package main

import (
	"fmt"
)

func main() {
	var cnt int
loop:
	for i := q1[0]; i <= q1[1]; i++ {
		p, j := i%10, i/10
		var d bool
		for j > 0 {
			if j%10 > p {
				continue loop
			} else if j%10 == p {
				d = true
			}
			p = j % 10
			j /= 10
		}
		if d {
			cnt++
		}
	}
	fmt.Println(cnt)
	cnt = 0
loop2:
	for i := q1[0]; i <= q1[1]; i++ {
		p, j := i%10, i/10
		var d bool
		var dd int
		for j > 0 {
			if j%10 > p {
				continue loop2
			} else if j%10 == p {
				dd++
			} else {
				if dd == 1 {
					d = true
				}
				dd = 0
			}
			p = j % 10
			j /= 10
		}
		if dd == 1 {
			d = true
		}
		if d {
			cnt++
		}
	}
	fmt.Println(cnt)
}

var q1 = []int{
	387638, 919123,
}

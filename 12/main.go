package main

import (
	"fmt"
)

type pt struct {
	p [3]int
	v [3]int
}

func (p pt) String() string {
	return fmt.Sprintf("pos=<x=%3d, y=%3d, z=%3d>, vel=<x=%3d, y=%3d, z=%3d>", p.p[0], p.p[1], p.p[2], p.v[0], p.v[1], p.v[2])
}

func (p pt) energy() int {
	var pe, ke int
	for i := range p.p {
		if p.p[i] < 0 {
			pe -= p.p[i]
		} else {
			pe += p.p[i]
		}
		if p.v[i] < 0 {
			ke -= p.v[i]
		} else {
			ke += p.v[i]
		}
	}
	return pe * ke
}

type Moons []pt

func (moons Moons) update() {
	for i, m := range moons {
		for j := i + 1; j < len(moons); j++ {
			n := moons[j]
			for d := 0; d < 3; d++ {
				if dir := m.p[d] - n.p[d]; dir < 0 {
					moons[i].v[d]++
					moons[j].v[d]--
				} else if dir > 0 {
					moons[i].v[d]--
					moons[j].v[d]++
				}
			}
		}
	}
	for i, m := range moons {
		for d := 0; d < 3; d++ {
			moons[i].p[d] += m.v[d]
		}
	}
}
func (moons Moons) energy() int {
	var sum int
	for _, m := range moons {
		sum += m.energy()
	}
	return sum
}

func main() {
	fmt.Println(Q1(q1, 1000))
	fmt.Println(Q2(q1))
}

func Q2(in []string) int64 {
	moons := make(Moons, len(in))
	for i, n := range in {
		fmt.Sscanf(n, "<x=%d, y=%d, z=%d>", &moons[i].p[0], &moons[i].p[1], &moons[i].p[2])
	}

	states := []map[[8]int]int64{
		make(map[[8]int]int64),
		make(map[[8]int]int64),
		make(map[[8]int]int64),
	}

	cycle := make([]int64, 3)
	var s [8]int
	for j := 0; j < 3; j++ {
		for i, m := range moons {
			s[i], s[i+1] = m.p[j], m.v[j]
		}
		states[j][s] = 0
	}
	var c int64
	for {
		c++
		if c%1000000 == 0 {
			fmt.Print(" Progress ", c)
		}
		moons.update()

		for j := 0; j < 3; j++ {
			for i, m := range moons {
				s[i], s[i+1] = m.p[j], m.v[j]
			}
			if co, ok := states[j][s]; !ok {
				// states[j][s] = c
			} else {
				fmt.Printf("d: %v cycle old: %v new %v\n", j, co, c)
				cycle[j] = c
			}
		}
		if cycle[0] != 0 && cycle[1] != 0 && cycle[2] != 0 {
			break
		}
	}
	return LCM(cycle[0], cycle[1], cycle[2])
}

func Q1(in []string, ts int) int {
	moons := make(Moons, len(in))
	for i, n := range in {
		fmt.Sscanf(n, "<x=%d, y=%d, z=%d>", &moons[i].p[0], &moons[i].p[1], &moons[i].p[2])
	}

	for t := 0; t < ts; t++ {
		moons.update()
	}
	return moons.energy()

}

var (
	q1 = []string{
		"<x=-14, y=-4, z=-11>",
		"<x=-9, y=6, z=-7>",
		"<x=4, y=1, z=4>",
		"<x=2, y=-14, z=-9>",
	}
)

func GCD(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(ints ...int64) int64 {
	result := ints[0] * ints[1] / GCD(ints[0], ints[1])
	for _, in := range ints[2:] {
		result = LCM(result, in)
	}
	return result
}

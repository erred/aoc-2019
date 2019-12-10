package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
)

func main() {
	x, y, c := Q1(q1)
	fmt.Println("Q1", x, y, c)
	x, y = Q2(q1, point{13, 17}, 200)
	fmt.Println("Q2", x, y)
}

type point struct {
	x, y int
}

func Q1(in string) (x, y, c int) {
	inl := strings.Split(in, "\n")

	m := make(map[point]int)
	for y := range inl {
		for x := range inl[y] {
			if inl[y][x] == '.' {
				continue
			}
			m[point{x, y}] = 0
		}
	}
	var wg sync.WaitGroup
	var cc = make(chan point)
	for p := range m {
		wg.Add(1)
		go func(p point) {
			defer wg.Done()
		loop:
			for p2 := range m {
				if p == p2 {
					continue
				}
				dx, dy := p2.x-p.x, p2.y-p.y
				g := gcd(dx, dy)
				dx, dy = dx/g, dy/g
				for mu := 1; mu < g; mu++ {
					pt := point{p.x + mu*dx, p.y + mu*dy}
					// if p == (point{4, 0}) {
					// 	fmt.Println(p, pt, p2)
					// }
					if _, ok := m[pt]; ok {
						continue loop
					}
				}
				cc <- p
			}

		}(p)
	}
	go func() {
		wg.Wait()
		close(cc)
	}()
	m2 := make(map[point]int)
	for p := range cc {
		m2[p]++
	}

	mp, mc := point{}, 0
	for p, c := range m2 {
		if c > mc {
			mc, mp = c, p
		}
	}
	return mp.x, mp.y, mc
}

func gcd(a, b int) int {
	if a < 0 {
		a *= -1
	}
	if b < 0 {
		b *= -1
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Q2(in string, p point, t int) (x, y int) {
	inl := strings.Split(in, "\n")

	m := make(map[float64][]point)
	for y := range inl {
		for x := range inl[y] {
			if inl[y][x] == '.' {
				continue
			}
			p2 := point{x, y}
			if p2 == p {
				continue
			}
			dx, dy := p2.x-p.x, p.y-p2.y
			r := -math.Atan2(float64(dy), float64(dx))
			if r < -math.Pi/2 {
				r += math.Pi * 2
			}
			m[r] = append(m[r], p2)
		}
	}
	var d point
	var s []float64
	for k := range m {
		s = append(s, k)
		sort.Slice(m[k], func(i, j int) bool {
			return m[k][i].y > m[k][j].y
		})
	}
	sort.Float64s(s)
	var c int
	for i := 0; i < t; i++ {
		// fmt.Println(s[c], m[s[c]])
		dp := m[s[c]]
		d = dp[0]
		if len(dp) == 1 {
			delete(m, s[c])
			s = append(s[:c], s[c+1:]...)
		} else {
			m[s[c]] = m[s[c]][1:]
			c++
		}
		if len(s) > 0 {
			c %= len(s)
		}
	}
	return d.x, d.y
}

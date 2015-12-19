package main

import (
	"fmt"
)

type Board struct {
	num, occupied, coverage int64
}

func (b *Board) String() string {
	s := ""
	for i := 0; i < 64; i++ {
		if b.Occupied(int64(i/8), int64(i%8)) {
			s += "X"
		} else {
			s += "-"
		}
		if i > 0 && i % 8 == 7 {
			s += "\n"
		}
	}
	return s
}

func Pos(r, c int64) int64 {
	return int64(1) << uint(r * int64(8) + c)
}

func (b *Board) Occupied(row, col int64) bool {
	return b.occupied & Pos(row, col) != 0
}

func (b *Board) Blocking(row, col int64) bool {
	return b.coverage & Pos(row, col) != 0
}

func Abs(a int64) int64 {
	if (a < 0) { 
		return -1*a
	}
	return a
}

func Put(b *Board, row, col int64) *Board {
	if b.Blocking(row, col) {
		return nil
	} 

	next_coverage := int64(b.coverage)
	for dr := int64(-7); dr <= 7; dr++ {
		for dc := int64(-7); dc <= 7; dc++ {
			nc := col + dc
			nr := row + dr
			if 0 <= nr && nr < 8 && 0 <= nc && nc < 8 {
				if dr == 0 || dc == 0 || Abs(dc) == Abs(dr) {
					next_coverage |= Pos(nr, nc)
				}
			}
		}
	}

	return &Board{b.num + 1, 
		b.occupied | Pos(row, col),
		next_coverage}
}

func (b *Board) Solve() {
	if b.num == 8 {
		fmt.Print(b.String())
		fmt.Println("*******************************")
	} else {
		for col := int64(0); col < 8; col++ {
			nextB := Put(b, b.num, col)
			if nextB != nil {
				nextB.Solve()
			}
		}
	}
}

func (b *Board) ParallellSolve(ch chan string) {
	if b.num == 8 {
		ch <- b.String()
	} else {
		for col := int64(0); col < 8; col++ {
			nextB := Put(b, b.num, col)
			if nextB != nil {
				go nextB.ParallellSolve(ch)
			}
		}
	}
}

func main() {
	// This is to solve normally
	// Board{0, 0, 0}.Solve()

	// This is to solve using goroutines
	ch := make(chan string)
	b := Board{0, 0, 0}
	go (&b).ParallellSolve(ch)
	for i := 0; i < 92; i++ {
		fmt.Print(<-ch)
		fmt.Println("********************************")
	}
}

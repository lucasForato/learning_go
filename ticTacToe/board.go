package main

import "fmt"

type Board [3][3]rune

type Vec2 struct {
	x int8
	y int8
}

func (b *Board) PrintBoard() {
	for i, row := range b {
		if i == 0 {
			fmt.Println("  A   B   C ")
		}
		for j, cell := range row {
			if j == 0 {
				fmt.Printf("%v", i+1)
			}
			fmt.Printf(" %c ", cell)
			if j < 2 {
				fmt.Printf("|")
			}
		}
		fmt.Printf("\n")
		if i < 2 {
			fmt.Printf(" ---|---|--- \n")
		}
	}
  fmt.Printf("\n")
}

type Player rune

const (
	X Player = 'X'
	O Player = 'O'
)

func (b *Board) RecordMove(player Player, position Vec2) {
	if player != X && player != O {
		panic("Invalid player")
	}

	b[position.x][position.y] = rune(player)
}

func (b *Board) GetWinner() (bool, Player) {
	if b[0][0] == b[0][1] && b[0][1] == b[0][2] {
		if b[0][0] != '-' {
			return true, Player(b[0][0])
		}
	}

	if b[1][0] == b[1][1] && b[1][1] == b[1][2] {
		if b[1][0] != '-' {
			return true, Player(b[1][0])
		}
	}

	if b[2][0] == b[2][1] && b[2][1] == b[2][2] {
		if b[2][0] != '-' {
			return true, Player(b[2][0])
		}
	}

	if b[0][0] == b[1][0] && b[1][0] == b[2][0] {
		if b[0][0] != '-' {
			return true, Player(b[0][0])
		}
	}

	if b[0][1] == b[1][1] && b[1][1] == b[2][1] {
		if b[0][1] != '-' {
			return true, Player(b[0][1])
		}
	}

	if b[0][2] == b[1][2] && b[1][2] == b[2][2] {
		if b[0][2] != '-' {
			return true, Player(b[0][2])
		}
	}

	if b[0][0] == b[1][1] && b[1][1] == b[2][2] {
		if b[0][0] != '-' {
			return true, Player(b[0][0])
		}
	}

	if b[2][0] == b[1][1] && b[1][1] == b[0][2] {
		if b[2][0] != '-' {
			return true, Player(b[2][0])
		}
	}

	return false, 0
}

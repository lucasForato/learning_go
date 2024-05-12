package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func ClearScreen() {
	var cmd *exec.Cmd
	cmd = exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	board := Board{
		{'-', '-', '-'},
		{'-', '-', '-'},
		{'-', '-', '-'},
	}
	var counter int

	for {
		ClearScreen()
		board.PrintBoard()

		var x string
		var y int8
		fmt.Print("Type the X coordinate (A, B, C): ")
		fmt.Scan(&x)
		fmt.Println(x)

		fmt.Print("Type the Y coordinate(1, 2, 3): ")
		fmt.Scan(&y)

		var vec, err = makeCoordinate(x, y)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		var player Player = X
		if counter%2 == 0 {
			player = X
		} else {
			player = O
		}

		board.RecordMove(player, vec)

		hasWinner, player := board.GetWinner()
		if hasWinner {
			fmt.Println("There is a winner:", player)
			break
		}
	}

}

func makeCoordinate(x string, y int8) (Vec2, error) {
	if y != 1 && y != 2 && y != 3 {
		return Vec2{}, errors.New("Invalid Y input.")
	}

	var vec Vec2

	switch x {
	case "A":
		vec.x = 0

	case "B":
		vec.x = 1

	case "C":
		vec.x = 2

	default:
		return Vec2{}, errors.New("Invalid X input.")
	}

	vec.y = y - 1

	return vec, nil
}

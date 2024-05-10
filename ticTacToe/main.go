package main

import (
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

	for {
		var x int8
		var y int8
		fmt.Print("Type the X coordinate: ")
		fmt.Scan(&x)

		fmt.Print("Type the Y coordinate: ")
		fmt.Scan(&y)
		board.RecordMove(X, Vec2{x, y})
		ClearScreen()
		board.PrintBoard()

    hasWinner, player := board.GetWinner()
    if hasWinner {
      fmt.Println("There is a winner:", player)
      break
    }
	}

}

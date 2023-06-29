package main

import (
	"flag"
	"fmt"
)

func main() {
	var rows int
	var cols int
	flag.IntVar(&rows, "rows", 20, "Number of rows wanted")
	flag.IntVar(&cols, "Columns", 20, "Number of columns wanted")
	gameDungeon := initDungeon(rows, cols)
	for i := 0; i < len(gameDungeon.rooms); i++ {
		for x := 0; x < len(gameDungeon.rooms[i]); x++ {
			fmt.Printf("%c ", gameDungeon.rooms[i][x])
			if x == (gameDungeon.rows - 1) {
				fmt.Println("")
			}
		}
	}
}

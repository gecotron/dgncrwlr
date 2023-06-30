package main

import (
	"flag"
	"fmt"
)

func main() {
	var rows int
	var columns int
	flag.IntVar(&rows, "r", 20, "Number of rows wanted")
	flag.IntVar(&columns, "c", 20, "Number of columns wanted")
	flag.Parse()
	gameDungeon := initDungeon(rows, columns)
	for i := 0; i < len(gameDungeon.rooms); i++ {
		for x := 0; x < len(gameDungeon.rooms[i]); x++ {
			fmt.Printf("%c ", gameDungeon.rooms[i][x])
			if x == (gameDungeon.rows - 1) {
				fmt.Println("")
			}
		}
	}
}

package main

import (
	"flag"
	"fmt"
)

func main() {
	// Game / Debugging flags
	var rows int
	var columns int
	flag.IntVar(&rows, "r", 20, "Number of rows wanted")
	flag.IntVar(&columns, "c", 20, "Number of columns wanted")
	flag.Parse()
	// Generate dungeon
	gameDungeon := initDungeon(rows, columns)
	// Generate player
	p := initPlayer('c', "Alix", rows, columns)
	// Display the dungeon and the player
	for i := 0; i < len(gameDungeon.rooms); i++ {
		for x := 0; x < len(gameDungeon.rooms[i]); x++ {
			if i == p.X && x == p.Y {
				fmt.Printf("%c ", '@')
			} else {
				fmt.Printf("%c ", gameDungeon.rooms[i][x])
			}
			if x == (gameDungeon.rows - 1) {
				fmt.Println("")
			}
		}
	}
	saveUser(p, p.Name)
}
func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

package main

import "fmt"

func main() {
	gameDungeon := initDungeon(10, 10)
	for i := 0; i < len(gameDungeon.rooms); i++ {
		for x := 0; x < len(gameDungeon.rooms[i]); x++ {
			fmt.Printf("%c", gameDungeon.rooms[i][x])
			if x == (gameDungeon.rows - 1) {
				fmt.Println("")
			}
		}
	}
}
